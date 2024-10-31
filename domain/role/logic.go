package role

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"strings"
)

func New(args Args) (*Domain, error) {
	if err := validation(args); err != nil {
		return nil, err
	}
	return &Domain{
		Aggregation: ddd.NewAggregation(),
		name:        args.Name,
		description: args.Description,
		permissions: args.Permissions,
		school:      args.School,
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
		name:        args.Name,
		description: args.Description,
		permissions: args.Permissions,
		school:      args.School,
	}, nil
}

func validation(args Args) error {
	errors := make([]string, 0)

	if len(args.Name) == 0 {
		errors = append(errors, "name is required")
	}

	if len(args.Description) == 0 {
		errors = append(errors, "end is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}
