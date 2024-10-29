package school

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/utils/stn"
	"github.com/moriba-build/ose/domain"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"strings"
)

func New(args Args) (*Domain, error) {
	if err := validation(args); err != nil {
		return nil, err
	}
	Aggregation := ddd.NewAggregation()
	phones := make([]domain.Phone, len(args.Phones))

	for i, p := range args.Phones {
		phone, err := domain.NewPhone(domain.PhoneArgs{
			Primary: false,
			Number:  p.Number,
		})
		if err != nil {
			return nil, err
		}

		phones[i] = *phone
	}

	owner, err := user.New(args.Owner)
	if err != nil {
		return nil, err
	}

	return &Domain{
		Aggregation: Aggregation,
		name:        stn.Clean(args.Name),
		domain:      strings.ToLower(stn.Key(args.Domain)),
		email:       stn.Clean(args.Email),
		region:      stn.Clean(args.Region),
		district:    stn.Clean(args.District),
		chiefdom:    stn.Clean(args.Chiefdom),
		city:        stn.Clean(args.City),
		street:      stn.Clean(args.Street),
		phones:      phones,
		owner:       owner,
	}, nil
}

func Existing(args Args) (*Domain, error) {
	if err := validation(args); err != nil {
		return nil, err
	}
	phones := make([]domain.Phone, len(args.Phones))

	for i, p := range args.Phones {
		phone, err := domain.NewPhone(domain.PhoneArgs{
			Primary: false,
			Number:  p.Number,
		})
		if err != nil {
			return nil, err
		}

		phones[i] = *phone
	}

	owner, err := user.New(args.Owner)
	if err != nil {
		return nil, err
	}

	Aggregation, err := ddd.ExistingAggregation(args.Aggregation)
	if err != nil {
		return nil, err
	}

	return &Domain{
		Aggregation: Aggregation,
		name:        stn.Clean(args.Name),
		domain:      strings.ToLower(stn.Key(args.Domain)),
		email:       stn.Clean(args.Email),
		region:      stn.Clean(args.Region),
		district:    stn.Clean(args.District),
		chiefdom:    stn.Clean(args.Chiefdom),
		city:        stn.Clean(args.City),
		street:      stn.Clean(args.Street),
		phones:      phones,
		owner:       owner,
	}, nil
}

func validation(args Args) error {
	errors := make([]string, 0)

	if len(args.Name) < 0 {
		errors = append(errors, "name is required")
	}

	if len(args.Domain) < 0 {
		errors = append(errors, "domain is required")
	}

	if len(args.Email) < 0 {
		errors = append(errors, "email is required")
	}

	if len(args.Region) < 0 {
		errors = append(errors, "region is required")
	}

	if len(args.District) < 0 {
		errors = append(errors, "district is required")
	}

	if len(args.Chiefdom) < 0 {
		errors = append(errors, "chiefdom is required")
	}

	if len(args.City) < 0 {
		errors = append(errors, "city is required")
	}

	if len(args.Street) < 0 {
		errors = append(errors, "street is required")
	}

	if len(args.Phones) < 0 {
		errors = append(errors, "minimum of 1 phone is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}
