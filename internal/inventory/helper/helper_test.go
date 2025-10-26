package helper

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository/mock"
	"github.com/hifat/mallow-sale-api/pkg/utils"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testInventoryHelperSuite struct {
	suite.Suite

	mockInventoryRepo *mockInventoryRepository.MockIRepository

	underTest helper
}

func (s *testInventoryHelperSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.mockInventoryRepo = mockInventoryRepository.NewMockIRepository(ctrl)

	s.underTest = helper{
		inventoryRepo: s.mockInventoryRepo,
	}
}

func TestInventoryHelperSuite(t *testing.T) {
	suite.Run(t, &testInventoryHelperSuite{})
}

func (s *testInventoryHelperSuite) TestInventoryHelper_FindAndGetByID() {
	s.T().Parallel()

	s.Run("failed - find inventory in ids other error", func() {
		mockIDs := []string{"mock-id"}
		ctx := context.Background()

		mockErr := errors.New("mock-err")
		s.mockInventoryRepo.EXPECT().
			FindInIDs(ctx, mockIDs).
			Return(nil, mockErr).
			Times(1)

		_, err := s.underTest.FindAndGetByID(ctx, mockIDs)
		s.Require().NotNil(err)
		s.Require().Equal(mockErr.Error(), err.Error())
	})

	s.Run("failed - get inventory by id return nil", func() {
		maxRecords := 3
		mockInventories := make([]inventoryModule.Response, maxRecords)
		gofakeit.Slice(&mockInventories)

		mockIDs := []string{}
		for i, v := range mockInventories {
			if i == maxRecords-1 {
				break
			}

			mockIDs = append(mockIDs, v.ID)
		}

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindInIDs(ctx, mockIDs).
			Return(mockInventories[:maxRecords-1], nil).
			Times(1)

		getInventoryByID, err := s.underTest.FindAndGetByID(ctx, mockIDs)
		s.Require().Nil(err)

		inventory := getInventoryByID(mockInventories[2].ID)
		s.Require().Nil(inventory)
	})

	s.Run("succeed - get inventory by id", func() {
		mockInventories := make([]inventoryModule.Response, 3)
		gofakeit.Slice(&mockInventories)

		mockIDs := []string{}
		for _, v := range mockInventories {
			mockIDs = append(mockIDs, v.ID)
		}

		ctx := context.Background()

		s.mockInventoryRepo.EXPECT().
			FindInIDs(ctx, mockIDs).
			Return(mockInventories, nil).
			Times(1)

		getInventoryByID, err := s.underTest.FindAndGetByID(ctx, mockIDs)
		s.Require().Nil(err)

		inventory := getInventoryByID(mockInventories[1].ID)

		s.Require().Equal(&mockInventories[1], inventory)
	})
}

func (s *testInventoryHelperSuite) TestInventoryHelper_calPurchasePrice() {
	s.T().Parallel()

	s.Run("should increase purchase quantity is 0 should return request purchase quantity", func() {
		reqPurchasePrice := 5.5546
		inventory := inventoryModule.Response{
			Prototype: inventoryModule.Prototype{
				PurchaseQuantity: 0,
			},
		}

		currentPrice := s.underTest.calPurchasePrice(inventory, reqPurchasePrice, true)
		s.Require().Equal(utils.RoundToDecimals(reqPurchasePrice, 3), currentPrice)
	})

	s.Run("should increase inventory purchase price", func() {
		reqPurchasePrice := 5.5546
		inventory := inventoryModule.Response{
			Prototype: inventoryModule.Prototype{
				PurchasePrice:    10,
				PurchaseQuantity: 10,
			},
		}

		currentPrice := s.underTest.calPurchasePrice(inventory, reqPurchasePrice, true)
		s.Require().Equal(utils.RoundToDecimals(inventory.PurchasePrice+reqPurchasePrice, 3), currentPrice)
	})

	s.Run("should decrease purchase quantity is 0 should return 0", func() {
		reqPurchasePrice := 5.5546
		inventory := inventoryModule.Response{
			Prototype: inventoryModule.Prototype{
				PurchaseQuantity: 0,
			},
		}

		currentPrice := s.underTest.calPurchasePrice(inventory, reqPurchasePrice, false)
		s.Require().Equal(utils.RoundToDecimals(0, 3), currentPrice)
	})

	s.Run("should decrease inventory purchase price", func() {
		reqPurchasePrice := 5.5546
		inventory := inventoryModule.Response{
			Prototype: inventoryModule.Prototype{
				PurchasePrice:    10,
				PurchaseQuantity: 10,
			},
		}

		currentPrice := s.underTest.calPurchasePrice(inventory, reqPurchasePrice, false)
		s.Require().Equal(utils.RoundToDecimals(inventory.PurchasePrice-reqPurchasePrice, 3), currentPrice)
	})
}

