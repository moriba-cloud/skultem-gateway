package user

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	core2 "github.com/moriba-build/ose/ddd/core"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos/role"
	"time"
)

type (
	User struct {
		ID             string
		GivenNames     string
		FamilyName     string
		Email          string
		Phone          int
		Role           role.Role
		RoleId         string
		Password       string
		PasswordTxt    string
		PasswordStatus string
		State          string
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)

func (u User) Domain() (*user.Domain, error) {
	phone, err := core2.NewPhone(u.Phone)
	if err != nil {
		return nil, err
	}

	return user.Existing(user.Args{
		Aggregation: ddd.AggregationArgs{
			Id:        u.ID,
			State:     ddd.State(u.State),
			CreatedAt: &u.CreatedAt,
			UpdatedAt: &u.UpdatedAt,
		},
		GivenNames: u.GivenNames,
		FamilyName: u.FamilyName,
		Email:      u.Email,
		Role:       u.Role.Reference(),
		Password: core.Password{
			Value: u.PasswordTxt,
			Hash:  u.Password,
			State: core.PasswordState(u.PasswordStatus),
		},
		Phone: phone.Phone(),
	})
}

func (u User) Reference() core.Reference {
	return core.Reference{
		Id:    u.ID,
		Value: fmt.Sprintf("%s %s (%d)", u.GivenNames, u.FamilyName, u.Phone),
	}
}

func Model(args *user.Domain) User {
	return User{
		ID:             args.ID(),
		GivenNames:     args.GivenNames(),
		FamilyName:     args.FamilyName(),
		Email:          args.Email(),
		Phone:          args.Phone(),
		RoleId:         args.Role().Id,
		Password:       args.Password().Hash,
		PasswordTxt:    args.Password().Value,
		PasswordStatus: string(args.Password().State),
		State:          string(args.State()),
		CreatedAt:      *args.CreatedAt(),
		UpdatedAt:      *args.UpdatedAt(),
	}
}
