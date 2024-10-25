package permission

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
		feature:     args.Feature,
		create:      args.Create,
		read:        args.Read,
		readAll:     args.ReadAll,
		edit:        args.Edit,
		delete:      args.Delete,
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
		feature:     args.Feature,
		create:      args.Create,
		read:        args.Read,
		readAll:     args.ReadAll,
		edit:        args.Edit,
		delete:      args.Delete,
	}, nil
}

func validation(args Args) error {
	errors := make([]string, 0)

	if len(args.Feature) == 0 {
		errors = append(errors, "feature is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}
