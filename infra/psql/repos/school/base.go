package school

import (
	"github.com/moriba-build/ose/ddd"
	ose "github.com/moriba-build/ose/domain"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/school"
	domainUser "github.com/moriba-cloud/skultem-gateway/domain/user"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos/user"
	"time"
)

type (
	School struct {
		ID        string
		Name      string
		Domain    string
		Email     string
		Region    string
		District  string
		Chiefdom  string
		City      string
		Street    string
		State     string
		Phones    []Phone
		Users     []user.User
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

func (s *School) ToDomain() (*school.Domain, error) {
	phones := make([]ose.PhoneArgs, len(s.Phones))
	for i, p := range s.Phones {
		phones[i] = ose.PhoneArgs{
			Aggregation: ddd.AggregationArgs{
				Id:        p.ID,
				State:     ddd.State(p.State),
				CreatedAt: &p.CreatedAt,
				UpdatedAt: &p.UpdatedAt,
			},
			Primary: p.Primary,
			Number:  p.Number,
			Country: p.Country,
		}
	}

	users := make([]domainUser.Args, len(s.Users))
	for i, u := range s.Users {
		usr, err := u.Args()
		if err != nil {
			return nil, err
		}

		users[i] = *usr
	}

	return school.Existing(school.Args{
		Aggregation: ddd.AggregationArgs{
			Id:        s.ID,
			State:     ddd.State(s.State),
			CreatedAt: &s.CreatedAt,
			UpdatedAt: &s.UpdatedAt,
		},
		Name:     s.Name,
		Email:    s.Email,
		Region:   s.Region,
		Chiefdom: s.Chiefdom,
		District: s.District,
		City:     s.City,
		Street:   s.Street,
		Domain:   s.Domain,
		Phones:   phones,
		Users:    users,
	})
}

func (s *School) Reference() core.Reference {
	return core.Reference{
		Id:    s.ID,
		Value: s.Name,
	}
}

func Model(args *school.Domain) School {
	return School{
		ID:        args.ID(),
		Name:      args.Name(),
		Domain:    args.Domain(),
		Email:     args.Email(),
		Region:    args.Region(),
		District:  args.District(),
		Chiefdom:  args.Chiefdom(),
		City:      args.City(),
		Street:    args.Street(),
		State:     string(args.State()),
		CreatedAt: *args.CreatedAt(),
		UpdatedAt: *args.UpdatedAt(),
	}
}
