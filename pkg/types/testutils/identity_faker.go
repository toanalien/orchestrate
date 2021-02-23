package testutils

import (
	"time"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/multitenancy"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/utils"
)

func FakeAccount() *entities.Account {
	return &entities.Account{
		Alias:               "MyAccount",
		TenantID:            multitenancy.DefaultTenant,
		Attributes:          make(map[string]string),
		Address:             "0x5Cc634233E4a454d47aACd9fC68801482Fb02610",
		PublicKey:           ethcommon.HexToHash(utils.RandHexString(12)).String(),
		CompressedPublicKey: ethcommon.HexToHash(utils.RandHexString(12)).String(),
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}