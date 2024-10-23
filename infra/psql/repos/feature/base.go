package feature

import (
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/feature"
	"time"
)

type (
	Feature struct {
		ID          string
		Name        string
		Description string
		State       string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)

func (f Feature) Domain() (*feature.Domain, error) {
	return feature.Existing(feature.Args{
		Aggregation: ddd.AggregationArgs{
			Id:        f.ID,
			State:     ddd.State(f.State),
			CreatedAt: &f.CreatedAt,
			UpdatedAt: &f.UpdatedAt,
		},
		Name:        f.Name,
		Description: f.Description,
	})
}

func (f Feature) Reference() core.Reference {
	return core.Reference{
		Id:    f.ID,
		Value: f.Name,
	}
}

func Model(args *feature.Domain) Feature {
	return Feature{
		ID:          args.ID(),
		Name:        args.Name(),
		Description: args.Description(),
		State:       string(args.State()),
		CreatedAt:   *args.CreatedAt(),
		UpdatedAt:   *args.UpdatedAt(),
	}
}
