// +build unit

package jobs

import (
	"context"
	"fmt"
	mocks2 "github.com/Shopify/sarama/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/mocks"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/models/testutils"
	"testing"
)

func TestStartJob_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockJobDA := mocks.NewMockJobAgent(ctrl)
	mockLogDA := mocks.NewMockLogAgent(ctrl)
	mockKafkaProducer := mocks2.NewSyncProducer(t, nil)
	txCrafterTopic := "tx-crafter-topic"
	usecase := NewStartJobUseCase(mockJobDA, mockLogDA, mockKafkaProducer, txCrafterTopic)
	ctx := context.Background()

	t.Run("should execute use case successfully", func(t *testing.T) {
		job := testutils.FakeJob(1)
		job.ID = 1
		job.UUID = "6380e2b6-b828-43ee-abdc-de0f8d57dc5f"
		job.Transaction.Sender = "0xfrom"
		job.Schedule = testutils.FakeSchedule()

		mockJobDA.EXPECT().FindOneByUUID(ctx, job.UUID).Return(job, nil)
		mockKafkaProducer.ExpectSendMessageAndSucceed()
		mockLogDA.EXPECT().Insert(ctx, gomock.Any()).Return(nil)

		err := usecase.Execute(ctx, job.UUID)

		assert.Nil(t, err)
	})

	t.Run("should fail with same error if FindOne fails", func(t *testing.T) {
		job := testutils.FakeJob(1)
		job.UUID = "6380e2b6-b828-43ee-abdc-de0f8d57dc5f"
		expectedErr := errors.NotFoundError("error")

		mockJobDA.EXPECT().FindOneByUUID(ctx, job.UUID).Return(nil, expectedErr)

		err := usecase.Execute(ctx, job.UUID)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(startJobComponent), err)
	})

	t.Run("should fail with KafkaConnectionError if Produce fails", func(t *testing.T) {
		job := testutils.FakeJob(1)
		job.UUID = "6380e2b6-b828-43ee-abdc-de0f8d57dc5f"
		job.Transaction.Sender = "0xfrom"
		job.Schedule = testutils.FakeSchedule()

		mockJobDA.EXPECT().FindOneByUUID(ctx, job.UUID).Return(job, nil)
		mockKafkaProducer.ExpectSendMessageAndFail(fmt.Errorf("error"))

		err := usecase.Execute(ctx, job.UUID)
		assert.True(t, errors.IsKafkaConnectionError(err))
	})

	t.Run("should fail with same error if Insert log fails", func(t *testing.T) {
		job := testutils.FakeJob(1)
		job.UUID = "6380e2b6-b828-43ee-abdc-de0f8d57dc5f"
		job.Transaction.Sender = "0xfrom"
		job.Schedule = testutils.FakeSchedule()
		expectedErr := errors.PostgresConnectionError("error")

		mockJobDA.EXPECT().FindOneByUUID(ctx, job.UUID).Return(job, nil)
		mockKafkaProducer.ExpectSendMessageAndSucceed()
		mockLogDA.EXPECT().Insert(ctx, gomock.Any()).Return(expectedErr)

		err := usecase.Execute(ctx, job.UUID)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(startJobComponent), err)
	})
}