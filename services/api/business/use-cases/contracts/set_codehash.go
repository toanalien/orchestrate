package contracts

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/log"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/use-cases"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store"
	models2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/models"
)

const setCodeHashComponent = "use-cases.set-codehash"

type setCodeHashUseCase struct {
	agent  store.CodeHashAgent
	logger *log.Logger
}

func NewSetCodeHashUseCase(agent store.CodeHashAgent) usecases.SetContractCodeHashUseCase {
	return &setCodeHashUseCase{
		agent:  agent,
		logger: log.NewLogger().SetComponent(setCodeHashComponent),
	}
}

func (uc *setCodeHashUseCase) Execute(ctx context.Context, chainID, address, codeHash string) error {
	ctx = log.WithFields(ctx, log.Field("chain_id", chainID), log.Field("address", chainID))
	logger := uc.logger.WithContext(ctx)
	logger.Debug("setting code-hash is starting ...")

	codehash := &models2.CodehashModel{
		ChainID:  chainID,
		Address:  address,
		Codehash: codeHash,
	}

	err := uc.agent.Insert(ctx, codehash)
	if err != nil {
		return errors.FromError(err).ExtendComponent(setCodeHashComponent)
	}

	logger.Debug("code-hash updated successfully")
	return nil
}