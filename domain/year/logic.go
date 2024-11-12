package year

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd"
	"strings"
)

func New(args Args) (*Domain, error) {
	if err := validation(args); err != nil {
		return nil, err
	}
	return &Domain{
		Aggregation: ddd.NewAggregation(),
		start:       args.Start,
		end:         args.End,
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
		start:       args.Start,
		end:         args.End,
	}, nil
}

func validation(args Args) error {
	errors := make([]string, 0)

	if args.Start == 0 {
		errors = append(errors, "start is required")
	}

	if args.End == 0 {
		errors = append(errors, "end is required")
	}

	different := args.End - args.Start

	if args.End < args.Start {
		return fiber.NewError(fiber.StatusBadRequest, "start must be greater than end year")
	}

	if different < 1 || different > 1 {
		errors = append(errors, "year must be between 1 year")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}
