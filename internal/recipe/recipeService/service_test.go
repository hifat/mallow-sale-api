package recipeService

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/inventoryRepository/mock"
	"github.com/hifat/mallow-sale-api/internal/recipe"
	mockRecipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/recipeRepository/mock"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository/mock"
	"github.com/hifat/mallow-sale-api/pkg/throw"
	mockCore "github.com/hifat/mallow-sale-api/pkg/utils/mock/core"
	mockRules "github.com/hifat/mallow-sale-api/pkg/utils/mock/rules"
	"github.com/hifat/mallow-sale-api/pkg/utils/response"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testRecipeServiceSuite struct {
	suite.Suite

	mockLogger            *mockCore.Mocklogger
	mockValidator         *mockRules.Mockvalidator
	mockHelper            *mockCore.Mockhelper
	mockRecipeRepo        *mockRecipeRepository.MockIRecipeRepository
	mockUsageUnitGrpcRepo *mockUsageUnitRepository.MockIUsageUnitGRPCRepository
	mockInventoryRepo     *mockInventoryRepository.MockIInventoryRepository
	mockInventoryGrpcRepo *mockInventoryRepository.MockIInventoryGRPCRepository

	underTest IRecipeService
}

func (s *testRecipeServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockCore.NewMocklogger(ctrl)
	s.mockValidator = mockRules.NewMockvalidator(ctrl)
	s.mockHelper = mockCore.NewMockhelper(ctrl)
	s.mockRecipeRepo = mockRecipeRepository.NewMockIRecipeRepository(ctrl)
	s.mockInventoryGrpcRepo = mockInventoryRepository.NewMockIInventoryGRPCRepository(ctrl)

	s.underTest = &recipeService{
		logger:            s.mockLogger,
		validator:         s.mockValidator,
		helper:            s.mockHelper,
		recipeRepo:        s.mockRecipeRepo,
		usageUnitGRPCRepo: s.mockUsageUnitGrpcRepo,
		inventoryRepo:     s.mockInventoryRepo,
		inventoryGRPCRepo: s.mockInventoryGrpcRepo,
	}
}

func TestRecipeServiceSuite(t *testing.T) {
	suite.Run(t, &testRecipeServiceSuite{})
}

func (s *testRecipeServiceSuite) TestRecipeService_Find() {
	s.T().Parallel()

	s.Run("fail - find", func() {
		errFind := errors.New("mock-error")
		s.mockRecipeRepo.EXPECT().
			Find(context.Background()).
			Return(nil, errFind)

		s.mockLogger.EXPECT().
			Error(errFind)

		res, err := s.underTest.Find(context.Background())
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
		s.Require().NotNil(res)

		errRes := err.(response.ResponseErr)

		s.Require().Equal(throw.CodeInternalServer, errRes.Code)
		s.Require().Equal(http.StatusInternalServerError, errRes.Status)
		s.Require().Equal(throw.ErrInternalServer.Error(), errRes.Message)
	})

	s.Run("success - find", func() {
		recipes := make([]recipe.RecipeRes, 2)
		gofakeit.Slice(&recipes)

		s.mockRecipeRepo.EXPECT().
			Find(context.Background()).
			Return(recipes, nil)

		res, err := s.underTest.Find(context.Background())
		s.Require().Nil(err)
		s.Require().NotEmpty(res)
	})
}

func (s *testRecipeServiceSuite) TestRecipeService_FindByID() {
	s.T().Parallel()

	s.Run("fail - find by id record not found", func() {
		errNotFound := throw.ErrRecordNotFound
		s.mockRecipeRepo.EXPECT().
			FindByID(context.Background(), "mock-id").
			Return(nil, errNotFound)

		s.mockLogger.EXPECT().
			Error(errNotFound)

		res, err := s.underTest.FindByID(context.Background(), "mock-id")
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
		s.Require().Nil(res)

		errRes := err.(response.ResponseErr)

		s.Require().Equal(throw.CodeRecordNotFound, errRes.Code)
		s.Require().Equal(http.StatusNotFound, errRes.Status)
		s.Require().Equal(throw.ErrRecordNotFound.Error(), errRes.Message)
	})

	s.Run("fail - find by id", func() {
		errFindByID := errors.New("mock-err")
		s.mockRecipeRepo.EXPECT().
			FindByID(context.Background(), "mock-id").
			Return(nil, errFindByID)

		s.mockLogger.EXPECT().
			Error(errFindByID)

		res, err := s.underTest.FindByID(context.Background(), "mock-id")
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
		s.Require().Nil(res)

		errRes := err.(response.ResponseErr)

		s.Require().Equal(throw.CodeInternalServer, errRes.Code)
		s.Require().Equal(http.StatusInternalServerError, errRes.Status)
		s.Require().Equal(throw.ErrInternalServer.Error(), errRes.Message)
	})

	s.Run("fail - find in", func() {
		_recipe := &recipe.RecipeRes{}
		_recipe.Ingredients = make([]recipe.RecipeInventoryRes, 2)
		if err := gofakeit.Struct(&_recipe); err != nil {
			s.T().Fatal(err)
		}

		s.mockRecipeRepo.EXPECT().
			FindByID(context.Background(), "mock-id").
			Return(_recipe, nil)

		errFindIn := errors.New("mock-error")
		s.mockInventoryGrpcRepo.EXPECT().
			FindIn(context.Background(), inventory.FilterReq{
				IDs: _recipe.GetInventoryIDs(),
			}).
			Return(nil, errFindIn)

		s.mockLogger.EXPECT().
			Error(errFindIn)

		res, err := s.underTest.FindByID(context.Background(), "mock-id")
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
		s.Require().Nil(res)

		errRes := err.(response.ResponseErr)

		s.Require().Equal(throw.CodeInternalServer, errRes.Code)
		s.Require().Equal(http.StatusInternalServerError, errRes.Status)
		s.Require().Equal(throw.ErrInternalServer.Error(), errRes.Message)
	})

	s.Run("warn - some inventory id not round", func() {
		_recipe := &recipe.RecipeRes{}
		_recipe.Ingredients = make([]recipe.RecipeInventoryRes, 1)
		if err := gofakeit.Struct(&_recipe); err != nil {
			s.T().Fatal(err)
		}

		s.mockRecipeRepo.EXPECT().
			FindByID(context.Background(), "mock-id").
			Return(_recipe, nil)

		inventories := make([]inventory.Inventory, 1)
		s.mockInventoryGrpcRepo.EXPECT().
			FindIn(context.Background(), inventory.FilterReq{
				IDs: _recipe.GetInventoryIDs(),
			}).
			Return(inventories, nil)

		for _, v := range _recipe.Ingredients {
			s.mockLogger.EXPECT().
				Warn(fmt.Sprintf("%s: %s", MsgNotFoundInventoryID, v.InventoryID))
		}

		res, err := s.underTest.FindByID(context.Background(), "mock-id")
		s.Require().Nil(err)
		s.Require().NotNil(res)
	})

	s.Run("fail - **", func() {
		// TODO: Check Warn when inventory id not found
		// TODO: Check when copy error
	})

	// TODO: Success case
}
