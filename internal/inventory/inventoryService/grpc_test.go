package inventoryService

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/inventoryRepository/mock"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository/mock"
	mockCore "github.com/hifat/mallow-sale-api/pkg/utils/mock/core"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testInventoryGRPCServiceSuite struct {
	suite.Suite

	mockLogger            *mockCore.Mocklogger
	mockInventoryRepo     *mockInventoryRepository.MockIInventoryRepository
	mockUsageUnitGRPCRepo *mockUsageUnitRepository.MockIUsageUnitGRPCRepository

	underTest IInventoryGRPCService
}

func (s *testInventoryGRPCServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockCore.NewMocklogger(ctrl)
	s.mockInventoryRepo = mockInventoryRepository.NewMockIInventoryRepository(ctrl)
	s.mockUsageUnitGRPCRepo = mockUsageUnitRepository.NewMockIUsageUnitGRPCRepository(ctrl)

	s.underTest = &inventoryGRPCService{
		logger:            s.mockLogger,
		inventoryRepo:     s.mockInventoryRepo,
		usageUnitGRPCRepo: s.mockUsageUnitGRPCRepo,
	}
}

func TestInventoryGRPCServiceSuite(t *testing.T) {
	suite.Run(t, &testInventoryGRPCServiceSuite{})
}

func (s *testInventoryGRPCServiceSuite) TestInventoryService_FindIn() {
	s.T().Parallel()

	s.Run("fail - find", func() {
		filter := inventory.FilterReq{
			Codes: []string{"mock-code"},
		}
		errFind := errors.New("mock-error")
		s.mockInventoryRepo.EXPECT().
			FindIn(context.Background(), filter).
			Return(nil, errFind)

		s.mockLogger.EXPECT().
			Error(errFind)

		res, err := s.underTest.FindIn(context.Background(), filter)
		s.Require().NotNil(err)
		s.Require().Nil(res)
	})

	s.Run("success - find", func() {
		filter := inventory.FilterReq{
			Codes: []string{"mock-code"},
		}

		inventories := make([]inventory.Inventory, 2)
		gofakeit.Slice(&inventories)

		s.mockInventoryRepo.EXPECT().
			FindIn(context.Background(), filter).
			Return(inventories, nil)

		res, err := s.underTest.FindIn(context.Background(), filter)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().NotEmpty(res.Items)
	})
}
