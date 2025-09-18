package inventoryService

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository/mock"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository/mock"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
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

func (s *testInventoryServiceSuite) TestInventoryService_Create() {
	s.T().Parallel()

	s.Run("failed - find inventory by name other error", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		ctx := context.Background()

		mockErr := errors.New("mock err")
		s.mockInventoryRepo.EXPECT().
			FindByName(ctx, mockReq.Name).
			Return(nil, mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := s.underTest.Create(ctx, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - duplicated inventory", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockRes := inventoryModule.Response{}
		if err := gofakeit.Struct(&mockRes); err != nil {
			s.T().Fatal(err)
		}

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindByName(ctx, mockReq.Name).
			Return(&mockRes, nil).
			Times(1)

		res, err := s.underTest.Create(ctx, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInventoryNameAlreadyExists, resErr.Code)
		s.Require().Equal(define.MsgInventoryNameAlreadyExists, resErr.Message)
		s.Require().Equal(http.StatusBadRequest, resErr.Status)
	})

	s.Run("failed - find usage unit by code other error", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindByName(ctx, mockReq.Name).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		mockUsageUnitRes := usageUnitModule.Prototype{}
		if err := gofakeit.Struct(&mockUsageUnitRes); err != nil {
			s.T().Fatal(err)
		}

		mockErr := errors.New("mock err")
		s.mockUsageUnitRepo.EXPECT().
			FindByCode(ctx, mockReq.PurchaseUnit.Code).
			Return(nil, mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := s.underTest.Create(ctx, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - invalid usage unit code", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindByName(ctx, mockReq.Name).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		s.mockUsageUnitRepo.EXPECT().
			FindByCode(ctx, mockReq.PurchaseUnit.Code).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		res, err := s.underTest.Create(ctx, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInvalidUsageUnit, resErr.Code)
		s.Require().Equal(define.MsgInvalidUsageUnit, resErr.Message)
		s.Require().Equal(http.StatusBadRequest, resErr.Status)
	})

	s.Run("failed - created inventory", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockUsageUnitRes := usageUnitModule.Prototype{}
		if err := gofakeit.Struct(&mockUsageUnitRes); err != nil {
			s.T().Fatal(err)
		}

		mockReq.PurchaseUnit.Code = mockUsageUnitRes.Code

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindByName(ctx, mockReq.Name).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		s.mockUsageUnitRepo.EXPECT().
			FindByCode(ctx, mockReq.PurchaseUnit.Code).
			Return(&mockUsageUnitRes, nil).
			Times(1)

		mockErr := errors.New("mock err")
		s.mockInventoryRepo.EXPECT().
			Create(ctx, &mockReq).
			Return(mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := s.underTest.Create(ctx, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - created inventory", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockUsageUnitRes := usageUnitModule.Prototype{}
		if err := gofakeit.Struct(&mockUsageUnitRes); err != nil {
			s.T().Fatal(err)
		}

		mockReq.PurchaseUnit.Code = mockUsageUnitRes.Code

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindByName(ctx, mockReq.Name).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		s.mockUsageUnitRepo.EXPECT().
			FindByCode(ctx, mockReq.PurchaseUnit.Code).
			Return(&mockUsageUnitRes, nil).
			Times(1)

		s.mockInventoryRepo.EXPECT().
			Create(ctx, &mockReq).
			Return(nil).
			Times(1)

		res, err := s.underTest.Create(ctx, &mockReq)
		s.Require().NotNil(res)
		s.Require().Nil(err)

		s.Require().Equal(&mockReq, res.Item)
	})
}

func (s *testInventoryServiceSuite) TestInventoryService_Find() {
	s.T().Parallel()

	s.Run("failed - count inventories", func() {
		ctx := context.Background()

		mockErr := errors.New("mock err")
		s.mockInventoryRepo.EXPECT().
			Count(ctx).
			Return(int64(0), mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		mockQuery := utilsModule.QueryReq{}

		res, err := s.underTest.Find(ctx, &mockQuery)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - find inventories", func() {
		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			Count(ctx).
			Return(int64(5), nil).
			Times(1)

		mockQuery := utilsModule.QueryReq{}

		mockErr := errors.New("mock err")
		s.mockInventoryRepo.EXPECT().
			Find(ctx, &mockQuery).
			Return(nil, mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := s.underTest.Find(ctx, &mockQuery)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - return empty slice", func() {
		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			Count(ctx).
			Return(int64(5), nil).
			Times(1)

		mockQuery := utilsModule.QueryReq{}

		s.mockInventoryRepo.EXPECT().
			Find(ctx, &mockQuery).
			Return(nil, nil).
			Times(1)

		res, err := s.underTest.Find(ctx, &mockQuery)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().IsType([]inventoryModule.Response{}, res.Items)
		s.Require().NotNil(res.Meta)

	})

	s.Run("succeed - find inventories", func() {
		ctx := context.Background()
		var total int64 = 3

		s.mockInventoryRepo.EXPECT().
			Count(ctx).
			Return(total, nil).
			Times(1)

		mockQuery := utilsModule.QueryReq{}

		mockInventories := make([]inventoryModule.Response, total)
		gofakeit.Slice(&mockInventories)

		s.mockInventoryRepo.EXPECT().
			Find(ctx, &mockQuery).
			Return(mockInventories, nil).
			Times(1)

		res, err := s.underTest.Find(ctx, &mockQuery)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().IsType([]inventoryModule.Response{}, res.Items)
		s.Require().Equal(total, int64(len(res.Items)))
		s.Require().NotNil(res.Meta)
		s.Require().Equal(total, res.Meta.Total)
	})
}

func (s *testInventoryServiceSuite) TestInventoryService_FindByID() {
	s.T().Parallel()

	s.Run("failed - find inventory by id other error", func() {
		ctx := context.Background()
		mockID := "mock-id"

		mockErr := errors.New("mock err")
		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := s.underTest.FindByID(ctx, mockID)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - find inventory by id not found", func() {
		ctx := context.Background()
		mockID := "mock-id"

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		res, err := s.underTest.FindByID(ctx, mockID)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeRecordNotFound, resErr.Code)
		s.Require().Equal(define.MsgRecordNotFound, resErr.Message)
		s.Require().Equal(http.StatusNotFound, resErr.Status)
	})

	s.Run("succeed - find inventory by id", func() {
		ctx := context.Background()
		mockID := "mock-id"

		mockInventory := inventoryModule.Response{}
		if err := gofakeit.Struct(&mockInventory); err != nil {
			s.T().Fatal(err)
		}

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&mockInventory, nil).
			Times(1)

		res, err := s.underTest.FindByID(ctx, mockID)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().Equal(&mockInventory, res.Item)
	})
}

func (s *testInventoryServiceSuite) TestInventoryService_UpdateByID() {
	s.T().Parallel()

	s.Run("failed - find inventory by id other error", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockID := "mock-id"

		ctx := context.Background()

		mockErr := errors.New("mock err")
		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := s.underTest.UpdateByID(ctx, mockID, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - find inventory by id record not found", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockID := "mock-id"

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		res, err := s.underTest.UpdateByID(ctx, mockID, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeRecordNotFound, resErr.Code)
		s.Require().Equal(define.MsgRecordNotFound, resErr.Message)
		s.Require().Equal(http.StatusNotFound, resErr.Status)
	})

	s.Run("failed - find usage unit by code other error", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockID := "mock-id"

		ctx := context.Background()

		mockInventory := inventoryModule.Response{}
		if err := gofakeit.Struct(&mockInventory); err != nil {
			s.T().Fatal(err)
		}

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&mockInventory, nil).
			Times(1)

		mockErr := errors.New("mock-err")
		s.mockUsageUnitRepo.EXPECT().
			FindByCode(ctx, mockReq.PurchaseUnit.Code).
			Return(nil, mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := s.underTest.UpdateByID(ctx, mockID, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - invalid usage unit code", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockID := "mock-id"

		ctx := context.Background()

		mockInventory := inventoryModule.Response{}
		if err := gofakeit.Struct(&mockInventory); err != nil {
			s.T().Fatal(err)
		}

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&mockInventory, nil).
			Times(1)

		s.mockUsageUnitRepo.EXPECT().
			FindByCode(ctx, mockReq.PurchaseUnit.Code).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		res, err := s.underTest.UpdateByID(ctx, mockID, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInvalidUsageUnit, resErr.Code)
		s.Require().Equal(define.MsgInvalidUsageUnit, resErr.Message)
		s.Require().Equal(http.StatusBadRequest, resErr.Status)
	})

	s.Run("failed - update inventory by id other error", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockID := "mock-id"

		ctx := context.Background()

		mockInventory := inventoryModule.Response{}
		if err := gofakeit.Struct(&mockInventory); err != nil {
			s.T().Fatal(err)
		}

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&mockInventory, nil).
			Times(1)

		mockUsageUnit := usageUnitModule.Prototype{}

		s.mockUsageUnitRepo.EXPECT().
			FindByCode(ctx, mockReq.PurchaseUnit.Code).
			Return(&mockUsageUnit, nil).
			Times(1)

		mockErr := errors.New("mock-err")
		s.mockInventoryRepo.EXPECT().
			UpdateByID(ctx, mockID, &mockReq).
			Return(mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := s.underTest.UpdateByID(ctx, mockID, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - update inventory by id", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockID := "mock-id"

		ctx := context.Background()

		mockInventory := inventoryModule.Response{}
		if err := gofakeit.Struct(&mockInventory); err != nil {
			s.T().Fatal(err)
		}

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&mockInventory, nil).
			Times(1)

		mockUsageUnit := usageUnitModule.Prototype{}

		s.mockUsageUnitRepo.EXPECT().
			FindByCode(ctx, mockReq.PurchaseUnit.Code).
			Return(&mockUsageUnit, nil).
			Times(1)

		s.mockInventoryRepo.EXPECT().
			UpdateByID(ctx, mockID, &mockReq).
			Return(nil).
			Times(1)

		res, err := s.underTest.UpdateByID(ctx, mockID, &mockReq)
		s.Require().NotNil(res)
		s.Require().Nil(err)

		s.Require().Equal(&mockReq, res.Item)
	})
}

func (s *testInventoryServiceSuite) TestInventoryService_DeleteByID() {
	s.Run("failed - find inventory by id other error", func() {
		mockID := "mock-id"

		ctx := context.Background()

		mockErr := errors.New("mock err")
		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		err := s.underTest.DeleteByID(ctx, mockID)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - find inventory by id record not found", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockID := "mock-id"

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		err := s.underTest.DeleteByID(ctx, mockID)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeRecordNotFound, resErr.Code)
		s.Require().Equal(define.MsgRecordNotFound, resErr.Message)
		s.Require().Equal(http.StatusNotFound, resErr.Status)
	})

	s.Run("failed - delete inventory by id other error", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockID := "mock-id"

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&inventoryModule.Response{}, nil).
			Times(1)

		mockErr := errors.New("mock-err")
		s.mockInventoryRepo.EXPECT().
			DeleteByID(ctx, mockID).
			Return(mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		err := s.underTest.DeleteByID(ctx, mockID)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - delete inventory by id", func() {
		mockReq := inventoryModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockID := "mock-id"

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&inventoryModule.Response{}, nil).
			Times(1)

		s.mockInventoryRepo.EXPECT().
			DeleteByID(ctx, mockID).
			Return(nil).
			Times(1)

		err := s.underTest.DeleteByID(ctx, mockID)
		s.Require().Nil(err)
	})
}
