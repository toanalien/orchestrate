// +build integration

package integrationtests

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	http2 "net/http"
	"strings"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gofrs/uuid"
	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/encoding/json"
	encoding "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/encoding/proto"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/ethclient/rpc"
	utils2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/ethclient/utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/http"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/http/httputil"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/tx"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/txscheduler"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/utils"
	"gopkg.in/h2non/gock.v1"
)

const (
	waitForEnvelopeTimeOut = 2 * time.Second
)

// txSignerEthereumTestSuite is a test suite for Transaction signer Ethereum
type txSignerEthereumTestSuite struct {
	suite.Suite
	env *IntegrationEnvironment
}

func (s *txSignerEthereumTestSuite) TestTxSender_Ethereum_Public() {
	signature := "0xd35c752d3498e6f5ca1630d264802a992a141ca4b6a3f439d673c75e944e5fb05278aaa5fabbeac362c321b54e298dedae2d31471e432c26ea36a8d49cf08f1e01"
	raw := "0xf85380839896808252088083989680808216b4a0d35c752d3498e6f5ca1630d264802a992a141ca4b6a3f439d673c75e944e5fb0a05278aaa5fabbeac362c321b54e298dedae2d31471e432c26ea36a8d49cf08f1e"
	txHash := "0x6621fbe1e2848446e38d99bfda159cdd83f555ae0ed7a4f3e1c3c79f7d6d74f3"

	s.T().Run("should sign and send public ethereum transaction successfully", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = s.env.nm.DeleteLastSent(envelope.PartitionKey())

		gock.New(keyManagerURL).
			Post(fmt.Sprintf("/ethereum/accounts/%s/sign-transaction", envelope.GetFromString())).
			Reply(200).BodyString(signature)

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_getTransactionCount", envelope.From.String(), "pending")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x0\"}")

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_sendRawTransaction", raw)).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"" + txHash + "\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusPending)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)
	})

	s.T().Run("should sign and send a public onetimekey ethereum transaction successfully", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_sendRawTransaction")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"" + txHash + "\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusPending)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		// IMPORTANT: As we cannot infer txHash before hand, status will be updated to WARNING
		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusWarning)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest().EnableTxFromOneTimeKey())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)
	})

	s.T().Run("should send envelope to tx-recover if key-manager fails to sign", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = s.env.nm.DeleteLastSent(envelope.PartitionKey())

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_getTransactionCount", envelope.From.String(), "pending")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x0\"}")

		url := fmt.Sprintf("/ethereum/accounts/%s/sign-transaction", envelope.GetFromString())
		gock.New(keyManagerURL).Post(url).Reply(http2.StatusUnauthorized).Status(422).
			JSON(httputil.ErrorResponse{
				Message: "not authorized requests",
				Code:    666,
			})

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusFailed)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)

		retrievedEnvelope, err := s.env.consumer.WaitForEnvelope(envelope.GetID(), s.env.srvConfig.RecoverTopic,
			waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, envelope.GetID(), retrievedEnvelope.GetID())
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Code)
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Message)
	})

	s.T().Run("should send envelope to tx-recover if transaction-scheduler fails to update", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = s.env.nm.DeleteLastSent(envelope.PartitionKey())

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_getTransactionCount", envelope.From.String(), "pending")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x0\"}")

		gock.New(keyManagerURL).
			Post(fmt.Sprintf("/ethereum/accounts/%s/sign-transaction", envelope.GetFromString())).
			Reply(200).BodyString(signature)

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_sendRawTransaction")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"" + txHash + "\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusPending)).
			Reply(422).JSON(httputil.ErrorResponse{
			Message: "cannot update status",
			Code:    666,
		})

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusFailed)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		retrievedEnvelope, err := s.env.consumer.WaitForEnvelope(envelope.GetID(), s.env.srvConfig.RecoverTopic,
			waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, envelope.GetID(), retrievedEnvelope.GetID())
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Code)
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Message)
	})

	s.T().Run("should send envelope to tx-recover if chain-proxy response with an error", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = s.env.nm.DeleteLastSent(envelope.PartitionKey())

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_getTransactionCount", envelope.From.String(), "pending")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x0\"}")

		gock.New(keyManagerURL).
			Post(fmt.Sprintf("/ethereum/accounts/%s/sign-transaction", envelope.GetFromString())).
			Reply(200).BodyString(signature)

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_sendRawTransaction")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"error\":\"invalid_raw\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusPending)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusFailed)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)

		retrievedEnvelope, err := s.env.consumer.WaitForEnvelope(envelope.GetID(), s.env.srvConfig.RecoverTopic,
			waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, envelope.GetID(), retrievedEnvelope.GetID())
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Code)
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Message)
	})
}

