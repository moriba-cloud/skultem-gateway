package year

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	"time"
)

type (
	Year struct {
		ID        string
		Start     int32
		End       int32
		IsActive  bool
		State     string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

func (y Year) Domain() (*year.Domain, error) {
	return year.Existing(year.Args{
		Aggregation: ddd.AggregationArgs{
			Id:        y.ID,
			State:     ddd.State(y.State),
			CreatedAt: &y.CreatedAt,
			UpdatedAt: &y.UpdatedAt,
		},
		Start: y.Start,
		End:   y.End,
	})
}

func (y Year) Reference() core.Reference {
	return core.Reference{
		Id:    y.ID,
		Value: fmt.Sprintf("%d - %d", y.Start, y.End),
	}
}

func Model(a *year.Domain) Year {
	return Year{
		ID:        a.ID(),
		Start:     a.Start(),
		End:       a.End(),
		IsActive:  a.IsActive(),
		State:     string(a.State()),
		CreatedAt: *a.CreatedAt(),
		UpdatedAt: *a.UpdatedAt(),
	}
}

func (y Year) TableName() string {
	return "academic_year"
}
