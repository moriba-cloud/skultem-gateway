package guardian

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/domain"
	"github.com/moriba-cloud/skultem-management/domain/core"
)

type (
	Domain struct {
		*ddd.Aggregation
		givenNames string
		familyName string
		profession string
		email      string
		region     string
		chiefdom   string
		district   string
		city       string
		street     string
		school     string
		phones     []domain.Phone
	}
	Reference struct {
		Id         string
		GivenNames string
		FamilyName string
		Profession string
		Email      string
		Region     string
		Chiefdom   string
		District   string
		City       string
		Street     string
		Phones     []domain.Phone
	}
	Args struct {
		Aggregation ddd.AggregationArgs
		GivenNames  string
		FamilyName  string
		Profession  string
		Email       string
		Region      string
		Chiefdom    string
		District    string
		City        string
		Street      string
		School      string
		Phones      []domain.PhoneArgs
	}
	App interface {
		New(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		One(ctx context.Context, id string) (*ddd.Response[Domain], error)
		ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List(ctx context.Context) (*ddd.Response[core.Option], error)
		Edit(ctx context.Context, args Args) (*ddd.Response[Domain], error)
		Remove(ctx context.Context, id string) (*ddd.Response[Domain], error)
	}
	Repo interface {
		Save(args *Domain) (*Domain, error)
		OneById(id string) (*Domain, error)
		ListByPage(args ddd.PaginationArgs) (*ddd.Response[Domain], error)
		List() (*ddd.Response[core.Option], error)
		Remove(args *Domain) (*Domain, error)
	}
)

func (s *Domain) GivenNames() string {
	return s.givenNames
}

func (s *Domain) FamilyName() string {
	return s.familyName
}

func (s *Domain) Profession() string {
	return s.profession
}

func (s *Domain) Region() string {
	return s.region
}

func (s *Domain) Chiefdom() string {
	return s.chiefdom
}

func (s *Domain) District() string {
	return s.district
}

func (s *Domain) City() string {
	return s.city
}

func (s *Domain) Street() string {
	return s.street
}

func (s *Domain) School() string {
	return s.school
}

func (s *Domain) Email() string {
	email := "no email"
	if len(s.email) > 0 {
		email = s.email
	}
	return email
}

func (s *Domain) Phones() []domain.Phone {
	return s.phones
}

func (s *Domain) ActivePhone() domain.Phone {
	var phone domain.Phone

	for _, o := range s.phones {
		if o.Primary() {
			phone = o
			break
		}
	}

	return phone
}

func (s *Domain) Args() *Args {
	phones := make([]domain.PhoneArgs, 0)

	return &Args{
		Aggregation: ddd.AggregationArgs{
			Id:        s.ID(),
			State:     s.State(),
			CreatedAt: s.CreatedAt(),
			UpdatedAt: s.UpdatedAt(),
		},
		GivenNames: s.givenNames,
		FamilyName: s.familyName,
		Profession: s.profession,
		Email:      s.email,
		Region:     s.region,
		Chiefdom:   s.chiefdom,
		District:   s.district,
		City:       s.city,
		Street:     s.street,
		Phones:     phones,
	}
}
