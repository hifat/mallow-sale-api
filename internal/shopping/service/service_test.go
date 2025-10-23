package shoppingService

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	mockShoppingRepository "github.com/hifat/mallow-sale-api/internal/shopping/repository/mock"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository/mock"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	mockLogger "github.com/hifat/mallow-sale-api/pkg/logger/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testShoppingServiceSuite struct {
	suite.Suite

	mockLogger       *mockLogger.MockLogger
	mockShoppingRepo *mockShoppingRepository.MockIRepository
	mockUsageUniRepo *mockUsageUnitRepository.MockIRepository

	underTest IService
}

func (s *testShoppingServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockLogger.NewMockLogger(ctrl)
	s.mockShoppingRepo = mockShoppingRepository.NewMockIRepository(ctrl)
	s.mockUsageUniRepo = mockUsageUnitRepository.NewMockIRepository(ctrl)

	s.underTest = &service{
		logger:        s.mockLogger,
		shoppingRepo:  s.mockShoppingRepo,
		usageUnitRepo: s.mockUsageUniRepo,
	}
}

func TestInventoryServiceSuite(t *testing.T) {
	suite.Run(t, &testShoppingServiceSuite{})
}

func (s *testShoppingServiceSuite) TestInventoryService_Create() {
	s.T().Parallel()

	s.Run("failed - find usage unit by code", func() {
		ctx := context.Background()

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		mockErr := errors.New("mock-err")
		s.mockUsageUniRepo.EXPECT().
			FindByCode(ctx, req.PurchaseUnit.Code).
			Return(nil, mockErr)

		s.mockLogger.EXPECT().
			Error(mockErr)

		res, err := s.underTest.Create(ctx, &req)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - invalid usage unit code", func() {
		ctx := context.Background()

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		mockUsgUnit := usageUnitModule.Prototype{}
		if err := gofakeit.Struct(&mockUsgUnit); err != nil {
			s.T().Fatal(err)
		}

		mockErr := define.ErrRecordNotFound
		s.mockUsageUniRepo.EXPECT().
			FindByCode(ctx, req.PurchaseUnit.Code).
			Return(nil, mockErr)

		res, err := s.underTest.Create(ctx, &req)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInvalidUsageUnit, resErr.Code)
		s.Require().Equal(define.MsgInvalidUsageUnit, resErr.Message)
		s.Require().Equal(http.StatusBadRequest, resErr.Status)
	})

	s.Run("failed - created shopping", func() {
		ctx := context.Background()

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		mockUsgUnit := usageUnitModule.Prototype{}
		if err := gofakeit.Struct(&mockUsgUnit); err != nil {
			s.T().Fatal(err)
		}

		s.mockUsageUniRepo.EXPECT().
			FindByCode(ctx, req.PurchaseUnit.Code).
			Return(&mockUsgUnit, nil)

		mockErr := errors.New("mock-err")
		s.mockShoppingRepo.EXPECT().
			Create(ctx, &req).
			Return(mockErr)

		s.mockLogger.EXPECT().
			Error(mockErr)

		res, err := s.underTest.Create(ctx, &req)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - created shopping", func() {
		ctx := context.Background()

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		mockUsgUnit := usageUnitModule.Prototype{}
		if err := gofakeit.Struct(&mockUsgUnit); err != nil {
			s.T().Fatal(err)
		}

		s.mockUsageUniRepo.EXPECT().
			FindByCode(ctx, req.PurchaseUnit.Code).
			Return(&mockUsgUnit, nil)

		s.mockShoppingRepo.EXPECT().
			Create(ctx, &req).
			Return(nil)

		res, err := s.underTest.Create(ctx, &req)
		s.Require().Nil(err)
		s.Require().IsType(&handling.Response{}, res)

		s.Require().Equal(define.CodeCreated, res.Code)
		s.Require().Equal(define.MsgCreated, res.Message)
		s.Require().Equal(http.StatusCreated, res.Status)
	})
}
