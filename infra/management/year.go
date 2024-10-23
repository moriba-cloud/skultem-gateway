package management

import (
	"context"
	"fmt"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
)

type (
	yearSeat struct{}
)

func (y yearSeat) TakeOff(ctx context.Context, args year.Domain) {
	fmt.Print("Event publish")
}

func NewYearSeat() year.Bus {
	return yearSeat{}
}
