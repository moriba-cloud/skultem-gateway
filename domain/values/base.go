package values

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
)

type (
	Reference struct {
		Id    string
		Value string
	}
	Batch  string
	Domain struct {
		*ddd.Aggregation
		key    string
		value  string
		school string
		batch  Batch
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		Key         string
		Value       string
		School      string
		Batch       Batch
	}
	App interface {
		New(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		List(ctx context.Context) (*ddd.Response[core.Option], error)
		ListByGroup(ctx context.Context, batch Batch) (*ddd.Response[core.Option], error)
		ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[Domain], error)
	}
	Service interface {
		Save(ctx context.Context, args Domain) (*ddd.Response[Domain], error)
		ListByBatch(ctx context.Context, group Batch) (*ddd.Response[core.Option], error)
		ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List(ctx context.Context) (*ddd.Response[core.Option], error)
	}
)

const (
	DESIGNATION  Batch = "DESIGNATION"
	SECTION      Batch = "SECTION"
	SUBJECT      Batch = "SUBJECT"
	RELIGION     Batch = "RELIGION"
	PAYMENT_PLAN Batch = "PAYMENT_PLAN"
)

func (d *Domain) Key() string {
	return d.key
}

func (d *Domain) School() string {
	return d.school
}

func (d *Domain) Value() string {
	return d.value
}

func (d *Domain) Batch() Batch {
	return d.batch
}

func (d *Domain) Args() *Args {
	return &Args{
		Aggregation: ddd.AggregationArgs{
			Id:        d.ID(),
			State:     d.State(),
			CreatedAt: d.CreatedAt(),
			UpdatedAt: d.UpdatedAt(),
		},
		Key:    d.key,
		Value:  d.value,
		Batch:  d.batch,
		School: d.school,
	}
}
