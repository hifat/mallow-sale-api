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

func (s *testShoppingServiceSuite) SetupTest() {
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

func (s *testShoppingServiceSuite) TestInventoryService_Find() {
	s.T().Parallel()

	s.Run("failed - find shopping", func() {
		ctx := context.Background()

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		mockErr := errors.New("mock-err")
		s.mockShoppingRepo.EXPECT().
			Find(ctx).
			Return(nil, mockErr)

		s.mockLogger.EXPECT().
			Error(mockErr)

		res, err := s.underTest.Find(ctx)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - find shopping return empty slice", func() {
		ctx := context.Background()

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		s.mockShoppingRepo.EXPECT().
			Find(ctx).
			Return(nil, nil)

		res, err := s.underTest.Find(ctx)

		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().IsType([]shoppingModule.Response{}, res.Items)
	})

	s.Run("succeed - find shopping", func() {
		ctx := context.Background()

		req := shoppingModule.Request{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		mockShps := make([]shoppingModule.Response, 3)
		gofakeit.Slice(&mockShps)

		s.mockShoppingRepo.EXPECT().
			Find(ctx).
			Return(mockShps, nil)

		res, err := s.underTest.Find(ctx)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().Equal(len(mockShps), len(res.Items))
		s.Require().Equal(int64(len(mockShps)), res.Meta.Total)

		s.Require().Equal(res.Items, mockShps)
	})
}

func (s *testShoppingServiceSuite) TestInventoryService_Create() {
	s.T().Parallel()

	s.Run("failed - find usage unit by code other error", func() {
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

	s.Run("failed - invalid usage unit code not found", func() {
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
		s.Require().IsType(&handling.ResponseItem[*shoppingModule.Request]{}, res)

		s.Require().NotNil(&req, res.Item)
	})
}

func (s *testShoppingServiceSuite) TestInventoryService_UpdateIsComplete() {
	s.T().Parallel()

	s.Run("failed - find shopping by id other error", func() {
		ctx := context.Background()
		mockShpID := "mock-shp-id"

		mockErr := errors.New("mock-err")
		s.mockShoppingRepo.EXPECT().
			FindByID(ctx, mockShpID).
			Return(nil, mockErr)

		s.mockLogger.EXPECT().
			Error(mockErr)

		mockReq := shoppingModule.ReqUpdateIsComplete{
			IsComplete: true,
		}

		res, err := s.underTest.UpdateIsComplete(ctx, mockShpID, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - find shopping by id error not found", func() {
		ctx := context.Background()
		mockShpID := "mock-shp-id"

		mockErr := define.ErrRecordNotFound
		s.mockShoppingRepo.EXPECT().
			FindByID(ctx, mockShpID).
			Return(nil, mockErr)

		mockReq := shoppingModule.ReqUpdateIsComplete{
			IsComplete: true,
		}

		res, err := s.underTest.UpdateIsComplete(ctx, mockShpID, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeRecordNotFound, resErr.Code)
		s.Require().Equal(define.MsgRecordNotFound, resErr.Message)
		s.Require().Equal(http.StatusNotFound, resErr.Status)
	})

	s.Run("failed - updated shopping is complete", func() {
		ctx := context.Background()
		mockShpID := "mock-shp-id"

		mockShp := shoppingModule.Response{}
		if err := gofakeit.Struct(&mockShp); err != nil {
			s.T().Fatal(err)
		}

		mockShp.ID = mockShpID

		s.mockShoppingRepo.EXPECT().
			FindByID(ctx, mockShpID).
			Return(&mockShp, nil)

		mockReq := shoppingModule.ReqUpdateIsComplete{
			IsComplete: true,
		}

		mockErr := errors.New("mock-err")
		s.mockShoppingRepo.EXPECT().
			UpdateIsComplete(ctx, mockShpID, &mockReq).
			Return(mockErr)

		s.mockLogger.EXPECT().
			Error(mockErr)

		res, err := s.underTest.UpdateIsComplete(ctx, mockShpID, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - updated shopping is complete", func() {
		ctx := context.Background()
		mockShpID := "mock-shp-id"

		mockShp := shoppingModule.Response{}
		if err := gofakeit.Struct(&mockShp); err != nil {
			s.T().Fatal(err)
		}

		mockShp.ID = mockShpID

		s.mockShoppingRepo.EXPECT().
			FindByID(ctx, mockShpID).
			Return(&mockShp, nil)

		mockReq := shoppingModule.ReqUpdateIsComplete{
			IsComplete: true,
		}

		s.mockShoppingRepo.EXPECT().
			UpdateIsComplete(ctx, mockShpID, &mockReq).
			Return(nil)

		res, err := s.underTest.UpdateIsComplete(ctx, mockShpID, &mockReq)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().IsType(&handling.Response{}, res)

		s.Require().Equal(define.CodeUpdated, res.Code)
		s.Require().Equal(define.MsgUpdated, res.Message)
		s.Require().Equal(http.StatusOK, res.Status)
	})
}

func (s *testShoppingServiceSuite) TestInventoryService_DeleteByID() {
	s.T().Parallel()

	s.Run("failed - find shopping by id other error", func() {
		ctx := context.Background()
		mockShpID := "mock-shp-id"

		mockErr := errors.New("mock-err")
		s.mockShoppingRepo.EXPECT().
			FindByID(ctx, mockShpID).
			Return(nil, mockErr)

		s.mockLogger.EXPECT().
			Error(mockErr)

		res, err := s.underTest.DeleteByID(ctx, mockShpID)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - find shopping by id error not found", func() {
		ctx := context.Background()
		mockShpID := "mock-shp-id"

		mockErr := define.ErrRecordNotFound
		s.mockShoppingRepo.EXPECT().
			FindByID(ctx, mockShpID).
			Return(nil, mockErr)

		res, err := s.underTest.DeleteByID(ctx, mockShpID)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeRecordNotFound, resErr.Code)
		s.Require().Equal(define.MsgRecordNotFound, resErr.Message)
		s.Require().Equal(http.StatusNotFound, resErr.Status)
	})

	s.Run("failed - deleted shopping is complete", func() {
		ctx := context.Background()
		mockShpID := "mock-shp-id"

		mockShp := shoppingModule.Response{}
		if err := gofakeit.Struct(&mockShp); err != nil {
			s.T().Fatal(err)
		}

		mockShp.ID = mockShpID

		s.mockShoppingRepo.EXPECT().
			FindByID(ctx, mockShpID).
			Return(&mockShp, nil)

		mockErr := errors.New("mock-err")
		s.mockShoppingRepo.EXPECT().
			DeleteByID(ctx, mockShpID).
			Return(mockErr)

		s.mockLogger.EXPECT().
			Error(mockErr)

		res, err := s.underTest.DeleteByID(ctx, mockShpID)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - deleted shopping is complete", func() {
		ctx := context.Background()
		mockShpID := "mock-shp-id"

		mockShp := shoppingModule.Response{}
		if err := gofakeit.Struct(&mockShp); err != nil {
			s.T().Fatal(err)
		}

		mockShp.ID = mockShpID

		s.mockShoppingRepo.EXPECT().
			FindByID(ctx, mockShpID).
			Return(&mockShp, nil)

		s.mockShoppingRepo.EXPECT().
			DeleteByID(ctx, mockShpID).
			Return(nil)

		res, err := s.underTest.DeleteByID(ctx, mockShpID)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().IsType(&handling.Response{}, res)

		s.Require().Equal(define.CodeDeleted, res.Code)
		s.Require().Equal(define.MsgDeleted, res.Message)
		s.Require().Equal(http.StatusOK, res.Status)
	})
}
