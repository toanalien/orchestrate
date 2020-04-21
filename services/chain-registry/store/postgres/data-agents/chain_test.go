// +build unit
// +build !race
// +build !integration

package dataagents

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	pgTestUtils "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/database/postgres/testutils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/models"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/postgres/migrations"
)

type ChainModelsTestSuite struct {
	pg *pgTestUtils.PGTestHelper
	ChainTestSuite
}

func TestModelsChain(t *testing.T) {
	s := new(ChainModelsTestSuite)
	s.pg = pgTestUtils.NewPGTestHelper(migrations.Collection)
	suite.Run(t, s)
}

func (s *ChainModelsTestSuite) SetupSuite() {
	s.pg.InitTestDB(s.T())
	s.ChainAgent = NewPGChainAgent(s.pg.DB)
}

func (s *ChainModelsTestSuite) SetupTest() {
	s.pg.Upgrade(s.T())
}

func (s *ChainModelsTestSuite) TearDownTest() {
	s.pg.Downgrade(s.T())
}

func (s *ChainModelsTestSuite) TearDownSuite() {
	s.pg.DropTestDB(s.T())
}


// ChainTestSuite is a test suite for ChainRegistry
type ChainTestSuite struct {
	suite.Suite
	ChainAgent *PGChainAgent
}

const (
	chainName1    = "testChain1"
	chainName2    = "testChain2"
	chainName3    = "testChain3"
	tenantID1     = "tenantID1"
	tenantID2     = "tenantID2"
	errorTenantID = "errorTenantID"
)

var tenantID1Chains = map[string]*models.Chain{
	chainName1: {
		Name:                    chainName1,
		TenantID:                tenantID1,
		URLs:                    []string{"http://testurlone.com", "http://testurltwo.com"},
		ListenerDepth:           &(&struct{ x uint64 }{1}).x,
		ListenerCurrentBlock:    &(&struct{ x uint64 }{1}).x,
		ListenerStartingBlock:   &(&struct{ x uint64 }{1}).x,
		ListenerBackOffDuration: &(&struct{ x string }{"1s"}).x,
	},
	chainName2: {
		Name:                    chainName2,
		TenantID:                tenantID1,
		URLs:                    []string{"http://localhost:8545", "https://localhost:443"},
		ListenerDepth:           &(&struct{ x uint64 }{2}).x,
		ListenerCurrentBlock:    &(&struct{ x uint64 }{2}).x,
		ListenerStartingBlock:   &(&struct{ x uint64 }{2}).x,
		ListenerBackOffDuration: &(&struct{ x string }{"2s"}).x,
	},
}
var tenantID2Chains = map[string]*models.Chain{
	chainName1: {
		Name:                      chainName1,
		TenantID:                  tenantID2,
		URLs:                      []string{"http://testurlone.com", "http://testurltwo.com"},
		ListenerDepth:             &(&struct{ x uint64 }{1}).x,
		ListenerCurrentBlock:      &(&struct{ x uint64 }{1}).x,
		ListenerStartingBlock:     &(&struct{ x uint64 }{1}).x,
		ListenerBackOffDuration:   &(&struct{ x string }{"1s"}).x,
		ListenerExternalTxEnabled: &(&struct{ x bool }{true}).x,
	},
	chainName2: {
		Name:                      chainName2,
		TenantID:                  tenantID2,
		URLs:                      []string{"http://testurlone.com", "http://testurltwo.com"},
		ListenerDepth:             &(&struct{ x uint64 }{2}).x,
		ListenerCurrentBlock:      &(&struct{ x uint64 }{2}).x,
		ListenerStartingBlock:     &(&struct{ x uint64 }{2}).x,
		ListenerBackOffDuration:   &(&struct{ x string }{"2s"}).x,
		ListenerExternalTxEnabled: &(&struct{ x bool }{true}).x,
	},
	chainName3: {
		Name:                      chainName3,
		TenantID:                  tenantID2,
		URLs:                      []string{"http://testurlone.com", "http://testurltwo.com"},
		ListenerDepth:             &(&struct{ x uint64 }{3}).x,
		ListenerCurrentBlock:      &(&struct{ x uint64 }{3}).x,
		ListenerStartingBlock:     &(&struct{ x uint64 }{3}).x,
		ListenerBackOffDuration:   &(&struct{ x string }{"3s"}).x,
		ListenerExternalTxEnabled: &(&struct{ x bool }{true}).x,
	},
}

var ChainsSample = map[string]map[string]*models.Chain{
	tenantID1: tenantID1Chains,
	tenantID2: tenantID2Chains,
}