func (s *txSignerEthereumTestSuite) TestTxSender_Ethereum_Raw_Public() {
	// signature := "0xd35c752d3498e6f5ca1630d264802a992a141ca4b6a3f439d673c75e944e5fb05278aaa5fabbeac362c321b54e298dedae2d31471e432c26ea36a8d49cf08f1e01"
	raw := "0xf85380839896808252088083989680808216b4a0d35c752d3498e6f5ca1630d264802a992a141ca4b6a3f439d673c75e944e5fb0a05278aaa5fabbeac362c321b54e298dedae2d31471e432c26ea36a8d49cf08f1e"
	txHash := "0x6621fbe1e2848446e38d99bfda159cdd83f555ae0ed7a4f3e1c3c79f7d6d74f3"

	s.T().Run("should send ethereum raw transaction successfully", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = envelope.SetJobType(tx.JobType_ETH_RAW_TX)
		_ = envelope.SetRawString(raw)
		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_sendRawTransaction", raw)).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"" + txHash + "\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusPending)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest().EnableTxFromOneTimeKey())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)
	})

	s.T().Run("should send envelope to tx-recover if send ethereum raw transaction fails", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = envelope.SetJobType(tx.JobType_ETH_RAW_TX)
		_ = envelope.SetRawString(raw)

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_sendRawTransaction", raw)).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"invalid_raw\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusPending)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusFailed)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest().EnableTxFromOneTimeKey())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)

		retrievedEnvelope, err := s.env.consumer.WaitForEnvelope(envelope.GetID(), s.env.srvConfig.RecoverTopic,
			waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, envelope.GetID(), retrievedEnvelope.GetID())
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Code)
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Message)
	})
}

func (s *txSignerEthereumTestSuite) TestTxSender_Ethereum_EEA() {
	signature := "0xd35c752d3498e6f5ca1630d264802a992a141ca4b6a3f439d673c75e944e5fb05278aaa5fabbeac362c321b54e298dedae2d31471e432c26ea36a8d49cf08f1e01"
	raw := "0xf8c380839896808252088083989680808216b4a0d35c752d3498e6f5ca1630d264802a992a141ca4b6a3f439d673c75e944e5fb0a05278aaa5fabbeac362c321b54e298dedae2d31471e432c26ea36a8d49cf08f1ea0035695b4cc4b0941e60551d7a19cf30603db5bfc23e5ac43a56f57f25f75486af842a0035695b4cc4b0941e60551d7a19cf30603db5bfc23e5ac43a56f57f25f75486aa0075695b4cc4b0941e60551d7a19cf30603db5bfc23e5ac43a56f57f25f75486a8a72657374726963746564"
	txHash := "0x6621fbe1e2848446e38d99bfda159cdd83f555ae0ed7a4f3e1c3c79f7d6d74f3"

	s.T().Run("should sign and send a EEA transaction successfully", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = envelope.SetJobType(tx.JobType_ETH_ORION_EEA_TX)
		_ = s.env.nm.DeleteLastSent(envelope.PartitionKey())

		gock.New(keyManagerURL).
			Post(fmt.Sprintf("/ethereum/accounts/%s/sign-eea-transaction", envelope.GetFromString())).
			Reply(200).BodyString(signature)

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "priv_getEeaTransactionCount", envelope.From.String(),
				envelope.PrivateFrom, envelope.PrivateFor)).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x0\"}")

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "priv_distributeRawTransaction", raw)).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"" + txHash + "\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusStored)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)
	})

	s.T().Run("should sign and send EEA transaction with oneTimeKey successfully", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = envelope.SetJobType(tx.JobType_ETH_ORION_EEA_TX)
		_ = s.env.nm.DeleteLastSent(envelope.PartitionKey())

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "priv_distributeRawTransaction")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"" + txHash + "\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusStored)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest().EnableTxFromOneTimeKey())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)
	})

	s.T().Run("should send a envelope to tx-recover if we fail to send EEA transaction", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = envelope.SetJobType(tx.JobType_ETH_ORION_EEA_TX)

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "priv_distributeRawTransaction")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"invalid_raw\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusFailed)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest().EnableTxFromOneTimeKey())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)

		retrievedEnvelope, err := s.env.consumer.WaitForEnvelope(envelope.GetID(), s.env.srvConfig.RecoverTopic,
			waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, envelope.GetID(), retrievedEnvelope.GetID())
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Code)
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Message)
	})
}

