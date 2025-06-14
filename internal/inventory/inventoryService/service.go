package inventoryService

import (
	"context"

	core "github.com/hifat/goroger-core"
	"github.com/hifat/goroger-core/rules"
	"github.com/hifat/mallow-sale-api/constant"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryRepository"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository"
	"github.com/hifat/mallow-sale-api/pkg/throw"
)

type IInventoryService interface {
	Create(ctx context.Context, req inventory.InventoryReq) error
	Find(ctx context.Context) ([]inventory.InventoryRes, error)
	FindByID(ctx context.Context, id string) (*inventory.InventoryRes, error)
	FindIn(ctx context.Context, filter inventory.FilterReq) ([]inventory.InventoryRes, error)
	Update(ctx context.Context, id string, req inventory.InventoryReq) error
	Delete(ctx context.Context, id string) error
}

type inventoryService struct {
	helper            core.Helper
	logger            core.Logger
	validator         rules.Validator
	inventoryRepo     inventoryRepository.IInventoryRepository
	usageUnitGRPCRepo usageUnitRepository.IUsageUnitGRPCRepository
}

func New(helper core.Helper, logger core.Logger, validator rules.Validator, inventoryRepo inventoryRepository.IInventoryRepository, usageUnitGRPCRepo usageUnitRepository.IUsageUnitGRPCRepository) IInventoryService {
	return &inventoryService{
		helper,
		logger,
		validator,
		inventoryRepo,
		usageUnitGRPCRepo,
	}
}

func (s *inventoryService) mapUsageUnit(ctx context.Context, codes []string) (map[string]string, error) {
	_usageUnits, err := s.usageUnitGRPCRepo.FindIn(ctx, usageUnit.FilterReq{
		Codes: codes,
	})
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	unitCodeMap := make(map[string]string)
	for _, usageUnit := range _usageUnits {
		unitCodeMap[usageUnit.Code] = usageUnit.Name
	}

	return unitCodeMap, nil
}

func (s *inventoryService) validateField(ctx context.Context, req inventory.InventoryReq, unitCodeMap map[string]string) error {
	if _, ok := unitCodeMap[req.PurchaseUnitCode]; !ok {
		return constant.ErrPurchaseUnitCode
	}

	return nil
}

func (s *inventoryService) Create(ctx context.Context, req inventory.InventoryReq) error {
	if err := s.validator.Validate(req); err != nil {
		return throw.ValidateErr(err)
	}

	reqUnitCodes := []string{
		req.PurchaseUnitCode,
	}

	unitCodeMap, err := s.mapUsageUnit(ctx, reqUnitCodes)
	if err != nil {
		return throw.InternalServerErr(err)
	}

	if err := s.validateField(ctx, req, unitCodeMap); err != nil {
		return throw.BadRequestErr(err)
	}

	req.PurchaseUnit.SetAttr(req.PurchaseUnitCode, unitCodeMap[req.PurchaseUnitCode])

	if _, err := s.inventoryRepo.Create(ctx, req); err != nil {
		s.logger.Error(err)
		return throw.InternalServerErr(err)
	}

	return nil
}

func (s *inventoryService) Find(ctx context.Context) ([]inventory.InventoryRes, error) {
	inventories, err := s.inventoryRepo.Find(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, throw.InternalServerErr(err)
	}

	res := []inventory.InventoryRes{}
	if err := s.helper.Copy(&res, inventories); err != nil {
		s.logger.Error(err)
		return nil, throw.InternalServerErr(err)
	}

	return res, nil
}

func (s *inventoryService) FindByID(ctx context.Context, id string) (*inventory.InventoryRes, error) {
	_inventory, err := s.inventoryRepo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, throw.WhenRecordNotFoundErr(err)
	}

	res := inventory.InventoryRes{}
	if err := s.helper.Copy(&res, _inventory); err != nil {
		s.logger.Error(err)
		return nil, throw.InternalServerErr(err)
	}

	return &res, nil
}

func (s *inventoryService) FindIn(ctx context.Context, filter inventory.FilterReq) ([]inventory.InventoryRes, error) {
	inventories, err := s.inventoryRepo.FindIn(ctx, filter)
	if err != nil {
		return []inventory.InventoryRes{}, throw.InternalServerErr(err)
	}

	res := make([]inventory.InventoryRes, 0, len(inventories))
	for _, v := range inventories {
		item := inventory.InventoryRes{
			InventoryPrototype: inventory.InventoryPrototype{
				ID:               v.ID,
				Name:             v.Name,
				PurchasePrice:    v.PurchasePrice,
				PurchaseQuantity: v.PurchaseQuantity,
				YieldPercentage:  v.YieldPercentage,
				Remark:           v.Remark,
				CreatedAt:        v.CreatedAt,
				UpdatedAt:        v.UpdatedAt,
			},
		}

		item.PurchaseUnit.SetAttr(v.PurchaseUnit.Code, v.PurchaseUnit.Name)

		res = append(res, item)
	}

	return res, nil
}

func (s *inventoryService) Update(ctx context.Context, id string, req inventory.InventoryReq) error {
	if err := s.validator.Validate(req); err != nil {
		return throw.ValidateErr(err)
	}

	reqUnitCodes := []string{
		req.PurchaseUnitCode,
	}

	unitCodeMap, err := s.mapUsageUnit(ctx, reqUnitCodes)
	if err != nil {
		s.logger.Error(err)
		return throw.InternalServerErr(err)
	}

	if err := s.validateField(ctx, req, unitCodeMap); err != nil {
		return throw.BadRequestErr(err)
	}

	req.PurchaseUnit.SetAttr(req.PurchaseUnitCode, unitCodeMap[req.PurchaseUnitCode])

	if err := s.inventoryRepo.Update(ctx, id, req); err != nil {
		s.logger.Error(err)
		return throw.InternalServerErr(err)
	}

	return nil
}

func (s *inventoryService) Delete(ctx context.Context, id string) error {
	if err := s.inventoryRepo.Delete(ctx, id); err != nil {
		s.logger.Error(err)
		return throw.InternalServerErr(err)
	}

	return nil
}
