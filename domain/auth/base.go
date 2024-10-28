package auth

import (
	"context"
	"github.com/moriba-build/ose/ddd"
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
	App interface {
		Login(ctx context.Context, email string, password string) (*ddd.Response[Domain], error)
		Access(ctx context.Context, refresh string) (*ddd.Response[Domain], error)
	}
)

func (d *Domain) Access() string {
	return d.access
}

func (d *Domain) Refresh() string {
	return d.refresh
}