func (s *txSignerEthereumTestSuite) TestTxSender_Tessera_Marking() {
	signature := "0xd35c752d3498e6f5ca1630d264802a992a141ca4b6a3f439d673c75e944e5fb05278aaa5fabbeac362c321b54e298dedae2d31471e432c26ea36a8d49cf08f1e01"
	raw := "0xf851808398968082520880839896808026a0d35c752d3498e6f5ca1630d264802a992a141ca4b6a3f439d673c75e944e5fb0a05278aaa5fabbeac362c321b54e298dedae2d31471e432c26ea36a8d49cf08f1e"
	txHash := "0x226d79b217b5ebfeddd08662f3ae1bb1b2cb339d50bbcb708b53ad5f4c71c5ea"

	s.T().Run("should sign and send Tessera marking transaction successfully", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}
	
		envelope := fakeEnvelope()
		_ = envelope.SetJobType(tx.JobType_ETH_TESSERA_MARKING_TX)
		_ = s.env.nm.DeleteLastSent(envelope.PartitionKey())
	
		gock.New(keyManagerURL).
			Post(fmt.Sprintf("/ethereum/accounts/%s/sign-quorum-private-transaction", envelope.GetFromString())).
			Reply(200).BodyString(signature)
	
		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_getTransactionCount", envelope.From.String(), "pending")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x0\"}")
	
		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_sendRawPrivateTransaction", raw, map[string]interface{}{
				"privateFor": envelope.PrivateFor,
			})).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"" + txHash + "\"}")
	
		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusPending)).
			Reply(200).JSON(&txscheduler.JobResponse{})
	
		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
	
		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)
	})
	
	s.T().Run("should sign and send Tessera marking transaction with oneTimeKey successfully", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}
	
		envelope := fakeEnvelope()
		_ = envelope.SetJobType(tx.JobType_ETH_TESSERA_MARKING_TX)
	
		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_sendRawPrivateTransaction")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"" + txHash + "\"}")
	
		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusPending)).
			Reply(200).JSON(&txscheduler.JobResponse{})
		
		// HASH won't match
		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusWarning)).
			Reply(200).JSON(&txscheduler.JobResponse{})
	
		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest().EnableTxFromOneTimeKey())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
	
		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)
	})

	s.T().Run("should send a envelope to tx-recover if we fail to send Tessera marking transaction", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		_ = envelope.SetJobType(tx.JobType_ETH_TESSERA_MARKING_TX)
		_ = s.env.nm.DeleteLastSent(envelope.PartitionKey())

		gock.New(keyManagerURL).
			Post(fmt.Sprintf("/ethereum/accounts/%s/sign-quorum-private-transaction", envelope.GetFromString())).
			Reply(200).BodyString(signature)

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_getTransactionCount", envelope.From.String(), "pending")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x0\"}")

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/%s", envelope.GetChainUUID())).
			AddMatcher(ethCallMatcher(wg, "eth_sendRawPrivateTransaction")).
			Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"error\":\"invalid_raw\"}")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusPending)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusFailed)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)

		retrievedEnvelope, err := s.env.consumer.WaitForEnvelope(envelope.GetID(), s.env.srvConfig.RecoverTopic,
			waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, envelope.GetID(), retrievedEnvelope.GetID())
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Code)
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Message)
	})
}

