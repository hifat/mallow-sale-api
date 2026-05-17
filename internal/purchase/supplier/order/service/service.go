package purchaseSupplierOrderService

import (
	"context"

	purchaseSupplierOrderModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order"
	"golang.org/x/sync/errgroup"
)

type service struct {
	repo purchaseSupplierOrderModule.IRepository
}

func New(repo purchaseSupplierOrderModule.IRepository) purchaseSupplierOrderModule.IService {
	return &service{repo: repo}
}

func (s *service) CreateMany(ctx context.Context, reqs []purchaseSupplierOrderModule.CreateOrderRequest, supplierID string) error {
	if len(reqs) == 0 {
		return nil
	}

	g, ctx := errgroup.WithContext(ctx)
	for _, req := range reqs {
		req := req
		g.Go(func() error {
			return s.repo.Create(ctx, &req, supplierID)
		})
	}
	return g.Wait()
}

func (s *service) DeleteBySupplierID(ctx context.Context, supplierID string) error {
	return s.repo.DeleteBySupplierID(ctx, supplierID)
}

func (s *service) FindBySupplierID(ctx context.Context, supplierID string) ([]purchaseSupplierOrderModule.Response, error) {
	return s.repo.FindBySupplierID(ctx, supplierID)
}
