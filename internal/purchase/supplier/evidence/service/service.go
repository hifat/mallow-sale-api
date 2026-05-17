package purchaseSupplierEvidenceService

import (
	"context"

	purchaseSupplierModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier"
	purchaseSupplierEvidenceModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence"
	purchaseSupplierOrderModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order"
	"golang.org/x/sync/errgroup"
)

type service struct {
	repo                 purchaseSupplierEvidenceModule.IRepository
	supplierRepo         purchaseSupplierModule.IRepository
	supplierEvidanceRepo purchaseSupplierEvidenceModule.IRepository
	supplierOrderRepo    purchaseSupplierOrderModule.IRepository
}

func New(repo purchaseSupplierEvidenceModule.IRepository) purchaseSupplierEvidenceModule.IService {
	return &service{repo: repo}
}

func (s *service) CreateMany(ctx context.Context, reqs []purchaseSupplierEvidenceModule.CreateEvidenceRequest, supplierID string) error {
	if len(reqs) == 0 {
		return nil
	}

	g, ctx := errgroup.WithContext(ctx)
	for _, req := range reqs {
		req := req
		g.Go(func() error {
			return s.supplierEvidanceRepo.Create(ctx, &req, supplierID)
		})
	}

	return g.Wait()
}

func (s *service) DeleteBySupplierID(ctx context.Context, supplierID string) error {
	return s.repo.DeleteBySupplierID(ctx, supplierID)
}

func (s *service) FindBySupplierID(ctx context.Context, supplierID string) ([]purchaseSupplierEvidenceModule.Response, error) {
	return s.repo.FindBySupplierID(ctx, supplierID)
}
