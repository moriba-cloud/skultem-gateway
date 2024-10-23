package management

import (
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	"go.uber.org/zap"
)

type (
	Bus struct {
		Year year.Bus
	}
	Args struct {
		Logger *zap.Logger
	}
)

func NewBus() Bus {
	return Bus{
		Year: NewYearSeat(),
	}
}
