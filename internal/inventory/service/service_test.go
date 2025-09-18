package inventoryService

import (
	"testing"

	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository/mock"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository/mock"
	mockLogger "github.com/hifat/mallow-sale-api/pkg/logger/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testInventoryServiceSuite struct {
	suite.Suite

	mockLogger        *mockLogger.MockLogger
	mockInventoryRepo *mockInventoryRepository.MockIRepository
	mockUsageUnitRepo *mockUsageUnitRepository.MockRepository

	underTest IService
}

func (s *testInventoryServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockLogger.NewMockLogger(ctrl)
	s.mockInventoryRepo = mockInventoryRepository.NewMockIRepository(ctrl)
	s.mockUsageUnitRepo = mockUsageUnitRepository.NewMockRepository(ctrl)

	s.underTest = &service{
		logger:        s.mockLogger,
		inventoryRepo: s.mockInventoryRepo,
		usageUnitRepo: s.mockUsageUnitRepo,
	}
}

func TestInventoryServiceSuite(t *testing.T) {
	suite.Run(t, &testInventoryServiceSuite{})
}