func (s *testInventoryHelperSuite) TestInventoryHelper_IncreaseStock() {
	s.T().Parallel()

	reqPurchaseQuantity := float64(44)
	reqPurchasePrice := float64(55)

	s.Run("failed - find inventory by id error", func() {
		ctx := context.Background()
		inventoryID := "mock-id"

		mockErr := errors.New("mock-err")
		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, inventoryID).
			Return(nil, mockErr).
			Times(1)

		err := s.underTest.IncreaseStock(ctx, inventoryID, reqPurchaseQuantity, reqPurchasePrice)
		s.Require().NotNil(err)
		s.Require().Equal(mockErr.Error(), err.Error())
	})

	s.Run("failed - update inventory stock error", func() {
		ctx := context.Background()
		inventoryID := "mock-id"

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, inventoryID).
			Return(&inventoryModule.Response{}, nil).
			Times(1)

		mockErr := errors.New("mock-err")
		s.mockInventoryRepo.EXPECT().
			UpdateStock(ctx, inventoryID, reqPurchaseQuantity, reqPurchasePrice).
			Return(mockErr).
			Times(1)

		err := s.underTest.IncreaseStock(ctx, inventoryID, reqPurchaseQuantity, reqPurchasePrice)
		s.Require().NotNil(err)
		s.Require().Equal(mockErr.Error(), err.Error())
	})

	s.Run("succeed - update inventory stock", func() {
		ctx := context.Background()
		inventoryID := "mock-id"

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, inventoryID).
			Return(&inventoryModule.Response{}, nil).
			Times(1)

		s.mockInventoryRepo.EXPECT().
			UpdateStock(ctx, inventoryID, reqPurchaseQuantity, reqPurchasePrice).
			Return(nil).
			Times(1)

		err := s.underTest.IncreaseStock(ctx, inventoryID, reqPurchaseQuantity, reqPurchasePrice)
		s.Require().Nil(err)
	})
}

func (s *testInventoryHelperSuite) TestInventoryHelper_DecreaseStock() {
	s.T().Parallel()

	reqPurchasePrice := float64(40)
	reqPurchaseQuantity := float64(50)

	s.Run("failed - find inventory by id error", func() {
		ctx := context.Background()
		inventoryID := "mock-id"

		mockErr := errors.New("mock-err")
		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, inventoryID).
			Return(nil, mockErr).
			Times(1)

		err := s.underTest.DecreaseStock(ctx, inventoryID, reqPurchaseQuantity, reqPurchasePrice)
		s.Require().NotNil(err)
		s.Require().Equal(mockErr.Error(), err.Error())
	})

	s.Run("failed - update inventory stock error", func() {
		ctx := context.Background()
		inventoryID := "mock-id"

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, inventoryID).
			Return(&inventoryModule.Response{
				Prototype: inventoryModule.Prototype{
					PurchasePrice:    50,
					PurchaseQuantity: 60,
				},
			}, nil).
			Times(1)

		mockErr := errors.New("mock-err")
		s.mockInventoryRepo.EXPECT().
			UpdateStock(ctx, inventoryID, float64(10), float64(10)).
			Return(mockErr).
			Times(1)

		err := s.underTest.DecreaseStock(ctx, inventoryID, reqPurchaseQuantity, reqPurchasePrice)
		s.Require().NotNil(err)
		s.Require().Equal(mockErr.Error(), err.Error())
	})

	s.Run("succeed - update inventory stock", func() {
		ctx := context.Background()
		inventoryID := "mock-id"

		s.mockInventoryRepo.EXPECT().
			FindByID(ctx, inventoryID).
			Return(&inventoryModule.Response{
				Prototype: inventoryModule.Prototype{
					PurchasePrice:    50,
					PurchaseQuantity: 60,
				},
			}, nil).
			Times(1)

		s.mockInventoryRepo.EXPECT().
			UpdateStock(ctx, inventoryID, float64(10), float64(10)).
			Return(nil).
			Times(1)

		err := s.underTest.DecreaseStock(ctx, inventoryID, reqPurchaseQuantity, reqPurchasePrice)
		s.Require().Nil(err)
	})
}
