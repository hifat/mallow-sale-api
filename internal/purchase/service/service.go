package purchaseService

import (
	"context"

	purchaseModule "github.com/hifat/mallow-sale-api/internal/purchase"
	purchaseSupplierModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier"
	purchaseSupplierEvidenceModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence"
	purchaseSupplierOrderModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"golang.org/x/sync/errgroup"
)

type service struct {
	repo                 purchaseModule.IRepository
	supplierRepo         purchaseSupplierModule.IRepository
	supplierOrderRepo    purchaseSupplierOrderModule.IRepository
	supplierEvidenceRepo purchaseSupplierEvidenceModule.IRepository
	utilsRepo            utilsModule.IRepository
}

func New(
	repo purchaseModule.IRepository,
	supplierRepo purchaseSupplierModule.IRepository,
	supplierOrderRepo purchaseSupplierOrderModule.IRepository,
	supplierEvidenceRepo purchaseSupplierEvidenceModule.IRepository,
	utilsRepo utilsModule.IRepository,
) purchaseModule.IService {
	return &service{
		repo:                 repo,
		supplierRepo:         supplierRepo,
		supplierOrderRepo:    supplierOrderRepo,
		supplierEvidenceRepo: supplierEvidenceRepo,
		utilsRepo:            utilsRepo,
	}
}

func (s *service) Create(ctx context.Context, req *purchaseModule.CreatePurchaseRequest) error {
	req.ID = s.utilsRepo.NewID()
	purchaseID, err := s.repo.Create(ctx, req)
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)
	for _, supplierReq := range req.Suppliers {
		g.Go(func() error {
			supplierReq.PurchaseSupplierID = s.utilsRepo.NewID()
			supplierID, err := s.supplierRepo.Create(ctx, &supplierReq, purchaseID)
			if err != nil {
				return err
			}

			sg, sctx := errgroup.WithContext(ctx)
			sg.SetLimit(10)

			for _, orderReq := range supplierReq.Orders {
				sg.Go(func() error {
					orderReq.PurchaseSupplierID = supplierReq.PurchaseSupplierID
					return s.supplierOrderRepo.Create(sctx, &orderReq, supplierID)
				})
			}

			for _, evidenceReq := range supplierReq.Evidences {
				sg.Go(func() error {
					evidenceReq.PurchaseSupplierID = supplierReq.PurchaseSupplierID
					return s.supplierEvidenceRepo.Create(sctx, &evidenceReq, supplierID)
				})
			}

			return sg.Wait()
		})
	}

	return g.Wait()
}

func (s *service) FindByID(ctx context.Context, id string) (*purchaseModule.Response, error) {
	purchase, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	suppliers, err := s.supplierRepo.FindByPurchaseID(ctx, purchase.ID)
	if err != nil {
		return nil, err
	}

	g, ctx := errgroup.WithContext(ctx)
	for i := range suppliers {
		i := i
		g.Go(func() error {
			orders, err := s.supplierOrderRepo.FindBySupplierID(ctx, suppliers[i].ID)
			if err != nil {
				return err
			}
			suppliers[i].Orders = orders

			evidences, err := s.supplierEvidenceRepo.FindBySupplierID(ctx, suppliers[i].ID)
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

	purchase.Suppliers = suppliers
	return purchase, nil
}

func (s *service) DeleteByID(ctx context.Context, id string) error {
	suppliers, err := s.supplierRepo.FindByPurchaseID(ctx, id)
	if err == nil {
		g, ctxDel := errgroup.WithContext(ctx)
		for _, supplier := range suppliers {
			supplierID := supplier.ID
			g.Go(func() error {
				if err := s.supplierOrderRepo.DeleteBySupplierID(ctxDel, supplierID); err != nil {
					return err
				}
				return s.supplierEvidenceRepo.DeleteBySupplierID(ctxDel, supplierID)
			})
		}
		if err := g.Wait(); err == nil {
			_ = s.supplierRepo.DeleteByPurchaseID(ctx, id)
		}
	}

	return s.repo.DeleteByID(ctx, id)
}

func (s *service) UpdateByID(ctx context.Context, id string, req *purchaseModule.CreatePurchaseRequest) error {
	if err := s.repo.UpdateByID(ctx, id, req); err != nil {
		return err
	}

	suppliers, err := s.supplierRepo.FindByPurchaseID(ctx, id)
	if err == nil {
		gDel, ctxDel := errgroup.WithContext(ctx)
		for _, supplier := range suppliers {
			supplierID := supplier.ID
			gDel.Go(func() error {
				if err := s.supplierOrderRepo.DeleteBySupplierID(ctxDel, supplierID); err != nil {
					return err
				}
				return s.supplierEvidenceRepo.DeleteBySupplierID(ctxDel, supplierID)
			})
		}
		if err := gDel.Wait(); err == nil {
			_ = s.supplierRepo.DeleteByPurchaseID(ctx, id)
		}
	}

	g, ctx := errgroup.WithContext(ctx)
	for _, supplierReq := range req.Suppliers {
		supplierReq := supplierReq
		g.Go(func() error {
			supplierID, err := s.supplierRepo.Create(ctx, &supplierReq, id)
			if err != nil {
				return err
			}

			sg, sctx := errgroup.WithContext(ctx)

			for _, orderReq := range supplierReq.Orders {
				orderReq := orderReq
				sg.Go(func() error {
					return s.supplierOrderRepo.Create(sctx, &orderReq, supplierID)
				})
			}

			for _, evidenceReq := range supplierReq.Evidences {
				evidenceReq := evidenceReq
				sg.Go(func() error {
					return s.supplierEvidenceRepo.Create(sctx, &evidenceReq, supplierID)
				})
			}

			return sg.Wait()
		})
	}
	return g.Wait()
}
