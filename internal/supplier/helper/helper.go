package supplierHelper

import (
	"context"

	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
)

type IHelper interface {
	FindAndGetByID(ctx context.Context, ids []string) (func(id string) *supplierModule.Response, error)
}

type helper struct {
	supplierRepository supplierModule.IRepository
}

func New(supplierRepository supplierModule.IRepository) IHelper {
	return &helper{
		supplierRepository: supplierRepository,
	}
}

func (h *helper) FindAndGetByID(ctx context.Context, ids []string) (func(id string) *supplierModule.Response, error) {
	if len(ids) == 0 {
		return nil, nil
	}

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
