package school

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/domain"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
)

type (
	Reference struct {
		Id    string
		Value string
	}
	Domain struct {
		*ddd.Aggregation
		name     string
		email    string
		region   string
		chiefdom string
		district string
		city     string
		street   string
		phones   []domain.Phone
		owner    core.Reference
	}
	AuthArgs struct {
		Email    string
		Password string
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		Name        string
		Email       string
		Region      string
		Chiefdom    string
		District    string
		City        string
		Street      string
		Phones      []string
	}
	App interface {
		New(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		FindById(ctx context.Context, id string) (*ddd.Response[Domain], error)
		Update(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		List(ctx context.Context) (*ddd.Response[core.Option], error)
		ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		Remove(ctx context.Context, id string) (*ddd.Response[Domain], error)
	}
	Repo interface {
		Save(args Domain) (*Domain, error)
		FindById(id string) (*ddd.Response[Domain], error)
		FindByEmail(email string) (*ddd.Response[Domain], error)
		Check(phone int, email string) (*Domain, error)
		List() (*ddd.Response[core.Option], error)
		ListByPage(args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		Remove(args Domain) (*Domain, error)
	}
)

func (d *Domain) Name() string {
	return d.name
}

func (d *Domain) Region() string {
	return d.region
}

func (d *Domain) District() string {
	return d.district
}

func (d *Domain) Chiefdom() string {
	return d.chiefdom
}

func (d *Domain) City() string {
	return d.city
}

func (d *Domain) Street() string {
	return d.street
}

func (d *Domain) Owner() core.Reference {
	return d.owner
}

func (d *Domain) Phones() []domain.Phone {
	return d.phones
}

func (d *Domain) Update(args Args) error {
	return nil
}
