package supplierHelper

import (
	"context"

	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
)

type Helper interface {
	FindAndGetByID(ctx context.Context, ids []string) (func(id string) *supplierModule.Response, error)
}

type helper struct {
	supplierRepository supplierRepository.IRepository
}

func New(supplierRepository supplierRepository.IRepository) Helper {
	return &helper{
		supplierRepository: supplierRepository,
	}
}

func (h *helper) FindAndGetByID(ctx context.Context, ids []string) (func(id string) *supplierModule.Response, error) {
	suppliers, err := h.supplierRepository.FindInIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return func(id string) *supplierModule.Response {
		for _, supplier := range suppliers {
			if supplier.ID == id {
				return &supplier
			}
		}

		return nil
	}, nil
}
