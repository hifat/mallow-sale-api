package recipeService

import (
	"testing"

	mockInventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper/mock"
	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository/mock"
	mockRecipeHelper "github.com/hifat/mallow-sale-api/internal/recipe/helper/mock"
	mockRecipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/repository/mock"
	mockUsageUnitHelper "github.com/hifat/mallow-sale-api/internal/usageUnit/helper/mock"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository/mock"
	mockLogger "github.com/hifat/mallow-sale-api/pkg/logger/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testRecipeServiceSuite struct {
	suite.Suite

	mockLogger           *mockLogger.MockLogger
	mockRecipeRepository *mockRecipeRepository.MockIRepository
	mockInventoryRepo    *mockInventoryRepository.MockIRepository
	mockUsageUnitRepo    *mockUsageUnitRepository.MockIRepository
	mockUsageUnitHelper  *mockUsageUnitHelper.MockIHelper
	mockInventoryHelper  *mockInventoryHelper.MockIHelper
	mockRecipeTypeHelper *mockRecipeHelper.MockIRecipeTypeHelper

	underTest IService
}

func (s *testRecipeServiceSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockLogger.NewMockLogger(ctrl)
	s.mockRecipeRepository = mockRecipeRepository.NewMockIRepository(ctrl)
	s.mockInventoryRepo = mockInventoryRepository.NewMockIRepository(ctrl)
	s.mockUsageUnitRepo = mockUsageUnitRepository.NewMockIRepository(ctrl)
	s.mockUsageUnitHelper = mockUsageUnitHelper.NewMockIHelper(ctrl)
	s.mockInventoryHelper = mockInventoryHelper.NewMockIHelper(ctrl)
	s.mockRecipeTypeHelper = mockRecipeHelper.NewMockIRecipeTypeHelper(ctrl)

	s.underTest = &service{
		logger:           s.mockLogger,
		recipeRepo:       s.mockRecipeRepository,
		inventoryRepo:    s.mockInventoryRepo,
		usageUnitRepo:    s.mockUsageUnitRepo,
		usageUnitHelper:  s.mockUsageUnitHelper,
		inventoryHelper:  s.mockInventoryHelper,
		recipeTypeHelper: s.mockRecipeTypeHelper,
	}
}

func TestRecipeServiceSuite(t *testing.T) {
	suite.Run(t, &testRecipeServiceSuite{})
}

// func (s *testRecipeServiceSuite) TestInventoryService_Create() {
// 	s.T().Parallel()

// 	s.Run("failed - find inventory by name other error", func() {
// 		mockReq := inventoryModule.Request{}
// 		if err := gofakeit.Struct(&mockReq); err != nil {
// 			s.T().Fatal(err)
// 		}

// 		ctx := context.Background()

// 		mockErr := errors.New("mock err")
// 		s.mockInventoryRepo.EXPECT().
// 			FindByName(ctx, mockReq.Name).
// 			Return(nil, mockErr).
// 			Times(1)

// 		s.mockLogger.EXPECT().
// 			Error(mockErr).
// 			Times(1)

// 		s.mockUsageUnitRepo.EXPECT().
// 			FindByCode(gomock.Any(), gomock.Any()).
// 			Return(&usageUnitModule.Prototype{}, nil).
// 			Times(1)

// 		res, err := s.underTest.Create(ctx, &mockReq)
// 		s.Require().Nil(res)
// 		s.Require().NotNil(err)
// 		s.Require().IsType(handling.ErrorResponse{}, err)

// 		resErr := err.(handling.ErrorResponse)

// 		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
// 		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
// 		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
// 	})

// 	s.Run("failed - duplicated inventory", func() {
// 		mockReq := inventoryModule.Request{}
// 		if err := gofakeit.Struct(&mockReq); err != nil {
// 			s.T().Fatal(err)
// 		}

// 		mockRes := inventoryModule.Response{}
// 		if err := gofakeit.Struct(&mockRes); err != nil {
// 			s.T().Fatal(err)
// 		}

// 		ctx := context.Background()

// 		s.mockInventoryRepo.EXPECT().
// 			FindByName(ctx, mockReq.Name).
// 			Return(&mockRes, nil).
// 			Times(1)

// 		s.mockUsageUnitRepo.EXPECT().
// 			FindByCode(gomock.Any(), gomock.Any()).
// 			Return(&usageUnitModule.Prototype{}, nil).
// 			Times(1)

// 		res, err := s.underTest.Create(ctx, &mockReq)
// 		s.Require().Nil(res)
// 		s.Require().NotNil(err)
// 		s.Require().IsType(handling.ErrorResponse{}, err)

// 		resErr := err.(handling.ErrorResponse)

// 		s.Require().Equal(define.CodeDuplicatedInventoryName, resErr.Code)
// 		s.Require().Equal(define.MsgDuplicatedInventoryName, resErr.Message)
// 		s.Require().Equal(http.StatusBadRequest, resErr.Status)
// 	})

// 	s.Run("failed - find usage unit by code other error", func() {
// 		mockReq := inventoryModule.Request{}
// 		if err := gofakeit.Struct(&mockReq); err != nil {
// 			s.T().Fatal(err)
// 		}

