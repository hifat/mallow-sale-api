package purchaseSupplierService

import (
	"context"

	purchaseSupplierModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier"
	purchaseSupplierEvidenceModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence"
	purchaseSupplierOrderModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order"
	"golang.org/x/sync/errgroup"
)

type service struct {
	repo            purchaseSupplierModule.IRepository
	orderService    purchaseSupplierOrderModule.IService
	evidenceService purchaseSupplierEvidenceModule.IService
}

func New(
	repo purchaseSupplierModule.IRepository,
	orderService purchaseSupplierOrderModule.IService,
	evidenceService purchaseSupplierEvidenceModule.IService,
) purchaseSupplierModule.IService {
	return &service{
		repo:            repo,
		orderService:    orderService,
		evidenceService: evidenceService,
	}
}

func (s *service) Create(ctx context.Context, req *purchaseSupplierModule.CreateSupplierRequest, purchaseID string) error {
	supplierID, err := s.repo.Create(ctx, req, purchaseID)
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return s.orderService.CreateMany(ctx, req.Orders, supplierID)
	})

	g.Go(func() error {
		return s.evidenceService.CreateMany(ctx, req.Evidences, supplierID)
	})

	return g.Wait()
}

func (s *service) DeleteByPurchaseID(ctx context.Context, purchaseID string) error {
	suppliers, err := s.repo.FindByPurchaseID(ctx, purchaseID)
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)
	for _, supplier := range suppliers {
		supplierID := supplier.ID
		g.Go(func() error {
			if err := s.orderService.DeleteBySupplierID(ctx, supplierID); err != nil {
				return err
			}
			return s.evidenceService.DeleteBySupplierID(ctx, supplierID)
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	return s.repo.DeleteByPurchaseID(ctx, purchaseID)
}

func (s *service) FindByPurchaseID(ctx context.Context, purchaseID string) ([]purchaseSupplierModule.Response, error) {
	suppliers, err := s.repo.FindByPurchaseID(ctx, purchaseID)
	if err != nil {
		return nil, err
	}

	g, ctx := errgroup.WithContext(ctx)
	for i := range suppliers {
		i := i
		g.Go(func() error {
			orders, err := s.orderService.FindBySupplierID(ctx, suppliers[i].ID)
			if err != nil {
				return err
			}
			suppliers[i].Orders = orders

			evidences, err := s.evidenceService.FindBySupplierID(ctx, suppliers[i].ID)
			if err != nil {
				return err
			}
			suppliers[i].Evidences = evidences

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return suppliers, nil
}
