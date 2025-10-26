package supplierService

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	mockSupplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository/mock"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	mockLogger "github.com/hifat/mallow-sale-api/pkg/logger/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testSupplierServiceSuite struct {
	suite.Suite

	// mockLogger       *mockLogger.MockLogger
	// mockSupplierRepo *mockSupplierRepository.MockIRepository

	// underTest IService
}

type mockSupplierService struct {
	mockLogger       *mockLogger.MockLogger
	mockSupplierRepo *mockSupplierRepository.MockIRepository
}

func NewMock(t *testing.T) (mockSupplierService, func()) {
	ctrl := gomock.NewController(t)

	return mockSupplierService{
			mockLogger:       mockLogger.NewMockLogger(ctrl),
			mockSupplierRepo: mockSupplierRepository.NewMockIRepository(ctrl),
		}, func() {
			ctrl.Finish()
		}
}

func NewUnderTest(m *mockSupplierService) *service {
	return &service{
		logger:             m.mockLogger,
		supplierRepository: m.mockSupplierRepo,
	}
}

func (s *testSupplierServiceSuite) SetupTest() {
	// ctrl := gomock.NewController((s.T()))
	// s.T().Cleanup(func() {
	// 	ctrl.Finish()
	// })

	// m.mockLogger = mockLogger.NewMockLogger(ctrl)
	// m.mockSupplierRepo = mockSupplierRepository.NewMockIRepository(ctrl)

	// underTest = &service{
	// 	logger:             m.mockLogger,
	// 	supplierRepository: m.mockSupplierRepo,
	// }
}

func TestSupplierServiceSuite(t *testing.T) {
	suite.Run(t, &testSupplierServiceSuite{})
}

func (s *testSupplierServiceSuite) TestSupplierService_Create() {
	s.T().Parallel()

	s.Run("failed - create supplier error", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()
		mockReq := &supplierModule.Request{}

		mockErr := errors.New("mock-err")
		m.mockSupplierRepo.EXPECT().
			Create(ctx, mockReq).
			Return(mockErr).
			Times(1)

		m.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := underTest.Create(ctx, mockReq)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
	})

	s.Run("succeed - create supplier", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()
		mockReq := &supplierModule.Request{}
		if err := gofakeit.Struct(mockReq); err != nil {
			s.T().Fatal(err)
		}

		m.mockSupplierRepo.EXPECT().
			Create(ctx, mockReq).
			Return(nil).
			Times(1)

		res, err := underTest.Create(ctx, mockReq)
		s.Require().Nil(err)
		s.Require().NotNil(res)

		s.Require().Equal(mockReq, res.Item)
	})
}

func (s *testSupplierServiceSuite) TestSupplierService_Find() {
	s.T().Parallel()

	s.Run("failed - count supplier error", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		var total int64 = 3
		mockErr := errors.New("mock-err")
		m.mockSupplierRepo.EXPECT().
			Count(ctx).
			Return(total, mockErr).
			Times(1)

		m.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := underTest.Find(ctx, &utilsModule.QueryReq{})
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
	})

	s.Run("failed - find supplier error", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		var total int64 = 3
		m.mockSupplierRepo.EXPECT().
			Count(ctx).
			Return(total, nil).
			Times(1)

		mockErr := errors.New("mock-err")
		m.mockSupplierRepo.EXPECT().
			Find(ctx, &utilsModule.QueryReq{}).
			Return(nil, mockErr).
			Times(1)

		m.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := underTest.Find(ctx, &utilsModule.QueryReq{})
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
	})

	s.Run("succeed - create supplier", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		var total int64 = 3
		m.mockSupplierRepo.EXPECT().
			Count(ctx).
			Return(total, nil).
			Times(1)

		mockSuppliers := make([]supplierModule.Response, total)
		m.mockSupplierRepo.EXPECT().
			Find(ctx, &utilsModule.QueryReq{}).
			Return(mockSuppliers, nil).
			Times(1)

		res, err := underTest.Find(ctx, &utilsModule.QueryReq{})
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().IsType([]supplierModule.Response{}, res.Items)
		s.Require().Equal(total, int64(len(res.Items)))
		s.Require().NotNil(res.Meta)
		s.Require().Equal(total, res.Meta.Total)
	})
}

