package role

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/permission"
)

type (
	Domain struct {
		*ddd.Aggregation
		name        string
		description string
		school      string
		permissions []*permission.Domain
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		Name        string
		Description string
		School      string
		Permissions []*permission.Domain
	}
	App interface {
		New(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		FindById(ctx context.Context, id string) (*ddd.Response[Domain], error)
		ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List(ctx context.Context) (*ddd.Response[core.Option], error)
		Update(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		Remove(ctx context.Context, id string) (*ddd.Response[Domain], error)
	}
	Repo interface {
		Save(args Domain) (*Domain, error)
		Check(name string) (*Domain, error)
		FindById(id string) (*Domain, error)
		ListByPage(args ddd.PaginationArgs, school string) (*ddd.Response[Domain], error)
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

func (d *Domain) School() string {
	return d.school
}

func (d *Domain) Permissions() []*permission.Domain {
	return d.permissions
}

func (d *Domain) Update(args Args) {
	if len(args.Name) > 0 {
		d.name = args.Name
	}

	if len(args.Description) > 0 {
		d.description = args.Description
	}

	d.permissions = args.Permissions
}
