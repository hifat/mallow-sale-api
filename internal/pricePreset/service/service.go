package pricePresetService

import (
	"context"

	inventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper"
	pricePresetModule "github.com/hifat/mallow-sale-api/internal/pricePreset"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type service struct {
	pricePresetRepository pricePresetModule.IRepository
	inventoryHelper       inventoryHelper.IHelper
	logger                logger.ILogger
}

func New(
	pricePresetRepository pricePresetModule.IRepository,
	inventoryHelper inventoryHelper.IHelper,
	logger logger.ILogger,
) pricePresetModule.IService {
	return &service{
		pricePresetRepository: pricePresetRepository,
		inventoryHelper:       inventoryHelper,
		logger:                logger,
	}
}

func (s *service) Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[pricePresetModule.Response], error) {
	count, err := s.pricePresetRepository.Count(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	presets, err := s.pricePresetRepository.Find(ctx, query)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	inventoryIDs := make([]string, 0, len(presets))
	for _, v := range presets {
		inventoryIDs = append(inventoryIDs, v.InventoryID)
	}

	getInventoryByID, err := s.inventoryHelper.FindAndGetByID(ctx, inventoryIDs)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	for i, v := range presets {
		inventory := getInventoryByID(v.InventoryID)
		if inventory != nil {
			presets[i].Inventory = &inventory.Prototype
		}
	}

	return &handling.ResponseItems[pricePresetModule.Response]{
		Items: presets,
		Meta:  handling.MetaResponse{Total: count},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*pricePresetModule.Response], error) {
	preset, err := s.pricePresetRepository.FindByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	if preset.InventoryID != "" {
		getInventoryByID, err := s.inventoryHelper.FindAndGetByID(ctx, []string{preset.InventoryID})
		if err != nil {
			s.logger.Error(err)
			return nil, handling.ThrowErr(err)
		}
		inventory := getInventoryByID(preset.InventoryID)
		if inventory != nil {
			preset.Inventory = &inventory.Prototype
		}
	}

	return &handling.ResponseItem[*pricePresetModule.Response]{Item: preset}, nil
}
