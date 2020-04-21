package usecases

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/store"
)

const getMethodsComponent = component + ".get-methods"

//go:generate mockgen -source=get_methods.go -destination=mocks/mock_get_methods.go -package=mocks

type GetMethodsUseCase interface {
	Execute(ctx context.Context, account *common.AccountInstance, selector []byte) (abi string, methodsABI []string, err error)
}

// GetMethods is a use case to get methods
type GetMethods struct {
	methodDataAgent store.MethodDataAgent
}

// NewGetMethods creates a new GetMethods
func NewGetMethods(methodDataAgent store.MethodDataAgent) *GetMethods {
	return &GetMethods{
		methodDataAgent: methodDataAgent,
	}
}

// Execute validates and registers a new contract in DB
func (usecase *GetMethods) Execute(ctx context.Context, account *common.AccountInstance, selector []byte) (abi string, methodsABI []string, err error) {
	method, err := usecase.methodDataAgent.FindOneByAccountAndSelector(ctx, account, selector)
	if errors.IsConnectionError(err) {
		return "", nil, errors.FromError(err).ExtendComponent(getMethodsComponent)
	}
	if method != nil {
		return method.ABI, nil, nil
	}

	defaultMethods, err := usecase.methodDataAgent.FindDefaultBySelector(ctx, selector)
	if err != nil {
		return "", nil, errors.FromError(err).ExtendComponent(getMethodsComponent)
	}

	for _, m := range defaultMethods {
		methodsABI = append(methodsABI, m.ABI)
	}

	return "", methodsABI, nil
}