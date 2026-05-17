package purchaseSupplierEvidenceModule

import (
	"context"
)

type IRepository interface {
	Create(ctx context.Context, req *CreateEvidenceRequest, supplierID string) error
	DeleteBySupplierID(ctx context.Context, supplierID string) error
	FindBySupplierID(ctx context.Context, supplierID string) ([]Response, error)
}

type IService interface {
	CreateMany(ctx context.Context, reqs []CreateEvidenceRequest, supplierID string) error
	DeleteBySupplierID(ctx context.Context, supplierID string) error
	FindBySupplierID(ctx context.Context, supplierID string) ([]Response, error)
}
