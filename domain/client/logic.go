package client

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/utils/stn"
	"strings"
)

func New(args Args) (*Domain, error) {
	if err := validation(args); err != nil {
		return nil, err
	}
	return &Domain{
		Aggregation: ddd.NewAggregation(),
		key:         stn.Key(args.Key),
		value:       stn.Clean(args.Value),
		batch:       args.Batch,
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
		key:         stn.Key(args.Key),
		value:       stn.Clean(args.Value),
		batch:       args.Batch,
	}, nil
}

func validation(args Args) error {
	errors := make([]string, 0)

	if args.Key == "" {
		errors = append(errors, "key is required")
	}

	if args.Value == "" {
		errors = append(errors, "value is required")
	}

	if args.Batch == "" {
		errors = append(errors, "batch is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}
