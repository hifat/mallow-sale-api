package purchaseService

import (
	"context"

	purchasePresetModule "github.com/hifat/mallow-sale-api/internal/purchase/preset"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type service struct {
	log  logger.ILogger
	repo purchasePresetModule.IRepository
}

func New(log logger.ILogger, repo purchasePresetModule.IRepository) purchasePresetModule.IService {
	return &service{log: log, repo: repo}
}

func (s *service) Create(ctx context.Context, req *purchasePresetModule.Request) (*handling.ResponseItem[*purchasePresetModule.Response], error) {
	id, err := s.repo.Create(ctx, req)
	if err != nil {
		s.log.Error("Error creating purchase preset", "error", err)
		return nil, handling.ThrowErr(err)
	}

	res := &purchasePresetModule.Response{
		ID: id,
	}

	return &handling.ResponseItem[*purchasePresetModule.Response]{
		Item: res,
	}, nil
}

func (s *service) Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[*purchasePresetModule.Response], error) {
	g, gctx := errgroup.WithContext(ctx)

	var total int64
	g.Go(func() error {
		count, err := s.repo.Count(gctx, query)
		if err != nil {
			return err
		}

		total = count
		return nil
	})

	var items []*purchasePresetModule.Response
	g.Go(func() error {
		res, err := s.repo.Find(gctx, query)
		if err != nil {
			return err
		}

		items := make([]*purchasePresetModule.Response, len(res))
		for i := range res {
			items[i] = &res[i]
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		s.log.Error("Error finding purchase presets", "error", err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[*purchasePresetModule.Response]{
		Items: items,
		Meta: handling.MetaResponse{
			Total: total,
		},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*purchasePresetModule.Response], error) {
	res, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if err != define.ErrRecordNotFound {
			s.log.Error("Error finding purchase preset", "error", err)
		}

		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*purchasePresetModule.Response]{
		Item: res,
	}, nil
}

func (s *service) UpdateByID(ctx context.Context, id string, req *purchasePresetModule.Request) (*handling.ResponseItem[*purchasePresetModule.Response], error) {
	if err := s.repo.UpdateByID(ctx, id, req); err != nil {
		return nil, handling.ThrowErr(err)
	}

	res, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if err != define.ErrRecordNotFound {
			s.log.Error("Error finding purchase preset", "error", err)
		}

		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*purchasePresetModule.Response]{
		Item: res,
	}, nil
}

func (s *service) DeleteByID(ctx context.Context, id string) error {
	if err := s.repo.DeleteByID(ctx, id); err != nil {
		s.log.Error("Error deleting purchase preset", "error", err)
		return handling.ThrowErr(err)
	}

	return nil
}