func CompareChains(t *testing.T, chain1, chain2 *models.Chain) {
	assert.Equal(t, chain1.Name, chain2.Name, "Should get the same chain name")
	assert.Equal(t, chain1.TenantID, chain2.TenantID, "Should get the same chain tenantID")
	assert.Equal(t, chain1.URLs, chain2.URLs, "Should get the same chain URLs")
	assert.Equal(t, chain1.ListenerDepth, chain2.ListenerDepth, "Should get the same chain ListenerDepth")
	assert.Equal(t, chain1.ListenerCurrentBlock, chain2.ListenerCurrentBlock, "Should get the same chain")
	assert.Equal(t, chain1.ListenerStartingBlock, chain2.ListenerStartingBlock, "Should get the same chain ListenerBlockPosition")
	assert.Equal(t, chain1.ListenerBackOffDuration, chain2.ListenerBackOffDuration, "Should get the same chain ListenerBackOffDuration")
	assert.Equal(t, chain1.ListenerExternalTxEnabled, chain2.ListenerExternalTxEnabled, "Should get the same chain ListenerExternalTxEnabled")
}

func (s *ChainTestSuite) TestRegisterChain() {
	err := s.ChainAgent.RegisterChain(context.Background(), ChainsSample[tenantID1][chainName1])
	assert.NoError(s.T(), err, "Should register chain properly")

	err = s.ChainAgent.RegisterChain(context.Background(), ChainsSample[tenantID1][chainName1])
	assert.Error(s.T(), err, "Should get an error violating the 'unique' constraint")
}

func (s *ChainTestSuite) TestRegisterChainWithMissingURLsFieldErr() {
	listenerDepth := uint64(2)
	listenerCurrentBlock := uint64(2)
	listenerStartingBlock := uint64(2)
	listenerBackOffDuration := "2s"
	chainError := &models.Chain{
		Name:                    "chainNameErr",
		TenantID:                "tenantID1",
		ListenerDepth:           &listenerDepth,
		ListenerCurrentBlock:    &listenerCurrentBlock,
		ListenerStartingBlock:   &listenerStartingBlock,
		ListenerBackOffDuration: &listenerBackOffDuration,
	}
	chainError.SetDefault()
	err := s.ChainAgent.RegisterChain(context.Background(), chainError)
	assert.Error(s.T(), err, "Should get an error when a field is missing URLs")
	assert.True(s.T(), errors.IsDataError(err), "Should be a DataError")
}

func (s *ChainTestSuite) TestRegisterChainWithInvalidBackOffDurationErr() {
	listenerDepth := uint64(2)
	listenerCurrentBlock := uint64(2)
	listenerStartingBlock := uint64(2)
	listenerBackOffDuration := "2000"
	chainError := &models.Chain{
		Name:                    "chainNameErr",
		TenantID:                "tenantID1",
		URLs:                    []string{"http://testurlthree.com"},
		ListenerDepth:           &listenerDepth,
		ListenerCurrentBlock:    &listenerCurrentBlock,
		ListenerStartingBlock:   &listenerStartingBlock,
		ListenerBackOffDuration: &listenerBackOffDuration,
	}
	chainError.SetDefault()
	err := s.ChainAgent.RegisterChain(context.Background(), chainError)
	assert.Error(s.T(), err, "Should get an error when invalid backOffDuration")
	assert.True(s.T(), errors.IsDataError(err), "Should be a DataError")
}

func (s *ChainTestSuite) TestRegisterChainWithInvalidUrlsErr() {
	listenerDepth := uint64(2)
	listenerCurrentBlock := uint64(2)
	listenerStartingBlock := uint64(2)
	listenerBackOffDuration := "2s"
	chainError := &models.Chain{
		Name:                    "chainNameErr",
		TenantID:                "tenantID1",
		URLs:                    []string{"%!1231"},
		ListenerDepth:           &listenerDepth,
		ListenerCurrentBlock:    &listenerCurrentBlock,
		ListenerStartingBlock:   &listenerStartingBlock,
		ListenerBackOffDuration: &listenerBackOffDuration,
	}
	chainError.SetDefault()
	err := s.ChainAgent.RegisterChain(context.Background(), chainError)
	assert.Error(s.T(), err, "Should get an error when invalid URI")
	assert.True(s.T(), errors.IsDataError(err), "Should be a DataError")
}

func (s *ChainTestSuite) TestRegisterChains() {
	for _, chains := range ChainsSample {
		for _, chain := range chains {
			chain.SetDefault()
			err := s.ChainAgent.RegisterChain(context.Background(), chain)
			assert.NoError(s.T(), err, "Should not fail to register")
		}
	}
}

func (s *ChainTestSuite) TestGetChains() {
	s.TestRegisterChains()

	chains, err := s.ChainAgent.GetChains(context.Background(), nil)
	assert.NoError(s.T(), err, "Should get chains without errors")
	assert.Len(s.T(), chains, len(tenantID1Chains)+len(tenantID2Chains), "Should get the same number of chains")

	for _, chain := range chains {
		CompareChains(s.T(), chain, ChainsSample[chain.TenantID][chain.Name])
	}
}

