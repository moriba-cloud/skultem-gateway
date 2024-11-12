package management

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/config"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	academicv1 "github.com/moriba-cloud/skultem-gateway/infra/management/grpc/gen/go/academic/v1"
	commonv1 "github.com/moriba-cloud/skultem-gateway/infra/management/grpc/gen/go/common/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

type (
	yearService struct {
		conn   *grpc.ClientConn
		logger *zap.Logger
	}
)

func ToDomain(record *academicv1.Academic) (*year.Domain, error) {
	createdAt := record.CreatedAt.AsTime()
	updatedAt := record.UpdatedAt.AsTime()

	start, err := strconv.Atoi(record.GetStart())
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(record.GetEnd())
	if err != nil {
		return nil, err
	}

	return year.Existing(year.Args{
		Aggregation: ddd.AggregationArgs{
			Id:        record.Id,
			State:     ddd.State(record.State),
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		},
		Start: int64(start),
		End:   int64(end),
	})
}

func (i *yearService) Save(ctx context.Context, args year.Domain) (*ddd.Response[year.Domain], error) {
	conn := academicv1.NewAcademicServiceClient(i.conn)
	res, err := conn.Create(ctx, &academicv1.CreateRequest{
		Start: args.Start(),
		End:   args.End(),
	})
	if err != nil {
		return nil, CleanError(err.Error())
	}

	record, err := ToDomain(res.GetRecord())
	if err != nil {
		return nil, err
	}

	i.logger.Info(fmt.Sprintf("saved academic year with id: %s", record.ID()))
	return ddd.NewResponse(ddd.ResponseArgs[year.Domain]{
		Record: record,
	}), nil
}

func (i *yearService) OneById(ctx context.Context, id string) (*ddd.Response[year.Domain], error) {
	conn := academicv1.NewAcademicServiceClient(i.conn)
	res, err := conn.Read(ctx, &academicv1.ReadRequest{
		Id: id,
	})
	if err != nil {
		return nil, CleanError(err.Error())
	}

	record, err := ToDomain(res.GetRecord())
	if err != nil {
		return nil, err
	}

	i.logger.Info(fmt.Sprintf("fetched academic year with id: %s", id))
	return ddd.NewResponse(ddd.ResponseArgs[year.Domain]{
		Record: record,
	}), nil
}

func (i *yearService) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[year.Domain], error) {
	conn := academicv1.NewAcademicServiceClient(i.conn)
	ctx = GrpcMetadata(ctx)
	res, err := conn.ReadAll(ctx, &academicv1.ReadAllRequest{Query: &commonv1.Query{
		Limit: uint32(args.Limit),
		Page:  uint64(args.Page),
	}})

	if err != nil {
		return nil, CleanError(err.Error())
	}

	records := make([]*year.Domain, 0)
	for _, o := range res.Records {
		record, err := ToDomain(o)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	i.logger.Info(fmt.Sprintf("fetch %d user by limit: %d and page: %d", len(records), args.Limit, args.Page))
	return ddd.NewResponse(ddd.ResponseArgs[year.Domain]{
		Pagination: ddd.PaginationArgs{
			Limit: int(res.Option.Pagination.GetLimit()),
			Page:  int(res.Option.Pagination.GetPage()),
			Size:  int(res.Option.Pagination.GetSize()),
			Pages: int(res.Option.Pagination.GetPages()),
		},
		Records: records,
	}), nil
}

func (i *yearService) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	//TODO implement me
	panic("implement me")
}

func NewYear(logger *zap.Logger) year.Service {
	addr := config.NewEnvs().EnvStr("MANAGEMENT_SERVER_ADDR")
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &yearService{
		conn:   conn,
		logger: logger,
	}
}
