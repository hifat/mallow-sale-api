package purchaseModule

import (
	"context"
)

type IRepository interface {
	Create(ctx context.Context, req *CreatePurchaseRequest) (string, error)
	FindByID(ctx context.Context, id string) (*Response, error)
	DeleteByID(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, id string, req *CreatePurchaseRequest) error
}

type IService interface {
	Create(ctx context.Context, req *CreatePurchaseRequest) error
	FindByID(ctx context.Context, id string) (*Response, error)
	DeleteByID(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, id string, req *CreatePurchaseRequest) error
}
