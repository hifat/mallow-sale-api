package purchaseService

import (
	"context"

	purchaseModule "github.com/hifat/mallow-sale-api/internal/purchase"
	purchaseSupplierModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier"
	purchaseSupplierEvidenceModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence"
	purchaseSupplierOrderModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type service struct {
	repo                 purchaseModule.IRepository
	supplierRepo         purchaseSupplierModule.IRepository
	supplierOrderRepo    purchaseSupplierOrderModule.IRepository
	supplierEvidenceRepo purchaseSupplierEvidenceModule.IRepository
	utilsRepo            utilsModule.IRepository
	logger               logger.ILogger
}

func New(
	repo purchaseModule.IRepository,
	supplierRepo purchaseSupplierModule.IRepository,
	supplierOrderRepo purchaseSupplierOrderModule.IRepository,
	supplierEvidenceRepo purchaseSupplierEvidenceModule.IRepository,
	utilsRepo utilsModule.IRepository,
	logger logger.ILogger,
) purchaseModule.IService {
	return &service{
		repo:                 repo,
		supplierRepo:         supplierRepo,
		supplierOrderRepo:    supplierOrderRepo,
		supplierEvidenceRepo: supplierEvidenceRepo,
		utilsRepo:            utilsRepo,
		logger:               logger,
	}
}

func (s *service) Create(ctx context.Context, req *purchaseModule.CreatePurchaseRequest) (*handling.ResponseItem[*purchaseModule.CreatePurchaseRequest], error) {
	req.ID = s.utilsRepo.NewID()

	g, ctxGroup := errgroup.WithContext(ctx)

	g.Go(func() error {
		_, err := s.repo.Create(ctxGroup, req)
		return err
	})

	sg, sctx := errgroup.WithContext(ctx)
	sg.SetLimit(10)

	for _, supplierReq := range req.Suppliers {
		sg.Go(func() error {
			supplierReq.PurchaseSupplierID = s.utilsRepo.NewID()
			supplierID, err := s.supplierRepo.Create(sctx, &supplierReq, req.ID)
			if err != nil {
				return err
			}

			for _, orderReq := range supplierReq.Orders {
				sg.Go(func() error {
					orderReq.PurchaseSupplierID = supplierReq.PurchaseSupplierID
					return s.supplierOrderRepo.Create(sctx, &orderReq, supplierID)
				})
			}

			return nil
		})
	}

	if err := sg.Wait(); err != nil {
		s.DeleteByID(ctx, req.ID)
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	if err := g.Wait(); err != nil {
		s.DeleteByID(ctx, req.ID)
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*purchaseModule.CreatePurchaseRequest]{Item: req}, nil
}

func (s *service) Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[purchaseModule.Response], error) {
	count, err := s.repo.Count(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	purchases, err := s.repo.Find(ctx, query)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[purchaseModule.Response]{
		Items: purchases,
		Meta:  handling.MetaResponse{Total: count},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*purchaseModule.Response], error) {
	g, gctx := errgroup.WithContext(ctx)

	var purchase *purchaseModule.Response
	g.Go(func() error {
		var err error
		purchase, err = s.repo.FindByID(gctx, id)
		return err
	})

	var suppliers []purchaseSupplierModule.Response
	g.Go(func() error {
		var err error
		suppliers, err = s.supplierRepo.FindByPurchaseID(gctx, id)
		return err
	})

	if err := g.Wait(); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	sg, sctx := errgroup.WithContext(ctx)
	sg.SetLimit(10)

	for i := range suppliers {
		sg.Go(func() error {
			ig, ictx := errgroup.WithContext(sctx)
			ig.SetLimit(10)

			var orders []purchaseSupplierOrderModule.Response
			ig.Go(func() error {
				var err error
				orders, err = s.supplierOrderRepo.FindBySupplierID(ictx, suppliers[i].ID)
				return err
			})

			var evidences []purchaseSupplierEvidenceModule.Response
			ig.Go(func() error {
				var err error
				evidences, err = s.supplierEvidenceRepo.FindBySupplierID(ictx, suppliers[i].ID)
				return err
			})

			if err := ig.Wait(); err != nil {
				return err
			}

			suppliers[i].Orders = orders
			suppliers[i].Evidences = evidences

			return nil
		})
	}

	if err := sg.Wait(); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	purchase.Suppliers = suppliers
	return &handling.ResponseItem[*purchaseModule.Response]{Item: purchase}, nil
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

	err = s.repo.DeleteByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}
	return nil
}

func (s *service) UpdateByID(ctx context.Context, id string, req *purchaseModule.CreatePurchaseRequest) (*handling.ResponseItem[*purchaseModule.CreatePurchaseRequest], error) {
	if err := s.repo.UpdateByID(ctx, id, req); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	suppliers, err := s.supplierRepo.FindByPurchaseID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	gDel, ctxDel := errgroup.WithContext(ctx)
	gDel.SetLimit(10)

	for _, supplier := range suppliers {
		gDel.Go(func() error {
			if err := s.supplierOrderRepo.DeleteBySupplierID(ctxDel, supplier.ID); err != nil {
				return err
			}

			return s.supplierEvidenceRepo.DeleteBySupplierID(ctxDel, supplier.ID)
		})
	}

	if err := gDel.Wait(); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	if err := s.supplierRepo.DeleteByPurchaseID(ctx, id); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	g, ctxGrp := errgroup.WithContext(ctx)
	g.SetLimit(10)

	for _, supplierReq := range req.Suppliers {
		g.Go(func() error {
			supplierID, err := s.supplierRepo.Create(ctxGrp, &supplierReq, id)
			if err != nil {
				return err
			}

			sg, sctx := errgroup.WithContext(ctxGrp)

			for _, orderReq := range supplierReq.Orders {
				orderReq := orderReq
				sg.Go(func() error {
					return s.supplierOrderRepo.Create(sctx, &orderReq, supplierID)
				})
			}

			return sg.Wait()
		})
	}

	if err := g.Wait(); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*purchaseModule.CreatePurchaseRequest]{Item: req}, nil
}
