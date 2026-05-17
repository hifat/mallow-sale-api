package purchaseSupplierOrderModule

import (
	"context"
)

type IRepository interface {
	Create(ctx context.Context, req *CreateOrderRequest, supplierID string) error
	DeleteBySupplierID(ctx context.Context, supplierID string) error
	FindBySupplierID(ctx context.Context, supplierID string) ([]Response, error)
}

type IService interface {
	CreateMany(ctx context.Context, reqs []CreateOrderRequest, supplierID string) error
	DeleteBySupplierID(ctx context.Context, supplierID string) error
	FindBySupplierID(ctx context.Context, supplierID string) ([]Response, error)
}
