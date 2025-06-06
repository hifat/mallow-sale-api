package inventoryService

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/hifat/mallow-sale-api/internal/entity"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/inventoryRepository/mock"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository/mock"
	mockCore "github.com/hifat/mallow-sale-api/pkg/utils/mock/core"
	mockRules "github.com/hifat/mallow-sale-api/pkg/utils/mock/rules"
	mockValidator "github.com/hifat/mallow-sale-api/pkg/utils/mock/rules"
	"github.com/hifat/mallow-sale-api/pkg/utils/response"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testInventoryServiceSuite struct {
	suite.Suite

	mockHelper            *mockCore.Mockhelper
	mockLogger            *mockCore.Mocklogger
	mockValidator         *mockRules.Mockvalidator
	mockInventoryRepo     *mockInventoryRepository.MockIInventoryRepository
	mockUsageUnitGRPCRepo *mockUsageUnitRepository.MockIUsageUnitGRPCRepository

	underTest IInventoryService
}

func (s *testInventoryServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.mockHelper = mockCore.NewMockhelper(ctrl)
	s.mockLogger = mockCore.NewMocklogger(ctrl)
	s.mockValidator = mockValidator.NewMockvalidator(ctrl)
	s.mockInventoryRepo = mockInventoryRepository.NewMockIInventoryRepository(ctrl)
	s.mockUsageUnitGRPCRepo = mockUsageUnitRepository.NewMockIUsageUnitGRPCRepository(ctrl)

	s.underTest = &inventoryService{
		helper:            s.mockHelper,
		logger:            s.mockLogger,
		validator:         s.mockValidator,
		inventoryRepo:     s.mockInventoryRepo,
		usageUnitGRPCRepo: s.mockUsageUnitGRPCRepo,
	}
}

func TestInventoryServiceSuite(t *testing.T) {
	suite.Run(t, &testInventoryServiceSuite{})
}

func (s *testInventoryServiceSuite) TestInventoryService_Create() {
	s.T().Parallel()

	s.Run("fail - validate", func() {
		req := inventory.InventoryReq{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		errValidate := errors.New("validate error")
		s.mockValidator.EXPECT().
			Validate(req).
			Return(errValidate)

		err := s.underTest.Create(context.Background(), req)
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
	})

	s.Run("fail - mapUsageUnit_FindIn", func() {
		req := inventory.InventoryReq{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		s.mockValidator.EXPECT().
			Validate(req).
			Return(nil)

		errMapUsageUnit := errors.New("mapUsageUnit_FindIn error")
		s.mockUsageUnitGRPCRepo.EXPECT().
			FindIn(context.Background(), gomock.Any()).
			Return(nil, errMapUsageUnit)

		s.mockLogger.EXPECT().
			Error(errMapUsageUnit)

		err := s.underTest.Create(context.Background(), req)
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
	})

	s.Run("fail - validateFiled", func() {
		req := inventory.InventoryReq{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		s.mockValidator.EXPECT().
			Validate(req).
			Return(nil)

		s.mockUsageUnitGRPCRepo.EXPECT().
			FindIn(context.Background(), usageUnit.FilterReq{
				Codes: []string{
					req.PurchaseUnitCode,
				},
			}).
			Return([]usageUnit.UsageUnit{
				{
					Base: entity.Base{
						ID:        "mock",
						CreatedAt: &time.Time{},
						UpdatedAt: &time.Time{},
					},
					Code: "mock-code",
					Name: "mock-name",
				},
			}, nil)

		err := s.underTest.Create(context.Background(), req)
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
	})

	s.Run("fail - create", func() {
		req := inventory.InventoryReq{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		s.mockValidator.EXPECT().
			Validate(req).
			Return(nil)

		s.mockUsageUnitGRPCRepo.EXPECT().
			FindIn(context.Background(), usageUnit.FilterReq{
				Codes: []string{
					req.PurchaseUnitCode,
				},
			}).
			DoAndReturn(func(context.Context, usageUnit.FilterReq) ([]usageUnit.UsageUnit, error) {
				return []usageUnit.UsageUnit{
					{
						Base: entity.Base{
							ID:        "mock",
							CreatedAt: &time.Time{},
							UpdatedAt: &time.Time{},
						},
						Code: req.PurchaseUnitCode,
						Name: "mock-name",
					},
				}, nil
			})

		newReq := req
		newReq.PurchaseUnit.SetAttr(req.PurchaseUnitCode, "mock-name")

		errCreate := errors.New("create error")
		s.mockInventoryRepo.EXPECT().
			Create(context.Background(), newReq).
			Return("", errCreate)

		s.mockLogger.EXPECT().Error(errCreate)

		err := s.underTest.Create(context.Background(), req)

		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
	})

	s.Run("success - create", func() {
		req := inventory.InventoryReq{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		s.mockValidator.EXPECT().
			Validate(req).
			Return(nil)

		s.mockUsageUnitGRPCRepo.EXPECT().
			FindIn(context.Background(), usageUnit.FilterReq{
				Codes: []string{
					req.PurchaseUnitCode,
				},
			}).
			DoAndReturn(func(context.Context, usageUnit.FilterReq) ([]usageUnit.UsageUnit, error) {
				return []usageUnit.UsageUnit{
					{
						Base: entity.Base{
							ID:        "mock",
							CreatedAt: &time.Time{},
							UpdatedAt: &time.Time{},
						},
						Code: req.PurchaseUnitCode,
						Name: "mock-name",
					},
				}, nil
			})

		newReq := req
		newReq.PurchaseUnit.SetAttr(req.PurchaseUnitCode, "mock-name")

		s.mockInventoryRepo.EXPECT().
			Create(context.Background(), newReq).
			Return(gomock.Any().String(), nil)

		err := s.underTest.Create(context.Background(), req)

		s.Require().Nil(err)
	})
}

func (s *testInventoryServiceSuite) TestInventoryService_Find() {
	s.T().Parallel()

	s.Run("fail - find", func() {
		errFind := errors.New("mock-error")
		s.mockInventoryRepo.EXPECT().
			Find(context.Background()).
			Return(nil, errFind)

		s.mockLogger.EXPECT().
			Error(errFind)

		_, err := s.underTest.Find(context.Background())
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
	})

	s.Run("fail - copy", func() {
		s.mockInventoryRepo.EXPECT().
			Find(context.Background()).
			Return([]inventory.Inventory{}, nil)

		errCopy := errors.New("mock-error")
		s.mockHelper.EXPECT().
			Copy(gomock.Any(), gomock.Any()).
			Return(errCopy)

		s.mockLogger.EXPECT().
			Error(errCopy)

		_, err := s.underTest.Find(context.Background())
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
	})

	s.Run("success - find", func() {
		mockInventories := make([]inventory.Inventory, 2)
		for i := range mockInventories {
			if err := gofakeit.Struct(&mockInventories[i]); err != nil {
				s.T().Fatal(err)
			}
		}

		s.mockInventoryRepo.EXPECT().
			Find(context.Background()).
			Return(mockInventories, nil)

		inventories := []inventory.InventoryRes{}
		s.mockHelper.EXPECT().
			Copy(&inventories, mockInventories).
			Return(nil)

		res, err := s.underTest.Find(context.Background())
		s.Require().Nil(err)
		s.Require().Equal(res, inventories)
	})
}
