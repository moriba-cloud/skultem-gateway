package role

import (
	"github.com/moriba-build/ose/ddd"
)

func New(args Args) (*Domain, error) {
	return &Domain{
		Aggregation: ddd.NewAggregation(),
		create:      args.Create,
		read:        args.Read,
		readAll:     args.ReadAll,
		edit:        args.Edit,
		delete:      args.Delete,
	}, nil
}

func Existing(args Args) (*Domain, error) {
	Aggregation, err := ddd.ExistingAggregation(args.Aggregation)
	if err != nil {
		return nil, err
	}

	return &Domain{
		Aggregation: Aggregation,
		create:      args.Create,
		read:        args.Read,
		readAll:     args.ReadAll,
		edit:        args.Edit,
		delete:      args.Delete,
	}, nil
}
