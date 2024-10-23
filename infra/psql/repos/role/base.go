package role

import (
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/role"
	"time"
)

type (
	Role struct {
		ID          string
		Name        string
		Description string
		State       string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)

func (r Role) Domain() (*role.Domain, error) {
	return role.Existing(role.Args{
		Aggregation: ddd.AggregationArgs{
			Id:        r.ID,
			State:     ddd.State(r.State),
			CreatedAt: &r.CreatedAt,
			UpdatedAt: &r.UpdatedAt,
		},
		Name:        r.Name,
		Description: r.Description,
	})
}

func (r Role) Reference() core.Reference {
	return core.Reference{
		Id:    r.ID,
		Value: r.Name,
	}
}

func Model(args *role.Domain) Role {
	return Role{
		ID:          args.ID(),
		Name:        args.Name(),
		Description: args.Description(),
		State:       string(args.State()),
		CreatedAt:   *args.CreatedAt(),
		UpdatedAt:   *args.UpdatedAt(),
	}
}
