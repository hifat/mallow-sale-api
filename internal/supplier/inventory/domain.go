package supplierInventoryModule

import (
	"context"

	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type IRepository interface {
	FindGroupBySupplier(ctx context.Context, query *utilsModule.QueryReq) ([]GroupBySupplierResponse, error)
}

type IService interface {
	FindGroupBySupplier(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[GroupBySupplierResponse], error)
}
