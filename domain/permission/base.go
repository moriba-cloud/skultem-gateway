package permission

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
)

type (
	Domain struct {
		*ddd.Aggregation
		feature core.Reference
		create  bool
		read    bool
		readAll bool
		edit    bool
		delete  bool
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		Feature     core.Reference
		Create      bool
		Read        bool
		ReadAll     bool
		Edit        bool
		Delete      bool
	}
	App interface {
		Update(ctx context.Context, args []*Args, role string) (*ddd.Response[Domain], error)
	}
	Repo interface {
		Save(args []*Domain, roleId string) (*ddd.Response[Domain], error)
		RolePermissions(role string) (*ddd.Response[Domain], error)
		Check(feature string, role string) (*Domain, error)
	}
)

func (d *Domain) Feature() core.Reference {
	return d.feature
}

func (d *Domain) Create() bool {
	return d.create
}

func (d *Domain) Read() bool {
	return d.read
}

func (d *Domain) ReadAll() bool {
	return d.readAll
}

func (d *Domain) Edit() bool {
	return d.edit
}

func (d *Domain) Delete() bool {
	return d.delete
}

func (d *Domain) Update(args Args) error {
	if err := validation(args); err != nil {
		return err
	}

	d.feature = args.Feature
	d.create = args.Create
	d.readAll = args.ReadAll
	d.read = args.Read
	d.edit = args.Edit
	d.delete = args.Delete

	return nil
}
