package shoppingService

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository/mock"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	mockShoppingRepository "github.com/hifat/mallow-sale-api/internal/shopping/repository/mock"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	mockLogger "github.com/hifat/mallow-sale-api/pkg/logger/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testShoppingServiceSuite struct {
	suite.Suite

	mockLogger        *mockLogger.MockLogger
	mockInventoryRepo *mockInventoryRepository.MockIRepository
	mockShoppingRepo  *mockShoppingRepository.MockIRepository

	underTest IService
}

func (s *testShoppingServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockLogger.NewMockLogger(ctrl)
	s.mockInventoryRepo = mockInventoryRepository.NewMockIRepository(ctrl)
	s.mockShoppingRepo = mockShoppingRepository.NewMockIRepository(ctrl)

	s.underTest = &service{
		logger:        s.mockLogger,
		inventoryRepo: s.mockInventoryRepo,
		shoppingRepo:  s.mockShoppingRepo,
	}
}

func TestInventoryServiceSuite(t *testing.T) {
	suite.Run(t, &testShoppingServiceSuite{})
}

func (s *testShoppingServiceSuite) TestInventoryService_Create() {
	s.T().Parallel()

	s.Run("failed - find inventories in ids failed other error", func() {
		ctx := context.Background()
		mockInvIDs := []string{"1", "2"}
		mockErr := errors.New("mock-err")

		s.mockInventoryRepo.EXPECT().
			FindInIDs(ctx, mockInvIDs).
			Return(nil, mockErr)

		s.mockLogger.EXPECT().
			Error(mockErr)

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		err := s.underTest.Create(&req)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - invalid some inventory ids", func() {
		ctx := context.Background()
		mockInvIDs := []string{"1", "2"}

		mockInvs := make([]inventoryModule.Response, 3)
		gofakeit.Slice(&mockInvs)

		s.mockInventoryRepo.EXPECT().
			FindInIDs(ctx, mockInvIDs).
			Return(mockInvs, nil)

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		err := s.underTest.Create(&req)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInvalidInventoryID, resErr.Code)
		s.Require().Equal(define.MsgInvalidInventoryID, resErr.Message)
		s.Require().Equal(http.StatusBadRequest, resErr.Status)
	})

	s.Run("failed - created shopping", func() {
		ctx := context.Background()
		mockInvIDs := []string{"1", "2", "3"}

		mockInvs := make([]inventoryModule.Response, 3)
		gofakeit.Slice(&mockInvs)

		s.mockInventoryRepo.EXPECT().
			FindInIDs(ctx, mockInvIDs).
			Return(mockInvs, nil)

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		mockErr := errors.New("mock-err")
		s.mockShoppingRepo.EXPECT().
			Create(&req).
			Return(mockErr)

		s.mockLogger.EXPECT().
			Error(mockErr)

		err := s.underTest.Create(&req)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - created shopping", func() {
		ctx := context.Background()
		mockInvIDs := []string{}

		mockInvs := make([]inventoryModule.Response, 3)
		gofakeit.Slice(&mockInvs)

		s.mockInventoryRepo.EXPECT().
			FindInIDs(ctx, mockInvIDs).
			Return(mockInvs, nil)

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		s.mockShoppingRepo.EXPECT().
			Create(&req).
			Return(nil)

		err := s.underTest.Create(&req)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})
}