func (s *ChainTestSuite) TestGetChainsByTenant() {
	s.TestRegisterChains()

	chains, err := s.ChainAgent.GetChainsByTenant(context.Background(), nil, tenantID1)
	assert.NoError(s.T(), err, "Should get chains without errors")
	assert.Len(s.T(), chains, len(tenantID1Chains), "Should get the same number of chains for tenantID1")

	for _, chain := range chains {
		assert.Equal(s.T(), tenantID1, chain.TenantID)
	}
}

func (s *ChainTestSuite) TestGetChainByUUID() {
	s.TestRegisterChains()

	chainUUID := ChainsSample[tenantID1][chainName1].UUID

	chain, err := s.ChainAgent.GetChainByUUID(context.Background(), chainUUID)
	assert.NoError(s.T(), err, "Should get chain without errors")

	CompareChains(s.T(), chain, ChainsSample[tenantID1][chainName1])
}

func (s *ChainTestSuite) TestGetChainByUUIDByTenant() {
	s.TestRegisterChains()

	chainUUID := ChainsSample[tenantID1][chainName1].UUID

	chain, err := s.ChainAgent.GetChainByUUIDAndTenant(context.Background(), chainUUID, tenantID1)
	assert.NoError(s.T(), err, "Should get chain without errors")

	assert.Equal(s.T(), tenantID1, chain.TenantID)
}

func (s *ChainTestSuite) TestUpdateChainByUUID() {
	s.TestRegisterChains()

	testChain := ChainsSample[tenantID1][chainName2]
	testChain.URLs = []string{"http://testurlone.com"}
	err := s.ChainAgent.UpdateChainByUUID(context.Background(), testChain.UUID, testChain)
	assert.NoError(s.T(), err, "Should update chain without errors")

	chain, _ := s.ChainAgent.GetChainByUUID(context.Background(), testChain.UUID)
	CompareChains(s.T(), chain, testChain)
}

func (s *ChainTestSuite) TestNotFoundTenantErrorUpdateChainByName() {
	testChain := ChainsSample[tenantID1][chainName2]
	testChain.URLs = []string{"http://testurlone.com"}
	err := s.ChainAgent.UpdateChainByName(context.Background(), testChain.Name, testChain)
	assert.Error(s.T(), err, "Should get chain without errors")
}

func (s *ChainTestSuite) TestNotFoundNameErrorUpdateChainByName() {
	s.TestRegisterChains()

	testChain := &models.Chain{
		Name:     tenantID1,
		TenantID: errorTenantID,
		URLs:     []string{"http://testurlone.com"},
	}
	err := s.ChainAgent.UpdateChainByName(context.Background(), testChain.Name, testChain)
	assert.Error(s.T(), err, "Should get chain without errors")
}

func (s *ChainTestSuite) TestErrorNotFoundUpdateChainByUUID() {
	s.TestRegisterChains()

	testChain := &models.Chain{
		UUID: "0d60a85e-0b90-4482-a14c-108aea2557aa",
		URLs: []string{"http://testurlone.com"},
	}
	err := s.ChainAgent.UpdateChainByUUID(context.Background(), testChain.UUID, testChain)
	assert.Error(s.T(), err, "Should update chain with errors")
}

func (s *ChainTestSuite) TestDeleteChainByUUID() {
	s.TestRegisterChains()

	chainUUID := ChainsSample[tenantID1][chainName1].UUID

	err := s.ChainAgent.DeleteChainByUUID(context.Background(), chainUUID)
	assert.NoError(s.T(), err, "Should delete chain without errors")
}

func (s *ChainTestSuite) TestDeleteChainByUUIDByTenant() {
	s.TestRegisterChains()

	chainUUID := ChainsSample[tenantID1][chainName1].UUID

	err := s.ChainAgent.DeleteChainByUUIDAndTenant(context.Background(), chainUUID, tenantID1)
	assert.NoError(s.T(), err, "Should delete chain without errors")
}

func (s *ChainTestSuite) TestErrorNotFoundDeleteChainByUUIDAndTenant() {
	s.TestRegisterChains()

	// tenantID2 in the context but we try to delete the chainUUID of tenantID1
	chainUUID := ChainsSample[tenantID1][chainName1].UUID

	err := s.ChainAgent.DeleteChainByUUIDAndTenant(context.Background(), chainUUID, tenantID2)
	assert.Error(s.T(), err, "Should delete chain with errors")
}

func (s *ChainTestSuite) TestErrorNotFoundDeleteChainByUUID() {
	s.TestRegisterChains()

	err := s.ChainAgent.DeleteChainByUUID(context.Background(), "0d60a85e-0b90-4482-a14c-108aea2557aa")
	assert.Error(s.T(), err, "Should delete chain with errors")
}