func (s *testSupplierServiceSuite) TestSupplierService_FindByID() {
	s.T().Parallel()

	s.Run("failed - find supplier by id other error", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		mockErr := errors.New("mock-err")
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, mockErr).
			Times(1)

		m.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := underTest.FindByID(ctx, mockID)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
	})

	s.Run("failed - find supplier by id not found error", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		res, err := underTest.FindByID(ctx, mockID)
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusNotFound, resErr.Status)
		s.Require().Equal(define.CodeRecordNotFound, resErr.Code)
		s.Require().Equal(define.MsgRecordNotFound, resErr.Message)
	})

	s.Run("succeed - find supplier by id", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		mockSupplier := &supplierModule.Response{}
		if err := gofakeit.Struct(mockSupplier); err != nil {
			s.T().Fatal(err)
		}

		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(mockSupplier, nil).
			Times(1)

		res, err := underTest.FindByID(ctx, mockID)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().NotNil(res.Item)
		s.Require().Equal(mockSupplier, res.Item)
	})
}

func (s *testSupplierServiceSuite) TestSupplierService_UpdateByID() {
	s.T().Parallel()

	s.Run("failed - find supplier by id other error", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		mockErr := errors.New("mock-err")
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, mockErr).
			Times(1)

		m.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := underTest.UpdateByID(ctx, mockID, &supplierModule.Request{})
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
	})

	s.Run("failed - find supplier by id not found error", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		res, err := underTest.UpdateByID(ctx, mockID, &supplierModule.Request{})
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusNotFound, resErr.Status)
		s.Require().Equal(define.CodeRecordNotFound, resErr.Code)
		s.Require().Equal(define.MsgRecordNotFound, resErr.Message)
	})

	s.Run("failed - update supplier by id", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&supplierModule.Response{}, nil).
			Times(1)

		mockErr := errors.New("mock-err")
		m.mockSupplierRepo.EXPECT().
			UpdateByID(ctx, mockID, &supplierModule.Request{}).
			Return(mockErr).
			Times(1)

		m.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		res, err := underTest.UpdateByID(ctx, mockID, &supplierModule.Request{})
		s.Require().Nil(res)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
	})

	s.Run("succeed - update supplier by id", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&supplierModule.Response{}, nil).
			Times(1)

		mockReq := &supplierModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		m.mockSupplierRepo.EXPECT().
			UpdateByID(ctx, mockID, mockReq).
			Return(nil).
			Times(1)

		res, err := underTest.UpdateByID(ctx, mockID, mockReq)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().Equal(mockReq, res.Item)
	})
}

func (s *testSupplierServiceSuite) TestSupplierService_DeleteByID() {
	s.T().Parallel()

	s.Run("failed - find supplier by id other error", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		mockErr := errors.New("mock-err")
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, mockErr).
			Times(1)

		m.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		err := underTest.DeleteByID(ctx, mockID)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
	})

	s.Run("failed - find supplier by id not found error", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(nil, define.ErrRecordNotFound).
			Times(1)

		err := underTest.DeleteByID(ctx, mockID)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusNotFound, resErr.Status)
		s.Require().Equal(define.CodeRecordNotFound, resErr.Code)
		s.Require().Equal(define.MsgRecordNotFound, resErr.Message)
	})

	s.Run("failed - delete supplier by id", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&supplierModule.Response{}, nil).
			Times(1)

		mockErr := errors.New("mock-err")
		m.mockSupplierRepo.EXPECT().
			DeleteByID(ctx, mockID).
			Return(mockErr).
			Times(1)

		m.mockLogger.EXPECT().
			Error(mockErr).
			Times(1)

		err := underTest.DeleteByID(ctx, mockID)
		s.Require().NotNil(err)
		s.Require().IsType(handling.ErrorResponse{}, err)

		resErr := err.(handling.ErrorResponse)

		s.Require().Equal(http.StatusInternalServerError, resErr.Status)
		s.Require().Equal(define.CodeInternalServerError, resErr.Code)
		s.Require().Equal(define.MsgInternalServerError, resErr.Message)
	})

	s.Run("succeed - delete supplier by id", func() {
		m, cleanup := NewMock(s.T())
		defer cleanup()

		underTest := NewUnderTest(&m)

		ctx := context.Background()

		mockID := "mock-id"
		m.mockSupplierRepo.EXPECT().
			FindByID(ctx, mockID).
			Return(&supplierModule.Response{}, nil).
			Times(1)

		mockReq := &supplierModule.Request{}
		if err := gofakeit.Struct(&mockReq); err != nil {
			s.T().Fatal(err)
		}

		m.mockSupplierRepo.EXPECT().
			DeleteByID(ctx, mockID).
			Return(nil).
			Times(1)

		err := underTest.DeleteByID(ctx, mockID)
		s.Require().Nil(err)
	})
}
