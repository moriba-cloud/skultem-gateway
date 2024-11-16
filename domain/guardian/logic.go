package guardian

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/utils/stn"
	"github.com/moriba-build/ose/domain"
	"strings"
)

func New(args Args) (*Domain, error) {
	if err := validation(args); err != nil {
		return nil, err
	}

	Aggregation := ddd.NewAggregation()

	phones := make([]domain.Phone, 0)
	for _, phone := range args.Phones {
		p, err := domain.NewPhone(domain.PhoneArgs{
			Primary: true,
			Number:  phone.Number,
			Owner:   Aggregation.ID(),
		})
		if err != nil {
			return nil, err
		}

		phones = append(phones, *p)
	}

	return &Domain{
		Aggregation: Aggregation,
		givenNames:  stn.Clean(args.GivenNames),
		familyName:  stn.Clean(args.FamilyName),
		profession:  stn.Clean(args.Profession),
		email:       stn.Clean(args.Email),
		region:      stn.Clean(args.Region),
		chiefdom:    stn.Clean(args.Chiefdom),
		district:    stn.Clean(args.District),
		city:        stn.Clean(args.City),
		street:      stn.Clean(args.Street),
		school:      args.School,
		phones:      phones,
	}, nil
}

func Existing(args Args) (*Domain, error) {
	if err := validation(args); err != nil {
		return nil, err
	}

	Aggregation, err := ddd.ExistingAggregation(args.Aggregation)
	if err != nil {
		return nil, err
	}

	phones := make([]domain.Phone, 0)
	for _, phone := range args.Phones {
		p, err := domain.ExistingPhone(phone)
		if err != nil {
			return nil, err
		}

		phones = append(phones, *p)
	}

	return &Domain{
		Aggregation: Aggregation,
		givenNames:  stn.Clean(args.GivenNames),
		familyName:  stn.Clean(args.FamilyName),
		profession:  args.Profession,
		email:       stn.Clean(args.Email),
		region:      stn.Clean(args.Region),
		chiefdom:    stn.Clean(args.Chiefdom),
		district:    stn.Clean(args.District),
		city:        stn.Clean(args.City),
		street:      stn.Clean(args.Street),
		school:      args.School,
		phones:      phones,
	}, nil
}

func validation(args Args) error {
	errors := make([]string, 0)

	if len(args.GivenNames) == 0 {
		errors = append(errors, "given name is required")
	}

	if len(args.FamilyName) == 0 {
		errors = append(errors, "family name is required")
	}

	if len(args.Profession) == 0 {
		errors = append(errors, "profession is required")
	}

	if len(args.Phones) == 0 {
		errors = append(errors, "minimum of 1 phone is required")
	}

	if len(args.Region) == 0 {
		errors = append(errors, "region is required")
	}

	if len(args.District) == 0 {
		errors = append(errors, "district is required")
	}

	if args.Chiefdom == "" {
		errors = append(errors, "chiefdom is required")
	}

	if args.City == "" {
		errors = append(errors, "city is required")
	}

	if args.Street == "" {
		errors = append(errors, "street is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}
