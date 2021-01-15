package proxy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/encoding/json"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	ethclient "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/ethclient/utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/utils"
)

var rpcCachedMethods = map[string]bool{
	"eth_getBlockByNumber":      true,
	"eth_getTransactionReceipt": true,
}

func HTTPCacheRequest(req *http.Request) (c bool, k string, ttl time.Duration, err error) {
	logger := log.WithContext(req.Context())
	if req.Method != "POST" {
		return false, "", 0, nil
	}

	if req.Body == nil {
		return false, "", 0, nil
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return false, "", 0, errors.InternalError("can't read request body: %q", err)
	}

	// And now set a new body, which will simulate the same data we read
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var msg ethclient.JSONRpcMessage
	err = json.Unmarshal(body, &msg)
	// In case request does not correspond to one of expected call RPC call, we ignore
	if err != nil {
		logger.Debugf("HTTPCache: request is not an RPC message")
		return false, "", 0, nil
	}

	if _, ok := rpcCachedMethods[msg.Method]; !ok {
		logger.Debugf("HTTPCache: RPC method is ignored: %s", msg.Method)
		return false, "", 0, nil
	}

	cacheKey := fmt.Sprintf("%s(%s)", msg.Method, string(msg.Params))
	if msg.Method == "eth_getBlockByNumber" && strings.Contains(string(msg.Params), "latest") {
		return true, cacheKey, time.Second, nil
	}

	return true, cacheKey, 0, nil
}

func HTTPCacheResponse(resp *http.Response) bool {
	var msg ethclient.JSONRpcMessage
	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}
	err := json.UnmarshalBody(reader, &msg)
	if err != nil {
		log.WithError(err).Debugf("HTTPCache: cannot decode response")
		return false
	}

	if msg.Error != nil {
		log.WithField("error", msg.Error.Message).Debugf("HTTPCache: skip RPC error responses")
		return false
	}

	if len(msg.Result) == 0 {
		log.Debugf("HTTPCache: skip RPC empty response results")
		return false
	}

	return true
}

func httpCacheGenerateChainKey(chain *entities.Chain) string {
	// Order urls to identify common chain definitions
	sort.Sort(utils.Alphabetic(chain.URLs))

	var urls = strings.Join(chain.URLs, "_")

	if chain.PrivateTxManager != nil {
		urls = urls + "_" + chain.PrivateTxManager.URL
	}

	return urls
}