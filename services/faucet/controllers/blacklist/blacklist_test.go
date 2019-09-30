package blacklist

import (
	"context"
	"math/big"
	"sync"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/services/faucet/faucet/mock"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/services/faucet/types"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/services/faucet/types/testutils"
)

var (
	chains    = []*big.Int{big.NewInt(10), big.NewInt(2345), big.NewInt(1)}
	addresses = []ethcommon.Address{
		ethcommon.HexToAddress("0xab"),
		ethcommon.HexToAddress("0xcd"),
		ethcommon.HexToAddress("0xef"),
	}
)

func TestBlackList(t *testing.T) {
	// Create Controller and blacklist addresses
	cntrl := NewController()
	for i := range chains {
		cntrl.BlackList(chains[i], addresses[i])
	}
	credit := cntrl.Control(mock.Credit)

	// Prepare test data
	rounds := 600
	tests := make([]*testutils.TestRequest, 0)
	for i := 0; i < rounds; i++ {
		var expectedAmount *big.Int
		var expectErr bool
		if i%2 == 0 {
			expectedAmount = big.NewInt(0)
			expectErr = true
		} else {
			expectedAmount = big.NewInt(10)
			expectErr = false
		}
		tests = append(
			tests,
			&testutils.TestRequest{
				Req: &types.Request{
					ChainID:     chains[i%3],
					Beneficiary: addresses[(i+i%2)%3],
					Amount:      big.NewInt(10),
				},
				ExpectedOK:     i%2 == 1,
				ExpectedAmount: expectedAmount,
				ExpectedErr:    expectErr,
			},
		)
	}

	// Apply tests
	wg := &sync.WaitGroup{}
	for _, test := range tests {
		wg.Add(1)
		go func(test *testutils.TestRequest) {
			defer wg.Done()
			amount, ok, err := credit(context.Background(), test.Req)
			test.ResultAmount, test.ResultOK, test.ResultErr = amount, ok, err
		}(test)
	}
	wg.Wait()

	// Ensure results are correct
	for _, test := range tests {
		testutils.AssertRequest(t, test)
		if test.ResultErr != nil {
			assert.True(t, errors.IsFaucetWarning(test.ResultErr), "%v should be a faucet warning", test.ResultErr)
			assert.Equal(t, "controller.blacklist", errors.FromError(test.ResultErr).GetComponent(), "Error component should be correct")
		}
	}
}