package settingService

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	settingModule "github.com/hifat/mallow-sale-api/internal/settings"
	mockSettingRepository "github.com/hifat/mallow-sale-api/internal/settings/repository/mock"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	mockLogger "github.com/hifat/mallow-sale-api/pkg/logger/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testInventoryServiceSuite struct {
	suite.Suite

	mockLogger      *mockLogger.MockLogger
	mockSettingRepo *mockSettingRepository.MockIRepository

	underTest IService
}

func (s *testInventoryServiceSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.mockLogger = mockLogger.NewMockLogger(ctrl)
	s.mockSettingRepo = mockSettingRepository.NewMockIRepository(ctrl)

	s.underTest = &service{
		logger:      s.mockLogger,
		settingRepo: s.mockSettingRepo,
	}
}

func TestInventoryServiceSuite(t *testing.T) {
	suite.Run(t, &testInventoryServiceSuite{})
}

func (s *testInventoryServiceSuite) TestInventoryService_Create() {
	s.T().Parallel()

	s.Run("failed - update setting other error", func() {
		ctx := context.Background()
		var costPercentage float32 = 49.99

		mockErr := errors.New("mock err")
		s.mockSettingRepo.EXPECT().
			Update(ctx, costPercentage).
			Return(mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		err := s.underTest.Update(ctx, costPercentage)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - update setting", func() {
		ctx := context.Background()
		var costPercentage float32 = 49.99

		s.mockSettingRepo.EXPECT().
			Update(ctx, costPercentage).
			Return(nil).
			Times(1)

		err := s.underTest.Update(ctx, costPercentage)
		s.Require().Nil(err)
	})
}

func (s *testInventoryServiceSuite) TestInventoryService_Find() {
	s.T().Parallel()

	s.Run("failed - update setting other error", func() {
		ctx := context.Background()

		mockErr := errors.New("mock err")
		s.mockSettingRepo.EXPECT().
			Find(ctx).
			Return(nil, mockErr).
			Times(1)

		s.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := s.underTest.Find(ctx)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("succeed - update setting", func() {
		ctx := context.Background()
		mockSetting := &settingModule.Response{}
		if err := gofakeit.Struct(&mockSetting); err != nil {
			s.T().Fatal(err)
		}

		s.mockSettingRepo.EXPECT().
			Find(ctx).
			Return(mockSetting, nil).
			Times(1)

		res, err := s.underTest.Find(ctx)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().Equal(mockSetting, res.Item)
	})
}
