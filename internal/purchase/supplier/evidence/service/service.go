package purchaseSupplierEvidenceService

import (
	"context"

	purchaseSupplierEvidenceModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence"
	"golang.org/x/sync/errgroup"
)

type service struct {
	supplierEvidenceRepo purchaseSupplierEvidenceModule.IRepository
}

func New(supplierEvidenceRepo purchaseSupplierEvidenceModule.IRepository) purchaseSupplierEvidenceModule.IService {
	return &service{supplierEvidenceRepo: supplierEvidenceRepo}
}

func (s *service) CreateMany(ctx context.Context, reqs []purchaseSupplierEvidenceModule.CreateEvidenceRequest, supplierID string) error {
	if len(reqs) == 0 {
		return nil
	}

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(10)
	for _, req := range reqs {
		g.Go(func() error {
			return s.supplierEvidenceRepo.Create(ctx, &req, supplierID)
		})
	}

	return g.Wait()
}

func (s *service) DeleteBySupplierID(ctx context.Context, supplierID string) error {
	return s.supplierEvidenceRepo.DeleteBySupplierID(ctx, supplierID)
}

func (s *service) FindBySupplierID(ctx context.Context, supplierID string) ([]purchaseSupplierEvidenceModule.Response, error) {
	return s.supplierEvidenceRepo.FindBySupplierID(ctx, supplierID)
}