func (s *txSignerEthereumTestSuite) TestTxSender_Tessera_Private() {
	data := "0xf8c380839896808252088083989680808216b4a0d35c752d3498e6f5ca1630d264802a992a141ca4b6a3f439d673c75e944e5fb0a05278aaa5fabbeac362c321b54e298dedae2d31471e432c26ea36a8d49cf08f1ea0035695b4cc4b0941e60551d7a19cf30603db5bfc23e5ac43a56f57f25f75486af842a0035695b4cc4b0941e60551d7a19cf30603db5bfc23e5ac43a56f57f25f75486aa0075695b4cc4b0941e60551d7a19cf30603db5bfc23e5ac43a56f57f25f75486a8a72657374726963746564"
	enclaveKey := "0x226d79b217b5ebfeddd08662f3ae1bb1b2cb339d50bbcb708b53ad5f4c71c5ea"

	s.T().Run("should send a Tessera private transaction successfully", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}
	
		envelope := fakeEnvelope()
		envelope.SetDataString(data)
		_ = envelope.SetJobType(tx.JobType_ETH_TESSERA_PRIVATE_TX)
	
		encodedKey := base64.StdEncoding.EncodeToString([]byte(enclaveKey))
		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/tessera/%s/storeraw", envelope.GetChainUUID())).
			Reply(200).JSON(rpc.StoreRawResponse{Key: encodedKey})
	
		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusStored)).
			Reply(200).JSON(&txscheduler.JobResponse{})
	
		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
	
		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)
	})
	
	s.T().Run("fail to send a Tessera private transaction", func(t *testing.T) {
		defer gock.Off()
		wg := &multierror.Group{}

		envelope := fakeEnvelope()
		envelope.SetDataString(data)
		_ = envelope.SetJobType(tx.JobType_ETH_TESSERA_PRIVATE_TX)

		gock.New(chainRegistryURL).
			Post(fmt.Sprintf("/tessera/%s/storeraw", envelope.GetChainUUID())).
			Reply(200).BodyString("Invalid_Response")

		gock.New(txSchedulerURL).
			Patch(fmt.Sprintf("/jobs/%s", envelope.GetJobUUID())).
			AddMatcher(txStatusUpdateMatcher(wg, utils.StatusFailed)).
			Reply(200).JSON(&txscheduler.JobResponse{})

		err := s.sendEnvelope(envelope.TxEnvelopeAsRequest())
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		err = waitTimeout(wg, time.Second * 2)
		assert.NoError(t, err)

		retrievedEnvelope, err := s.env.consumer.WaitForEnvelope(envelope.GetID(), s.env.srvConfig.RecoverTopic,
			waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, envelope.GetID(), retrievedEnvelope.GetID())
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Code)
		assert.NotEmpty(t, retrievedEnvelope.GetErrors()[0].Message)
	})
}


func (s *txSignerEthereumTestSuite) TestTxSender_ZHealthCheck() {
	type healthRes struct {
		TransactionScheduler string `json:"transaction-scheduler,omitempty"`
		KeyManager           string `json:"key-manager,omitempty"`
		Kafka                string `json:"kafka,omitempty"`
	}

	httpClient := http.NewClient(http.NewDefaultConfig())
	ctx := s.env.ctx
	s.T().Run("should retrieve positive health check over service dependencies", func(t *testing.T) {
		req, err := http2.NewRequest("GET", fmt.Sprintf("%s/ready?full=1", s.env.metricsURL), nil)
		assert.NoError(s.T(), err)

		gock.New(txSchedulerMetricsURL).Get("/live").Reply(200)
		defer gock.Off()

		resp, err := httpClient.Do(req)
		if err != nil {
			assert.Fail(s.T(), err.Error())
			return
		}

		assert.Equal(s.T(), 200, resp.StatusCode)
		status := healthRes{}
		err = json.UnmarshalBody(resp.Body, &status)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "OK", status.TransactionScheduler)
		assert.Equal(s.T(), "OK", status.Kafka)
	})

	s.T().Run("should retrieve a negative health check over kafka service", func(t *testing.T) {
		req, err := http2.NewRequest("GET", fmt.Sprintf("%s/ready?full=1", s.env.metricsURL), nil)
		assert.NoError(s.T(), err)

		gock.New(txSchedulerMetricsURL).Get("/live").Reply(200)
		defer gock.Off()

		// Kill Kafka on first call so data is added in DB and status is CREATED but does not get updated to STARTED
		err = s.env.client.Stop(ctx, kafkaContainerID)
		assert.NoError(t, err)

		resp, err := httpClient.Do(req)
		if err != nil {
			assert.Fail(s.T(), err.Error())
			return
		}

		err = s.env.client.StartServiceAndWait(ctx, kafkaContainerID, 10*time.Second)
		assert.NoError(t, err)

		assert.Equal(s.T(), 503, resp.StatusCode)
		status := healthRes{}
		err = json.UnmarshalBody(resp.Body, &status)
		assert.NoError(s.T(), err)
		assert.NotEqual(s.T(), "OK", status.Kafka)
		assert.Equal(s.T(), "OK", status.TransactionScheduler)
	})
}

