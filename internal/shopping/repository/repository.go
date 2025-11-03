package shoppingRepository

import (
	"context"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository.go -package=mockShoppingRepository
type IRepository interface {
	Find(ctx context.Context) ([]shoppingModule.Response, error)
	FindByID(ctx context.Context, id string) (*shoppingModule.Response, error)
	Create(ctx context.Context, req *shoppingModule.Request) error
	UpdateIsComplete(ctx context.Context, id string, req *shoppingModule.ReqUpdateIsComplete) error
	ReOrderNo(ctx context.Context, reqs []shoppingModule.ReqReOrder) error
	DeleteByID(ctx context.Context, id string) error
}

//go:generate mockgen -source=./receiptGrpc.go -destination=./mock/receiptGrpc.go -package=mockShoppingRepository
type IReceiptGrpcRepository interface {
	ReadReceipt(ctx context.Context, fileName string, file []byte) ([]shoppingModule.ResReceiptReader, error)
}
