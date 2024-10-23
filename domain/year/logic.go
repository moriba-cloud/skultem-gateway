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
		isActive:    args.IsActive,
	}, nil
}

func validation(args Args) error {
	errors := make([]string, 0)

	start := fmt.Sprintf("%d", args.Start)
	end := fmt.Sprintf("%d", args.End)
	diff := args.End - args.Start

	if len(start) == 0 {
		errors = append(errors, "start is required")
	}

	if len(end) == 0 {
		errors = append(errors, "end is required")
	}

	if len(start) < 4 || len(start) < 4 {
		errors = append(errors, "invalid start date")
	}

	if len(end) < 4 || len(end) < 4 {
		errors = append(errors, "invalid end date")
	}

	if len(end) == 0 {
		errors = append(errors, "end is required")
	}

	if args.Start > args.End {
		return fiber.NewError(fiber.StatusBadRequest, "start must be lesser than end year")
	}

	if diff < 1 || diff > 1 {
		errors = append(errors, "year must be between 1 year")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}
