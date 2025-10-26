package usageUnitHelper

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository/mock"
	mockLogger "github.com/hifat/mallow-sale-api/pkg/logger/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testUsageUnitHelperSuite struct {
	suite.Suite

	mockLogger        *mockLogger.MockLogger
	mockUsageUnitRepo *mockUsageUnitRepository.MockIRepository

	underTest IHelper
}

func (s *testUsageUnitHelperSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockLogger.NewMockLogger(ctrl)
	s.mockUsageUnitRepo = mockUsageUnitRepository.NewMockIRepository(ctrl)

	s.underTest = &helper{
		logger:        s.mockLogger,
		usageUnitRepo: s.mockUsageUnitRepo,
	}
}

func TestUsageUnitHelperSuite(t *testing.T) {
	suite.Run(t, &testUsageUnitHelperSuite{})
}

func (s *testUsageUnitHelperSuite) TestUsageUnitHelper_GetNameByCode() {
	s.T().Parallel()

	s.Run("failed - find usage unit by code error", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		ctx := context.Background()
		mockCodes := []string{"mock-id"}

		mockErr := errors.New("mock err")
		s.mockUsageUnitRepo.EXPECT().
			FindInCodes(ctx, mockCodes).
			Return(nil, mockErr).
			Times(1)

		getNameByCode, err := s.underTest.GetNameByCode(ctx, mockCodes)
		s.Require().Nil(getNameByCode)
		s.Require().NotNil(err)
		s.Require().ErrorIs(mockErr, err)
	})

	s.Run("failed - get usage unit name by code not found", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		ctx := context.Background()

		mockUsgUnits := make([]usageUnitModule.Prototype, 3)
		gofakeit.Slice(&mockUsgUnits)
		limitUsgUnits := len(mockUsgUnits) - 1

		mockCodes := make([]string, 0, len(mockUsgUnits))
		for i, v := range mockUsgUnits {
			if i < limitUsgUnits {
				break
			}

			mockCodes = append(mockCodes, v.Code)
		}

		s.mockUsageUnitRepo.EXPECT().
			FindInCodes(ctx, mockCodes).
			Return(mockUsgUnits[:limitUsgUnits], nil).
			Times(1)

		getNameByCode, err := s.underTest.GetNameByCode(ctx, mockCodes)
		s.Require().Nil(err)
		s.Require().NotNil(getNameByCode)

		usgUnitName := getNameByCode(mockUsgUnits[limitUsgUnits].Code)
		s.Require().Equal("", usgUnitName)
	})

	s.Run("succeed - get usage unit name by code", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		ctx := context.Background()

		mockUsgUnits := make([]usageUnitModule.Prototype, 3)
		gofakeit.Slice(&mockUsgUnits)

		mockCodes := make([]string, 0, len(mockUsgUnits))
		for _, v := range mockUsgUnits {
			mockCodes = append(mockCodes, v.Code)
		}

		s.mockUsageUnitRepo.EXPECT().
			FindInCodes(ctx, mockCodes).
			Return(mockUsgUnits, nil).
			Times(1)

		getNameByCode, err := s.underTest.GetNameByCode(ctx, mockCodes)
		s.Require().Nil(err)
		s.Require().NotNil(getNameByCode)

		usgUnitName := getNameByCode(mockUsgUnits[1].Code)
		s.Require().Equal(mockUsgUnits[1].Name, usgUnitName)
	})
}
