package year

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
)

type (
	Domain struct {
		*ddd.Aggregation
		start int64
		end   int64
	}
	Reference struct {
		ID    string
		Start int64
		End   int64
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		Start       int64
		End         int64
	}
	Event struct {
		ID    string
		Year  string
		State string
	}
	Bus interface {
		TakeOff(ctx context.Context, args Domain)
	}
	App interface {
		New(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		One(ctx context.Context, id string) (*ddd.Response[Domain], error)
		ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List(ctx context.Context) (*ddd.Response[core.Option], error)
	}
	Service interface {
		Save(ctx context.Context, args Domain) (*ddd.Response[Domain], error)
		OneById(ctx context.Context, id string) (*ddd.Response[Domain], error)
		ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List(ctx context.Context) (*ddd.Response[core.Option], error)
	}
)

func (d Domain) Start() int64 {
	return d.start
}

func (d Domain) End() int64 {
	return d.end
}
