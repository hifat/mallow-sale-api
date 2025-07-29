package service

import (
	"context"

	promotionModule "github.com/hifat/mallow-sale-api/internal/promotion"
	promotionRepository "github.com/hifat/mallow-sale-api/internal/promotion/repository"
	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	recipeHelper "github.com/hifat/mallow-sale-api/internal/recipe/helper"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type Service interface {
	Create(ctx context.Context, req *promotionModule.Request) (*handling.ResponseItem[*promotionModule.Request], error)
	Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[promotionModule.Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*promotionModule.Response], error)
	UpdateByID(ctx context.Context, id string, req *promotionModule.Request) (*handling.ResponseItem[*promotionModule.Request], error)
	DeleteByID(ctx context.Context, id string) error
}

type service struct {
	logger              logger.Logger
	promotionRepository promotionRepository.Repository
	recipeHelper        recipeHelper.Helper
}

func New(
	logger logger.Logger,
	promotionRepository promotionRepository.Repository,
	recipeHelper recipeHelper.Helper,
) Service {
	return &service{
		logger:              logger,
		promotionRepository: promotionRepository,
		recipeHelper:        recipeHelper,
	}
}

func (s *service) Create(ctx context.Context, req *promotionModule.Request) (*handling.ResponseItem[*promotionModule.Request], error) {
	// Validate the request based on promotion type
	if err := req.Validate(); err != nil {
		return nil, handling.ThrowErr(err)
	}

	if err := s.promotionRepository.Create(ctx, req); err != nil {
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*promotionModule.Request]{
		Item: req,
	}, nil
}

func (s *service) Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[promotionModule.Response], error) {
	res, err := s.promotionRepository.Find(ctx, query)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	count, err := s.promotionRepository.Count(ctx)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[promotionModule.Response]{
		Items: res,
		Meta: handling.MetaResponse{
			Total: count,
		},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*promotionModule.Response], error) {
	res, err := s.promotionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	res.Products = make([]recipeModule.Response, 0, len(res.Products))
	getRecipeByID, err := s.recipeHelper.FindAndGetByID(ctx, res.GetProductIDs())
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	for _, v := range res.Products {
		recipe := getRecipeByID(v.ID)
		if recipe == nil {
			s.logger.Warn("recipe not found: ", v.ID)
			continue
		}

		res.Products = append(res.Products, *recipe)
	}

	return &handling.ResponseItem[*promotionModule.Response]{
		Item: res,
	}, nil
}

func (s *service) UpdateByID(ctx context.Context, id string, req *promotionModule.Request) (*handling.ResponseItem[*promotionModule.Request], error) {
	if err := req.Validate(); err != nil {
		return nil, handling.ThrowErr(err)
	}

	if err := s.promotionRepository.UpdateByID(ctx, id, req); err != nil {
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*promotionModule.Request]{
		Item: req,
	}, nil
}

func (s *service) DeleteByID(ctx context.Context, id string) error {
	_, err := s.promotionRepository.FindByID(ctx, id)
	if err != nil {
		return handling.ThrowErr(err)
	}

	if err := s.promotionRepository.DeleteByID(ctx, id); err != nil {
		return handling.ThrowErr(err)
	}

	return nil
}
