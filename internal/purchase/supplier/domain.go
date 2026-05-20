package purchaseSupplierModule

import (
	"context"
)

type IRepository interface {
	Create(ctx context.Context, req *CreateSupplierRequest, purchaseID string) (string, error)
	DeleteByPurchaseID(ctx context.Context, purchaseID string) error
	FindByPurchaseID(ctx context.Context, purchaseID string) ([]Response, error)
}
