package user

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	core2 "github.com/moriba-build/ose/ddd/core"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
)

type (
	Reference struct {
		Id    string
		Value string
	}
	Domain struct {
		*ddd.Aggregation
		givenNames string
		familyName string
		phone      int
		email      string
		role       core.Reference
		password   core.Password
		school     string
	}
	AuthArgs struct {
		Email    string
		Password string
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		GivenNames  string
		FamilyName  string
		Phone       int
		Email       string
		School      string
		Role        core.Reference
		Password    core.Password
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
		CheckByEmail(email string) (*Domain, error)
		CheckByPhone(phone int) (*Domain, error)
		List() (*ddd.Response[core.Option], error)
		ListByPage(args ddd.PaginationArgs, school string) (*ddd.Response[Domain], error)
		Remove(args Domain) (*Domain, error)
	}
)

func (d *Domain) GivenNames() string {
	return d.givenNames
}

func (d *Domain) FamilyName() string {
	return d.familyName
}

func (d *Domain) Phone() int {
	return d.phone
}

func (d *Domain) Email() string {
	return d.email
}

func (d *Domain) School() string {
	return d.school
}

func (d *Domain) Role() core.Reference {
	return d.role
}

func (d *Domain) Password() core.Password {
	return d.password
}

func (d *Domain) Update(args Args) error {
	phone, err := core2.NewPhone(args.Phone)
	if err != nil {
		return err
	}

	d.givenNames = args.GivenNames
	d.familyName = args.FamilyName
	d.phone = phone.Phone()
	d.email = args.Email
	d.role = args.Role
	d.password = args.Password

	return nil
}

func (d *Domain) ForgetPassword() error {
	password, err := core.GeneratePassword()
	if err != nil {
		return err
	}

	d.password = *password
	return nil
}

func (d *Domain) Reference() core.Reference {
	return core.Reference{
		Id:    d.Aggregation.ID(),
		Value: fmt.Sprintf("%s %s (%d)", d.givenNames, d.familyName, d.phone),
	}
}

func (d *Domain) Option() *core.Option {
	return &core.Option{
		Label: d.Aggregation.ID(),
		Value: fmt.Sprintf("%s %s (%d)", d.givenNames, d.familyName, d.phone),
	}
}

func (d *Domain) Args() *Args {
	return &Args{
		Aggregation: ddd.AggregationArgs{
			Id:        d.ID(),
			State:     d.State(),
			CreatedAt: d.CreatedAt(),
			UpdatedAt: d.UpdatedAt(),
		},
		GivenNames: d.GivenNames(),
		FamilyName: d.FamilyName(),
		Phone:      d.Phone(),
		Email:      d.Email(),
		Role:       d.Role(),
		Password:   d.Password(),
	}
}
