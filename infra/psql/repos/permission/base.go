package permission

import (
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/permission"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos/feature"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos/role"
	"time"
)

type (
	Permission struct {
		ID        string
		Create    bool
		Read      bool
		ReadAll   bool
		Edit      bool
		Delete    bool
		State     string
		RoleId    string `gorm:"index"`
		FeatureId string `gorm:"index"`
		Role      role.Role
		Feature   feature.Feature
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

func (p Permission) Domain() (*permission.Domain, error) {
	return permission.Existing(permission.Args{
		Aggregation: ddd.AggregationArgs{
			Id:        p.ID,
			State:     ddd.State(p.State),
			CreatedAt: &p.CreatedAt,
			UpdatedAt: &p.UpdatedAt,
		},
		Feature: p.Feature.Reference(),
		Create:  p.Create,
		Read:    p.Read,
		ReadAll: p.ReadAll,
		Edit:    p.Edit,
		Delete:  p.Delete,
	})
}

func Model(args *permission.Domain, role string) Permission {
	return Permission{
		ID:        args.ID(),
		RoleId:    role,
		FeatureId: args.Feature().Id,
		Create:    args.Create(),
		Read:      args.Read(),
		ReadAll:   args.ReadAll(),
		Edit:      args.Edit(),
		Delete:    args.Delete(),
		State:     string(args.State()),
		CreatedAt: *args.CreatedAt(),
		UpdatedAt: *args.UpdatedAt(),
	}
}
