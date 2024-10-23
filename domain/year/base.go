package year

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
)

type (
	Domain struct {
		*ddd.Aggregation
		start    int32
		end      int32
		isActive bool
	}
	Reference struct {
		ID    string
		Start int32
		End   int32
	}
	Event struct {
		ID    string
		Year  string
		State string
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		Start       int32
		End         int32
		IsActive    bool
	}
	Bus interface {
		TakeOff(ctx context.Context, args Domain)
	}
	App interface {
		New(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		One(ctx context.Context, id string) (*ddd.Response[Domain], error)
		ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List(ctx context.Context) (*ddd.Response[core.Option], error)
		Active(ctx context.Context, args ddd.FindByArgs) (*ddd.Response[Domain], error)
	}
	Repo interface {
		Save(args Domain) (*Domain, error)
		Check(start int32, end int32) bool
		OneById(id string) (*Domain, error)
		ListByPage(args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List() (*ddd.Response[core.Option], error)
	}
)

func (d Domain) Start() int32 {
	return d.start
}

func (d Domain) End() int32 {
	return d.end
}

func (d Domain) IsActive() bool {
	return d.isActive
}