func fakeEnvelope() *tx.Envelope {
	scheduleUUID := uuid.Must(uuid.NewV4()).String()
	jobUUID := uuid.Must(uuid.NewV4()).String()
	chainUUID := uuid.Must(uuid.NewV4()).String()

	envelope := tx.NewEnvelope()
	_ = envelope.SetID(scheduleUUID)
	_ = envelope.SetJobUUID(jobUUID)
	_ = envelope.SetJobType(tx.JobType_ETH_TX)
	_ = envelope.SetNonce(0)
	_ = envelope.SetFromString("0xeca84382E0f1dDdE22EedCd0D803442972EC7BE5")
	_ = envelope.SetGas(21000)
	_ = envelope.SetGasPriceString("10000000")
	_ = envelope.SetValueString("10000000")
	_ = envelope.SetDataString("0x")
	_ = envelope.SetChainIDString("2888")
	_ = envelope.SetChainUUID(chainUUID)
	_ = envelope.SetHeadersValue(utils.TenantIDMetadata, "tenantID")
	_ = envelope.SetPrivateFrom("A1aVtMxLCUHmBVHXoZzzBgPbW/wj5axDpW9X8l91SGo=")
	_ = envelope.SetPrivateFor([]string{"A1aVtMxLCUHmBVHXoZzzBgPbW/wj5axDpW9X8l91SGo=", "B1aVtMxLCUHmBVHXoZzzBgPbW/wj5axDpW9X8l91SGo="})

	return envelope
}

func (s *txSignerEthereumTestSuite) sendEnvelope(protoMessage proto.Message) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = s.env.srvConfig.ListenerTopic

	b, err := encoding.Marshal(protoMessage)
	if err != nil {
		return err
	}
	msg.Value = sarama.ByteEncoder(b)

	_, _, err = s.env.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func waitTimeout(wg *multierror.Group, duration time.Duration) error {
	c := make(chan bool, 1)
	var err error
	go func() {
		defer close(c)
		err = wg.Wait().ErrorOrNil()
	}()

	select {
	case <-c:
		return err
	case <-time.After(duration):
		return fmt.Errorf("timeout after %s", duration.String())
	}
}

func txStatusUpdateMatcher(wg *multierror.Group, status string) gock.MatchFunc {
	cerr := make(chan error, 1)
	wg.Go(func() error {
		return <-cerr
	})

	return func(rw *http2.Request, _ *gock.Request) (bool, error) {
		defer func() {
			cerr <- nil
		}()

		body, _ := ioutil.ReadAll(rw.Body)
		rw.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		req := txscheduler.UpdateJobRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			cerr <- err
			return false, err
		}

		if req.Status != status {
			err := fmt.Errorf("invalid status, got %s, expected %s", req.Status, status)
			cerr <- err
			return false, err
		}

		return true, nil
	}
}

func ethCallMatcher(wg *multierror.Group, method string, args ...interface{}) gock.MatchFunc {
	cerr := make(chan error, 1)
	wg.Go(func() error {
		err := <-cerr
		return err
	})

	return func(rw *http2.Request, grw *gock.Request) (bool, error) {
		defer func() {
			cerr <- nil
		}()

		body, _ := ioutil.ReadAll(rw.Body)
		rw.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		req := utils2.JSONRpcMessage{}
		if err := json.Unmarshal(body, &req); err != nil {
			cerr <- err
			return false, err
		}

		if req.Method != method {
			err := fmt.Errorf("invalid method, got %s, expected %s", req.Method, method)
			cerr <- err
			return false, err
		}

		if len(args) > 0 {
			params, _ := json.Marshal(args)
			if strings.ToLower(string(req.Params)) != strings.ToLower(string(params)) {
				err := fmt.Errorf("invalid params, got %s, expected %s", req.Params, params)
				cerr <- err
				return false, err
			}
		}

		return true, nil
	}
}