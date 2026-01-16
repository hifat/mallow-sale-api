package shoppingModule

import (
	"context"

	"github.com/hifat/mallow-sale-api/pkg/handling"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository.go -package=mockShoppingRepository
type IRepository interface {
	Find(ctx context.Context) ([]Response, error)
	FindByID(ctx context.Context, id string) (*Response, error)
	Create(ctx context.Context, req *Request) error
	UpdateIsComplete(ctx context.Context, id string, req *ReqUpdateIsComplete) error
	ReOrderNo(ctx context.Context, reqs []ReqReOrder) error
	DeleteByID(ctx context.Context, id string) error
}

type IService interface {
	Find(ctx context.Context) (*handling.ResponseItems[Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*Response], error)
	Create(ctx context.Context, req *Request) (*handling.ResponseItem[*Request], error)
	UpdateIsComplete(ctx context.Context, id string, req *ReqUpdateIsComplete) (*handling.Response, error)
	DeleteByID(ctx context.Context, id string) (*handling.Response, error)
}

/* -------------------------------------------------------------------------- */
/*                             Shopping Inventory                             */
/* -------------------------------------------------------------------------- */

type IInventoryRepository interface {
	Create(ctx context.Context, req *RequestShoppingInventory) error
	Find(ctx context.Context) ([]InventoryResponse, error)
	DeleteByID(ctx context.Context, id string) error
}

type IInventoryService interface {
	Create(ctx context.Context, req *RequestShoppingInventory) (*handling.ResponseItem[*RequestShoppingInventory], error)
	Find(ctx context.Context) (*handling.ResponseItems[InventoryResponse], error)
	DeleteByID(ctx context.Context, id string) error
}

/* -------------------------------------------------------------------------- */
/*                                   Receipt                                  */
/* -------------------------------------------------------------------------- */

//go:generate mockgen -source=./receiptGrpc.go -destination=./mock/receiptGrpc.go -package=mockShoppingRepository
type IReceiptGrpcRepository interface {
	ReadReceipt(ctx context.Context, fileName string, file []byte) ([]ResReceiptReader, error)
}

type IReceiptService interface {
	Reader(ctx context.Context, req *ReqReceiptReader) (*handling.ResponseItems[ResReceiptReader], error)
}
