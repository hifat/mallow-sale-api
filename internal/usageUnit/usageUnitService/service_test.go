package usageUnitService

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository/mock"
	mockCore "github.com/hifat/mallow-sale-api/pkg/utils/mock/core"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testUsageUnitServiceSuite struct {
	suite.Suite

	mockLogger        *mockCore.Mocklogger
	mockUsageUnitRepo *mockUsageUnitRepository.MockIUsageUnitRepository

	underTest IUsageUnitService
}

func (s *testUsageUnitServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockCore.NewMocklogger(ctrl)
	s.mockUsageUnitRepo = mockUsageUnitRepository.NewMockIUsageUnitRepository(ctrl)

	s.underTest = &usageUnitService{
		logger:        s.mockLogger,
		usageUnitRepo: s.mockUsageUnitRepo,
	}
}

func TestUsageUnitServiceSuite(t *testing.T) {
	suite.Run(t, &testUsageUnitServiceSuite{})
}

func (s *testUsageUnitServiceSuite) TestUsageUnitService_FindIn() {
	s.T().Parallel()

	s.Run("fail - find in codes", func() {
		filter := usageUnit.FilterReq{
			Codes: []string{"mock-code"},
		}
		errFind := errors.New("mock-error")
		s.mockUsageUnitRepo.EXPECT().
			FindInCodes(context.Background(), filter.Codes).
			Return(nil, errFind)

		s.mockLogger.EXPECT().
			Error(errFind)

		res, err := s.underTest.FindIn(context.Background(), filter)
		s.Require().NotNil(err)
		s.Require().Nil(res)
	})

	s.Run("success - find", func() {
		filter := usageUnit.FilterReq{
			Codes: []string{"mock-code"},
		}

		inventories := make([]usageUnit.UsageUnit, 2)
		gofakeit.Slice(&inventories)

		s.mockUsageUnitRepo.EXPECT().
			FindInCodes(context.Background(), filter.Codes).
			Return(inventories, nil)

		res, err := s.underTest.FindIn(context.Background(), filter)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().NotEmpty(res)
	})
}
