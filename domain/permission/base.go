package role

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
)

type (
	Domain struct {
		*ddd.Aggregation
		role    string
		create  bool
		read    bool
		readAll bool
		edit    bool
		delete  bool
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		Create      bool
		Read        bool
		ReadAll     bool
		Edit        bool
		Delete      bool
	}
	App interface {
		New(ctx context.Context, args Args) (*ddd.Response[Domain], error)
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

func (d *Domain) Create() bool {
	return d.create
}

func (d *Domain) Read() bool {
	return d.create
}

func (d *Domain) ReadAll() bool {
	return d.create
}

func (d *Domain) Edit() bool {
	return d.create
}

func (d *Domain) Delete() bool {
	return d.create
}

func (d *Domain) Update(args Args) error {
	if err := validation(args); err != nil {
		return err
	}
	d.create = args.Create
	d.readAll = args.ReadAll
	d.read = args.Read
	d.edit = args.Edit
	d.delete = args.Delete
	return nil
}