// 		ctx := context.Background()

// 		s.mockInventoryRepo.EXPECT().
// 			FindByName(ctx, mockReq.Name).
// 			Return(nil, define.ErrRecordNotFound).
// 			Times(1)

// 		mockErr := errors.New("mock err")
// 		s.mockUsageUnitRepo.EXPECT().
// 			FindByCode(ctx, mockReq.PurchaseUnit.Code).
// 			Return(nil, mockErr).
// 			Times(1)

// 		s.mockLogger.EXPECT().
// 			Error(mockErr).
// 			Times(1)

// 		res, err := s.underTest.Create(ctx, &mockReq)
// 		s.Require().Nil(res)
// 		s.Require().NotNil(err)
// 		s.Require().IsType(handling.ErrorResponse{}, err)

// 		resErr := err.(handling.ErrorResponse)

// 		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
// 		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
// 		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
// 	})

// 	s.Run("failed - invalid usage unit code", func() {
// 		mockReq := inventoryModule.Request{}
// 		if err := gofakeit.Struct(&mockReq); err != nil {
// 			s.T().Fatal(err)
// 		}

// 		ctx := context.Background()

// 		s.mockInventoryRepo.EXPECT().
// 			FindByName(ctx, mockReq.Name).
// 			Return(nil, define.ErrRecordNotFound).
// 			Times(1)

// 		s.mockUsageUnitRepo.EXPECT().
// 			FindByCode(ctx, mockReq.PurchaseUnit.Code).
// 			Return(nil, define.ErrRecordNotFound).
// 			Times(1)

// 		res, err := s.underTest.Create(ctx, &mockReq)
// 		s.Require().Nil(res)
// 		s.Require().NotNil(err)
// 		s.Require().IsType(handling.ErrorResponse{}, err)

// 		resErr := err.(handling.ErrorResponse)

// 		s.Require().Equal(define.CodeInvalidUsageUnit, resErr.Code)
// 		s.Require().Equal(define.MsgInvalidUsageUnit, resErr.Message)
// 		s.Require().Equal(http.StatusBadRequest, resErr.Status)
// 	})

// 	s.Run("failed - created inventory", func() {
// 		mockReq := inventoryModule.Request{}
// 		if err := gofakeit.Struct(&mockReq); err != nil {
// 			s.T().Fatal(err)
// 		}

// 		mockUsageUnitRes := usageUnitModule.Prototype{}
// 		if err := gofakeit.Struct(&mockUsageUnitRes); err != nil {
// 			s.T().Fatal(err)
// 		}

// 		mockReq.PurchaseUnit.Code = mockUsageUnitRes.Code

// 		ctx := context.Background()

// 		s.mockInventoryRepo.EXPECT().
// 			FindByName(ctx, mockReq.Name).
// 			Return(nil, define.ErrRecordNotFound).
// 			Times(1)

// 		s.mockUsageUnitRepo.EXPECT().
// 			FindByCode(ctx, mockReq.PurchaseUnit.Code).
// 			Return(&mockUsageUnitRes, nil).
// 			Times(1)

// 		mockErr := errors.New("mock err")
// 		s.mockInventoryRepo.EXPECT().
// 			Create(ctx, &mockReq).
// 			Return(mockErr).
// 			Times(1)

// 		s.mockLogger.EXPECT().
// 			Error(mockErr).
// 			Times(1)

// 		res, err := s.underTest.Create(ctx, &mockReq)
// 		s.Require().Nil(res)
// 		s.Require().NotNil(err)
// 		s.Require().IsType(handling.ErrorResponse{}, err)

// 		resErr := err.(handling.ErrorResponse)

// 		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
// 		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
// 		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
// 	})

// 	s.Run("succeed - created inventory", func() {
// 		mockReq := inventoryModule.Request{}
// 		if err := gofakeit.Struct(&mockReq); err != nil {
// 			s.T().Fatal(err)
// 		}

// 		mockUsageUnitRes := usageUnitModule.Prototype{}
// 		if err := gofakeit.Struct(&mockUsageUnitRes); err != nil {
// 			s.T().Fatal(err)
// 		}

// 		mockReq.PurchaseUnit.Code = mockUsageUnitRes.Code

// 		ctx := context.Background()

// 		s.mockInventoryRepo.EXPECT().
// 			FindByName(ctx, mockReq.Name).
// 			Return(nil, define.ErrRecordNotFound).
// 			Times(1)

// 		s.mockUsageUnitRepo.EXPECT().
// 			FindByCode(ctx, mockReq.PurchaseUnit.Code).
// 			Return(&mockUsageUnitRes, nil).
// 			Times(1)

// 		s.mockInventoryRepo.EXPECT().
// 			Create(ctx, &mockReq).
// 			Return(nil).
// 			Times(1)

// 		res, err := s.underTest.Create(ctx, &mockReq)
// 		s.Require().NotNil(res)
// 		s.Require().Nil(err)

// 		s.Require().Equal(&mockReq, res.Item)
// 	})
// }
