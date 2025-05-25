package usageUnitRepository

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitProto"
	"github.com/hifat/mallow-sale-api/pkg/rpc"
	"google.golang.org/grpc/metadata"
)

type IUsageUnitGRPCRepository interface {
	FindIn(ctx context.Context, filter usageUnit.FilterReq) ([]usageUnit.UsageUnit, error)
}

type grpcRepository struct {
	grpcClient rpc.GrpcClient
}

func NewGRPC(grpcClient rpc.GrpcClient) IUsageUnitGRPCRepository {
	return &grpcRepository{grpcClient}
}

func (r *grpcRepository) FindIn(ctx context.Context, filter usageUnit.FilterReq) ([]usageUnit.UsageUnit, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("x-api-key", "sdfsdf"))
	result, err := r.grpcClient.UsageUnit().FindIn(ctx, &usageUnitProto.InFilter{
		Codes: filter.Codes,
	})
	if err != nil {
		return nil, err
	}

	usageUnits := make([]usageUnit.UsageUnit, 0, len(result.Items))
	for _, v := range result.Items {
		_usageUnit := usageUnit.UsageUnit{
			Code: v.Code,
			Name: v.Name,
		}

		usageUnits = append(usageUnits, _usageUnit)
	}

	return usageUnits, nil
}
