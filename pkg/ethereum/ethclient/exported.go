package ethclient

import (
	"context"
	"sync"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethereum/ethclient/rpc"
)

var (
	client   Client
	initOnce = &sync.Once{}
)

func Init(ctx context.Context) {
	initOnce.Do(func() {
		if client != nil {
			return
		}

		// Initialize RPC client
		rpc.Init(ctx)

		client = rpc.GlobalClient()
	})
}

// GlobalClient returns global Client
func GlobalClient() Client {
	return client
}

// SetGlobalClient sets global Client
func SetGlobalClient(ec Client) {
	client = ec
}