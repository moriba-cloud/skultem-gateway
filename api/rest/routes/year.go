package client

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-cloud/skultem-gateway/api/rest/middlewares"
	"github.com/moriba-cloud/skultem-gateway/api/rest/routes"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	"go.uber.org/zap"
	"time"
)

type (
	apiYear struct {
		validation *validation.Validation
		app        year.App
		logger     *zap.Logger
	}
	Year struct {
		Id        string    `json:"id"`
		Start     int64     `json:"start"`
		End       int64     `json:"end"`
		Year      string    `json:"year"`
		State     ddd.State `json:"state"`
		CreatedAt string    `json:"createdAt"`
		UpdatedAt string    `json:"updatedAt"`
	}
	YearRequest struct {
		Start int64 `json:"start" validate:"required"`
		End   int64 `json:"end" validate:"required"`
	}
)

func YearResponse(o *year.Domain) *Year {
	return &Year{
		Id:        o.ID(),
		Year:      fmt.Sprintf("%d - %d", o.Start(), o.End()),
		Start:     o.Start(),
		End:       o.End(),
		State:     o.State(),
		CreatedAt: o.CreatedAt().Format(time.RFC850),
		UpdatedAt: o.UpdatedAt().Format(time.RFC850),
	}
}

func (a apiYear) one(c *fiber.Ctx) error {
	payload := new(dto.ById)
	if err := c.ParamsParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.One(c.Context(), payload.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.NewResponse[Year](dto.ResponseArgs[Year]{
		Record: YearResponse(res.Record()),
	}))
}

func (a apiYear) listByPage(c *fiber.Ctx) error {
	payload := new(dto.Pagination)
	if err := c.QueryParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.ListByPage(c.Context(), ddd.PaginationArgs{
		Limit: int(payload.Limit),
		Page:  int(payload.Page),
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	records := make([]*Year, 0)
	for _, record := range res.Records() {
		records = append(records, YearResponse(record))
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Year]{
		Pagination: &dto.Pagination{
			Limit: uint32(res.Pagination.Limit()),
			Page:  uint64(res.Pagination.Page()),
			Pages: uint64(res.Pagination.Pages()),
			Size:  uint64(res.Pagination.Size()),
		},
		Records: records,
	}))
}

func (a apiYear) list(c *fiber.Ctx) error {
	res, err := a.app.List(c.Context())
	if err != nil {
		return err
	}

	records := make([]*routes.Option, 0)
	for _, record := range res.Records() {
		records = append(records, &routes.Option{
			Label: record.Label,
			Value: record.Value,
		})
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[routes.Option]{
		Records: records,
	}))
}

func (a apiYear) new(c *fiber.Ctx) error {
	payload := new(YearRequest)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.New(c.Context(), year.Args{
		Start: payload.Start,
		End:   payload.End,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Year]{
		Record: YearResponse(res.Record()),
	}))
}

func YearRoute(api fiber.Router, app year.App, logger *zap.Logger) {
	r := &apiYear{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	router := api.Group("/management/year", middlewares.AccessTokenGuard)
	router.Get("", r.listByPage)
	router.Get("/option", r.list)
	router.Get("/:id", r.one)
	router.Post("", r.new)
}
