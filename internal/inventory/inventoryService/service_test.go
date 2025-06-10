package inventoryService

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/hifat/goroger-core/rules"
	"github.com/hifat/mallow-sale-api/internal/entity"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/inventoryRepository/mock"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	mockUsageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository/mock"
	"github.com/hifat/mallow-sale-api/pkg/throw"
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

func (s *testInventoryServiceSuite) testMapUsageUnit() {
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
}

func (s *testInventoryServiceSuite) testValidateField() {
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
}

func (s *testInventoryServiceSuite) testValidate(req inventory.InventoryReq) {
	errValidate := rules.ValidateErrs{
		"name": "the name is required",
	}
	s.mockValidator.EXPECT().
		Validate(req).
		Return(errValidate)

	err := s.underTest.Update(context.Background(), "mock-id", req)
	s.Require().NotNil(err)
	s.Require().IsType(response.ResponseErr{}, err)

	resErr := err.(response.ResponseErr)

	s.Require().Equal(resErr.Status, http.StatusBadRequest)
	s.Require().Equal(resErr.Code, throw.CodeInvalidForm)
	s.Require().Equal(resErr.Message, throw.MsgInvalidForm)
	s.Require().NotEmpty(resErr.Attribute)
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
			Return([]usageUnit.UsageUnit{
				{
					Base: entity.Base{
						ID:        "mock",
						CreatedAt: &time.Time{},
						UpdatedAt: &time.Time{},
					},
					Code: req.PurchaseUnitCode,
					Name: "mock-name",
				},
			}, nil)

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
		gofakeit.Slice(&mockInventories)

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

func (s *testInventoryServiceSuite) TestInventoryService_FindByID() {
	s.T().Parallel()

	mockID := "mock-id"

	s.Run("fail - find by id", func() {
		s.mockInventoryRepo.EXPECT().
			FindByID(context.Background(), mockID).
			Return(nil, throw.ErrRecordNotFound)

		s.mockLogger.EXPECT().
			Error(throw.ErrRecordNotFound)

		res, err := s.underTest.FindByID(context.Background(), mockID)
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
		s.Require().Nil(res)

		errRes := err.(response.ResponseErr)

		s.Require().Equal(http.StatusNotFound, errRes.Status)
		s.Require().Equal(throw.CodeRecordNotFound, errRes.Code)
		s.Require().Equal(throw.ErrRecordNotFound.Error(), errRes.Message)
	})

	s.Run("fail - copy", func() {
		_inventory := &inventory.Inventory{}
		s.mockInventoryRepo.EXPECT().
			FindByID(context.Background(), mockID).
			Return(_inventory, nil)

		mockRes := inventory.InventoryRes{}
		errCopy := errors.New("error-copy")
		s.mockHelper.EXPECT().
			Copy(&mockRes, _inventory).
			Return(errCopy)

		s.mockLogger.EXPECT().
			Error(errCopy)

		res, err := s.underTest.FindByID(context.Background(), mockID)
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)
		s.Require().Nil(res)

		errRes := err.(response.ResponseErr)

		s.Require().Equal(http.StatusInternalServerError, errRes.Status)
		s.Require().Equal(throw.CodeInternalServer, errRes.Code)
		s.Require().Equal(throw.ErrInternalServer.Error(), errRes.Message)
	})

	s.Run("success - find by id", func() {
		_inventory := &inventory.Inventory{}
		if err := gofakeit.Struct(_inventory); err != nil {
			s.T().Fatal(err)
		}

		s.mockInventoryRepo.EXPECT().
			FindByID(context.Background(), mockID).
			Return(_inventory, nil)

		mockRes := inventory.InventoryRes{}
		s.mockHelper.EXPECT().
			Copy(&mockRes, _inventory).
			Return(nil)

		res, err := s.underTest.FindByID(context.Background(), mockID)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().Equal(mockRes, *res)
	})
}

