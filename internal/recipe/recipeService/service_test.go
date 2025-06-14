package recipeService

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/hifat/mallow-sale-api/internal/recipe"
	mockRecipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/recipeRepository/mock"
	"github.com/hifat/mallow-sale-api/pkg/throw"
	mockCore "github.com/hifat/mallow-sale-api/pkg/utils/mock/core"
	"github.com/hifat/mallow-sale-api/pkg/utils/response"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testRecipeServiceSuite struct {
	suite.Suite

	mockLogger     *mockCore.Mocklogger
	mockRecipeRepo *mockRecipeRepository.MockIRecipeRepository

	underTest IRecipeService
}

func (s *testRecipeServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockCore.NewMocklogger(ctrl)
	s.mockRecipeRepo = mockRecipeRepository.NewMockIRecipeRepository(ctrl)

	s.underTest = &recipeService{
		logger:     s.mockLogger,
		recipeRepo: s.mockRecipeRepo,
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
