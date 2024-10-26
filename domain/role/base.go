package role

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
)

type (
	Domain struct {
		*ddd.Aggregation
		name        string
		description string
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		Name        string
		Description string
	}
	App interface {
		New(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		FindById(id string) (*ddd.Response[Domain], error)
		ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List(ctx context.Context) (*ddd.Response[core.Option], error)
		Update(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		Remove(ctx context.Context, id string) (*ddd.Response[Domain], error)
	}
	Repo interface {
		Save(args Domain) (*Domain, error)
		Check(name string) (*Domain, error)
		FindById(id string) (*Domain, error)
		ListByPage(args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List() (*ddd.Response[core.Option], error)
		Remove(args Domain) (*Domain, error)
	}
)

func (d *Domain) Name() string {
	return d.name
}

func (d *Domain) Description() string {
	return d.description
}

func (d *Domain) Update(args Args) error {
	if err := validation(args); err != nil {
		return err
	}
	d.name = args.Name
	d.description = args.Description
	return nil
}
