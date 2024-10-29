package school

import (
	"github.com/moriba-build/ose/ddd"
	ose "github.com/moriba-build/ose/domain"
	"time"
)

type Phone struct {
	ID        string
	Primary   bool
	Number    int
	Country   string
	SchoolId  string `gorm:"index"`
	State     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToPhone(o ose.Phone, schoolId string) Phone {
	return Phone{
		ID:        o.ID(),
		Primary:   o.Primary(),
		Number:    o.Number(),
		Country:   o.Country(),
		SchoolId:  schoolId,
		State:     string(o.State()),
		CreatedAt: *o.CreatedAt(),
		UpdatedAt: *o.UpdatedAt(),
	}
}

func (p Phone) Args() ose.PhoneArgs {
	return ose.PhoneArgs{
		Aggregation: ddd.AggregationArgs{
			Id:        p.ID,
			State:     ddd.State(p.State),
			CreatedAt: &p.CreatedAt,
			UpdatedAt: &p.UpdatedAt,
		},
		Primary: p.Primary,
		Number:  p.Number,
		Country: p.Country,
	}
}

func (p Phone) TableName() string {
	return "school_phone"
}
