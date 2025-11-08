package authService

import (
	"context"
	"errors"
	"log"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	authModule "github.com/hifat/mallow-sale-api/internal/auth"
	userModule "github.com/hifat/mallow-sale-api/internal/user"
	mockUserRepository "github.com/hifat/mallow-sale-api/internal/user/repository/mock"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	mockLogger "github.com/hifat/mallow-sale-api/pkg/logger/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

type mockAuthService struct {
	ctrl         *gomock.Controller
	mockLogger   *mockLogger.MockLogger
	mockUserRepo *mockUserRepository.MockIRepository
}

func NewMock(t *testing.T) mockAuthService {
	ctrl := gomock.NewController(t)

	return mockAuthService{
		ctrl: ctrl,

		mockLogger:   mockLogger.NewMockLogger(ctrl),
		mockUserRepo: mockUserRepository.NewMockIRepository(ctrl),
	}
}

func NewUnderTest(m mockAuthService) *service {
	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return &service{
		logger:   m.mockLogger,
		cfg:      cfg,
		userRepo: m.mockUserRepo,
	}
}

type testAuthServiceSuite struct {
	suite.Suite
}

func TestAuthServiceSuite(t *testing.T) {
	suite.Run(t, &testAuthServiceSuite{})
}

func (s *testAuthServiceSuite) TestAuthService_Signin() {
	s.Run("failed - find user by username other error", func() {
		m := NewMock(s.T())
		underTest := NewUnderTest(m)

		mockReq := authModule.SigninReq{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		ctx := context.Background()

		mockErr := errors.New("mock err")
		m.mockUserRepo.EXPECT().
			FindByUsername(ctx, mockReq.Username).
			Return(nil, mockErr).
			Times(1)

		m.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := underTest.Signin(ctx, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
	})

	s.Run("failed - invalid username", func() {
		m := NewMock(s.T())
		underTest := NewUnderTest(m)

		mockReq := authModule.SigninReq{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		ctx := context.Background()

		mockErr := define.ErrRecordNotFound
		m.mockUserRepo.EXPECT().
			FindByUsername(ctx, mockReq.Username).
			Return(nil, mockErr).
			Times(1)

		res, err := underTest.Signin(ctx, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInvalidCredentials, resErr.Code)
		s.Require().Equal(define.MsgInvalidCredentials, resErr.Message)
		s.Require().Equal(http.StatusUnauthorized, resErr.Status)
	})

	s.Run("failed - invalid password", func() {
		m := NewMock(s.T())
		underTest := NewUnderTest(m)

		mockReq := authModule.SigninReq{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockRes := userModule.Response{}
		if err := gofakeit.Struct(&mockRes); err != nil {
			s.T().Fatal(err)
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte("invalid_pass"), 10)
		if err != nil {
			s.T().Fatalf("Failed to hashed password: %v", err)
		}

		mockRes.Password = string(hashedPass)

		ctx := context.Background()

		m.mockUserRepo.EXPECT().
			FindByUsername(ctx, mockReq.Username).
			Return(&mockRes, nil).
			Times(1)

		res, err := underTest.Signin(ctx, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInvalidCredentials, resErr.Code)
		s.Require().Equal(define.MsgInvalidCredentials, resErr.Message)
		s.Require().Equal(http.StatusUnauthorized, resErr.Status)
	})

	s.Run("failed - invalid password", func() {
		m := NewMock(s.T())
		underTest := NewUnderTest(m)

		mockReq := authModule.SigninReq{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		mockRes := userModule.Response{}
		if err := gofakeit.Struct(&mockRes); err != nil {
			s.T().Fatal(err)
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte("invalid_pass"), 10)
		if err != nil {
			s.T().Fatalf("Failed to hashed password: %v", err)
		}

		mockRes.Password = string(hashedPass)

		ctx := context.Background()

		m.mockUserRepo.EXPECT().
			FindByUsername(ctx, mockReq.Username).
			Return(&mockRes, nil).
			Times(1)

		res, err := underTest.Signin(ctx, &mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(define.CodeInvalidCredentials, resErr.Code)
		s.Require().Equal(define.MsgInvalidCredentials, resErr.Message)
		s.Require().Equal(http.StatusUnauthorized, resErr.Status)
	})

	// s.Run("succeed - delete inventory by id", func() {
	// 	m := NewMock(s.T())
	// 	underTest := NewUnderTest(m)

	// 	mockReq := authModule.Request{}
	// 	if err := gofakeit.Struct(&mockReq); err != nil {
	// 		s.T().Fatal(err)
	// 	}

	// 	mockID := "mock-id"

	// 	ctx := context.Background()

	// 	m.mockUserRepo.EXPECT(). m
	// 		FindByID(ctx, mockID).
	// 		Return(&authModule.Response{}, nil).
	// 		Times(1)

	// 	m.mockUserRepo.EXPECT().
	// 		DeleteByID(ctx, mockID).
	// 		Return(nil).
	// 		Times(1)

	// 	err := underTest.DeleteByID(ctx, mockID)
	// 	s.Require().Nil(err)
	// })
}
