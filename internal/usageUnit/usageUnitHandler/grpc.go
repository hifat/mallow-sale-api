package usageUnitHandler

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitProto"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitService"
	"github.com/jinzhu/copier"
)

type usageUnitGRPC struct {
	usageUnitProto.UnimplementedUsageUnitGrpcServiceServer
	usageUnitSrv usageUnitService.IUsageUnitService
}

func NewGRPC(usageUnitSrv usageUnitService.IUsageUnitService) *usageUnitGRPC {
	return &usageUnitGRPC{
		usageUnitSrv: usageUnitSrv,
	}
}

func (h *usageUnitGRPC) FindIn(ctx context.Context, req *usageUnitProto.InFilter) (*usageUnitProto.UsageUnitRes, error) {
	filter := usageUnit.FilterReq{}
	if err := copier.Copy(&filter, req); err != nil {
		return nil, err
	}

	usageUnits, err := h.usageUnitSrv.FindIn(ctx, filter)
	if err != nil {
		return nil, err
	}

	res := make([]*usageUnitProto.UsageUnit, 0, len(usageUnits))
	for _, v := range usageUnits {
		item := &usageUnitProto.UsageUnit{
			Code: v.Code,
			Name: v.Name,
		}

		res = append(res, item)
	}

	return &usageUnitProto.UsageUnitRes{
		Items: res,
	}, nil
}
