package chains

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mockstore "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/mock"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/models"
)

func TestGetChains_ByUUID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	chainAgent := mockstore.NewMockChainAgent(mockCtrl)

	getChainsUC := NewGetChains(chainAgent)
	chainUUID := uuid.Must(uuid.NewV4()).String()

	filters := make(map[string]string)
	expectedChain := []*models.Chain{
		&models.Chain{
			UUID: chainUUID,
			Name: "testChain",
		},
	}
	chainAgent.EXPECT().GetChains(gomock.Any(), []string{}, gomock.Eq(filters)).Return(expectedChain, nil).Times(1)

	actualChains, err := getChainsUC.Execute(context.Background(), []string{}, filters)
	assert.NoError(t, err)
	assert.Equal(t, expectedChain, actualChains)
}

func TestGetChains_ByUUIDAndTenantID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	chainAgent := mockstore.NewMockChainAgent(mockCtrl)

	getChainsUC := NewGetChains(chainAgent)
	chainUUID := uuid.Must(uuid.NewV4()).String()
	tenantID := "tenantID_4"

	filters := make(map[string]string)
	expectedChain := []*models.Chain{
		&models.Chain{
			UUID:     chainUUID,
			TenantID: tenantID,
			Name:     "testChain",
		},
	}
	chainAgent.EXPECT().GetChains(gomock.Any(), []string{tenantID}, gomock.Eq(filters)).Return(expectedChain, nil).Times(1)

	actualChains, err := getChainsUC.Execute(context.Background(), []string{tenantID}, filters)
	assert.NoError(t, err)
	assert.Equal(t, expectedChain, actualChains)
}