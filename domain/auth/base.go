package auth

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
)

type (
	Domain struct {
		access  string
		refresh string
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		Refresh     string
	}
	User struct {
		Id         string
		GivenNames string
		FamilyName string
		Phone      int
		Email      string
		Role       core.Reference
		School     string
		State      string
	}
	App interface {
		Login(ctx context.Context, email string, password string) (*ddd.Response[Domain], error)
		Access(ctx context.Context) (*ddd.Response[Domain], error)
		Me(ctx context.Context) (*ddd.Response[User], error)
	}
)

func (d *Domain) Access() string {
	return d.access
}

func (d *Domain) Refresh() string {
	return d.refresh
}