func (s *testInventoryServiceSuite) TestInventoryService_FindIn() {
	s.T().Parallel()

	filter := inventory.FilterReq{
		Codes: []string{"mock-code"},
	}

	s.Run("fail - find in", func() {
		errFindIn := errors.New("mock error")
		s.mockInventoryRepo.EXPECT().
			FindIn(context.Background(), filter).
			Return(nil, errFindIn)

		res, err := s.underTest.FindIn(context.Background(), filter)
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)

		errRes := err.(response.ResponseErr)

		s.Require().Equal(http.StatusInternalServerError, errRes.Status)
		s.Require().Equal(throw.CodeInternalServer, errRes.Code)
		s.Require().Equal(throw.ErrInternalServer.Error(), errRes.Message)

		s.Require().NotNil(res)
		s.Require().Equal([]inventory.InventoryRes{}, res)
	})

	s.Run("success - find in return empty slice", func() {
		s.mockInventoryRepo.EXPECT().
			FindIn(context.Background(), filter).
			Return(nil, nil)

		res, err := s.underTest.FindIn(context.Background(), filter)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().Equal([]inventory.InventoryRes{}, res)
	})

	s.Run("success - find in", func() {
		inventories := make([]inventory.Inventory, 2)
		gofakeit.Slice(&inventories)

		s.mockInventoryRepo.EXPECT().
			FindIn(context.Background(), filter).
			Return(inventories, nil)

		res, err := s.underTest.FindIn(context.Background(), filter)
		s.Require().Nil(err)
		s.Require().NotNil(res)
		s.Require().IsType([]inventory.InventoryRes{}, res)
		s.Require().Equal(len(inventories), len(res))
	})
}

func (s *testInventoryServiceSuite) TestInventoryService_Update() {
	s.T().Parallel()

	s.Run("fail - validate", func() {
		req := inventory.InventoryReq{}
		if err := gofakeit.Struct(&req); err != nil {
			s.T().Fatal(err)
		}

		//TODO: Fix helper test to log this line(Parent func) on failed
		s.testValidate(req)
	})

	s.Run("fail - mapUsageUnit", func() {
		s.testMapUsageUnit()
	})

	s.Run("fail - validateField", func() {
		s.testValidateField()
	})

	s.Run("fail - update", func() {
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
					Code: req.PurchaseUnitCode,
					Name: "mock-name",
				},
			}, nil)

		newReq := req
		newReq.PurchaseUnit.SetAttr(req.PurchaseUnitCode, "mock-name")

		errUpdate := errors.New("error-update")
		s.mockInventoryRepo.EXPECT().
			Update(context.Background(), "mock-id", newReq).
			Return(errUpdate)

		s.mockLogger.EXPECT().
			Error(errUpdate)

		err := s.underTest.Update(context.Background(), "mock-id", req)
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)

		errRes := err.(response.ResponseErr)

		s.Require().Equal(http.StatusInternalServerError, errRes.Status)
		s.Require().Equal(throw.CodeInternalServer, errRes.Code)
		s.Require().Equal(throw.ErrInternalServer.Error(), errRes.Message)
	})

	s.Run("success - update", func() {
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
					Code: req.PurchaseUnitCode,
					Name: "mock-name",
				},
			}, nil)

		newReq := req
		newReq.PurchaseUnit.SetAttr(req.PurchaseUnitCode, "mock-name")

		s.mockInventoryRepo.EXPECT().
			Update(context.Background(), "mock-id", newReq).
			Return(nil)

		err := s.underTest.Update(context.Background(), "mock-id", req)
		s.Require().Nil(err)
	})
}

func (s *testInventoryServiceSuite) TestInventoryService_Delete() {
	s.T().Parallel()

	s.Run("fail - delete", func() {
		errDelete := errors.New("error-delete")
		s.mockInventoryRepo.EXPECT().
			Delete(context.Background(), "mock-id").
			Return(errDelete)

		s.mockLogger.EXPECT().
			Error(errDelete)

		err := s.underTest.Delete(context.Background(), "mock-id")
		s.Require().NotNil(err)
		s.Require().IsType(response.ResponseErr{}, err)

		errRes := err.(response.ResponseErr)

		s.Require().Equal(http.StatusInternalServerError, errRes.Status)
		s.Require().Equal(throw.CodeInternalServer, errRes.Code)
		s.Require().Equal(throw.ErrInternalServer.Error(), errRes.Message)
	})

	s.Run("success - delete", func() {
		s.mockInventoryRepo.EXPECT().
			Delete(context.Background(), "mock-id").
			Return(nil)

		err := s.underTest.Delete(context.Background(), "mock-id")
		s.Require().Nil(err)
	})
}
