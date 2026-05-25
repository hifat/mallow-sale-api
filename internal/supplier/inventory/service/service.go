package inventoryService

import (
	"context"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	inventorySupplierModule "github.com/hifat/mallow-sale-api/internal/supplier/inventory"
	supplierInventoryModule "github.com/hifat/mallow-sale-api/internal/supplier/inventory"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type service struct {
	logger              logger.ILogger
	supplierRepository  supplierModule.IRepository
	inventoryRepository inventoryModule.IRepository
}

func New(
	logger logger.ILogger,
	supplierRepository supplierModule.IRepository,
	inventoryRepository inventoryModule.IRepository,
) inventorySupplierModule.IService {
	return &service{
		logger:              logger,
		supplierRepository:  supplierRepository,
		inventoryRepository: inventoryRepository,
	}
}

func (s *service) FindGroupBySupplier(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[supplierInventoryModule.GroupBySupplierResponse], error) {
	errgroup, ctx := errgroup.WithContext(ctx)

	var suppliers []supplierModule.Response
	errgroup.Go(func() error {
		var err error
		suppliers, err = s.supplierRepository.Find(ctx, &utilsModule.QueryReq{})
		if err != nil {
			s.logger.Error(err)
			return err
		}

		return nil
	})

	var inventories []inventoryModule.Response
	errgroup.Go(func() error {
		var err error
		inventories, err = s.inventoryRepository.Find(ctx, &utilsModule.QueryReq{})
		if err != nil {
			s.logger.Error(err)
			return err
		}

		return nil
	})

	if err := errgroup.Wait(); err != nil {
		return nil, handling.ThrowErr(err)
	}

	supplierGroups := make(map[string][]inventoryModule.Response)
	for _, inventory := range inventories {
		supplierID := inventory.SupplierID
		supplierGroups[inventory.SupplierID] = append(supplierGroups[supplierID], inventory)
	}

	var res []supplierInventoryModule.GroupBySupplierResponse
	for _, supplier := range suppliers {
		inventories, ok := supplierGroups[supplier.ID]
		if !ok {
			inventories = []inventoryModule.Response{}
		}

		res = append(res, supplierInventoryModule.GroupBySupplierResponse{
			Prototype: supplierModule.Prototype{
				ID:     supplier.ID,
				Name:   supplier.Name,
				ImgUrl: supplier.ImgUrl,
			},
			Inventories: inventories,
		})
	}

	return &handling.ResponseItems[supplierInventoryModule.GroupBySupplierResponse]{
		Items: res,
		Meta:  handling.MetaResponse{Total: int64(len(res))},
	}, nil
}
