package school

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/utils/stn"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"strings"
)

func New(args Args) (*Domain, error) {
	if err := validation(args); err != nil {
		return nil, err
	}
	password, err := core.GeneratePassword()
	if err != nil {
		return nil, err
	}

	return &Domain{
		Aggregation: ddd.NewAggregation(),
		givenNames:  stn.Clean(args.GivenNames),
		familyName:  stn.Clean(args.FamilyName),
		phone:       args.Phone,
		email:       stn.Clean(args.Email),
		role:        args.Role,
		password:    *password,
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

	return &Domain{
		Aggregation: Aggregation,
		givenNames:  stn.Clean(args.GivenNames),
		familyName:  stn.Clean(args.FamilyName),
		phone:       args.Phone,
		email:       stn.Clean(args.Email),
		role:        args.Role,
		password:    args.Password,
	}, nil
}

func validation(args Args) error {
	errors := make([]string, 0)

	if len(args.GivenNames) < 0 {
		errors = append(errors, "given names is required")
	}

	if len(args.FamilyName) < 0 {
		errors = append(errors, "family name is required")
	}

	if len(args.Email) < 0 {
		errors = append(errors, "email is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}
