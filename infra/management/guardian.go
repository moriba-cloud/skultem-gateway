package management

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/config"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/values"
	commonv1 "github.com/moriba-cloud/skultem-gateway/infra/management/grpc/gen/go/common/v1"
	valuev1 "github.com/moriba-cloud/skultem-gateway/infra/management/grpc/gen/go/value/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type (
	valueService struct {
		conn   *grpc.ClientConn
		logger *zap.Logger
	}
)

func ToValueDomain(record *valuev1.Value) (*values.Domain, error) {
	createdAt := record.CreatedAt.AsTime()
	updatedAt := record.UpdatedAt.AsTime()

	return values.Existing(values.Args{
		Aggregation: ddd.AggregationArgs{
			Id:        record.Id,
			State:     ddd.State(record.State),
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		},
		Value: record.GetValue(),
		Batch: values.Batch(record.GetBatch()),
		Key:   record.GetKey(),
	})
}

func (i *valueService) Save(ctx context.Context, args values.Domain) (*ddd.Response[values.Domain], error) {
	ctx = GrpcMetadata(ctx)
	conn := valuev1.NewValueServiceClient(i.conn)

	res, err := conn.Create(ctx, &valuev1.CreateRequest{
		Key:   args.Key(),
		Value: args.Value(),
		Batch: string(args.Batch()),
	})
	if err != nil {
		return nil, CleanError(err.Error())
	}

	record, err := ToValueDomain(res.GetRecord())
	if err != nil {
		return nil, err
	}

	i.logger.Info(fmt.Sprintf("saved value with id: %s", record.ID()))
	return ddd.NewResponse(ddd.ResponseArgs[values.Domain]{
		Record: record,
	}), nil
}

func (i *valueService) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[values.Domain], error) {
	ctx = GrpcMetadata(ctx)

	conn := valuev1.NewValueServiceClient(i.conn)
	ctx = GrpcMetadata(ctx)
	res, err := conn.ReadAll(ctx, &valuev1.ReadAllRequest{Query: &commonv1.Query{
		Limit: uint32(args.Limit),
		Page:  uint64(args.Page),
	}})

	if err != nil {
		return nil, CleanError(err.Error())
	}

	records := make([]*values.Domain, 0)
	for _, o := range res.Records {
		record, err := ToValueDomain(o)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	i.logger.Info(fmt.Sprintf("fetch %d user by limit: %d and page: %d", len(records), args.Limit, args.Page))
	return ddd.NewResponse(ddd.ResponseArgs[values.Domain]{
		Pagination: ddd.PaginationArgs{
			Limit: int(res.Option.Pagination.GetLimit()),
			Page:  int(res.Option.Pagination.GetPage()),
			Size:  int(res.Option.Pagination.GetSize()),
			Pages: int(res.Option.Pagination.GetPages()),
		},
		Records: records,
	}), nil
}

func (i *valueService) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	//TODO implement me
	panic("implement me")
}

func (i *valueService) ListByBatch(ctx context.Context, batch values.Batch) (*ddd.Response[core.Option], error) {
	ctx = GrpcMetadata(ctx)

	conn := valuev1.NewValueServiceClient(i.conn)
	ctx = GrpcMetadata(ctx)
	res, err := conn.ReadByBatch(ctx, &valuev1.ReadByGroupRequest{
		Batch: string(batch),
	})

	if err != nil {
		return nil, CleanError(err.Error())
	}

	records := make([]*core.Option, len(res.Records))
	for i, o := range res.Records {
		records[i] = &core.Option{
			Label: o.Label,
			Value: o.Value,
		}
	}

	i.logger.Info(fmt.Sprintf("fetch %d user by batch: %s", len(records), batch))
	return ddd.NewResponse(ddd.ResponseArgs[core.Option]{
		Records: records,
	}), nil
}

func NewValue(logger *zap.Logger) values.Service {
	addr := config.NewEnvs().EnvStr("MANAGEMENT_SERVER_ADDR")
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &valueService{
		conn:   conn,
		logger: logger,
	}
}
