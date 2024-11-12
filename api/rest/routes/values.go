package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-build/ose/ddd/utils/stn"
	"github.com/moriba-cloud/skultem-management/api/rest/validations"
	"github.com/moriba-cloud/skultem-management/domain/values"
	"go.uber.org/zap"
	"time"
)

type (
	rValues struct {
		validation *validation.Validation
		app        values.App
		logger     *zap.Logger
	}
	Values struct {
		Id        string    `json:"id"`
		Key       string    `json:"key"`
		Value     string    `json:"value"`
		Batch     string    `json:"batch"`
		State     ddd.State `json:"state"`
		CreatedAt string    `json:"createdAt"`
		UpdatedAt string    `json:"updatedAt"`
	}
	ValuesRequest struct {
		Key   string `json:"key" validate:"required,min=3"`
		Value string `json:"value" validate:"required,min=3"`
		Batch string `json:"batch" validate:"required,batch"`
	}
)

func ValuesResponse(o *values.Domain) *Values {
	return &Values{
		Id:        o.ID(),
		Key:       o.Key(),
		Value:     o.Value(),
		Batch:     string(o.Batch()),
		State:     o.State(),
		CreatedAt: o.CreatedAt().Format(time.RFC850),
		UpdatedAt: o.UpdatedAt().Format(time.RFC850),
	}
}

func (r *rValues) one(c *fiber.Ctx) error {
	payload := new(dto.ById)
	if err := c.ParamsParser(payload); err != nil {
		return err
	}
	if err := r.validation.Run(payload); err != nil {
		return err
	}

	res, err := r.app.One(c.Context(), payload.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	log.Infof("values with id: %s fetched", res.Record().ID())
	return c.JSON(dto.NewResponse[Values](dto.ResponseArgs[Values]{
		Record: ValuesResponse(res.Record()),
	}))
}

func (r *rValues) listByPage(c *fiber.Ctx) error {
	payload := new(dto.Pagination)
	if err := c.QueryParser(payload); err != nil {
		return err
	}
	if err := r.validation.Run(payload); err != nil {
		return err
	}

	res, err := r.app.ListByPage(c.Context(), ddd.PaginationArgs{
		Limit: int(payload.Limit),
		Page:  int(payload.Page),
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	records := make([]*Values, 0)
	for _, record := range res.Records() {
		records = append(records, ValuesResponse(record))
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Values]{
		Pagination: &dto.Pagination{
			Limit: uint32(res.Pagination.Limit()),
			Page:  uint64(res.Pagination.Page()),
			Pages: uint64(res.Pagination.Pages()),
			Size:  uint64(res.Pagination.Size()),
		},
		Records: records,
	}))
}

func (r *rValues) listByGroup(c *fiber.Ctx) error {
	value := c.Params("value")
	if len(value) <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "value is empty")
	}

	res, err := r.app.ListByGroup(c.Context(), values.Batch(value))
	if err != nil {
		return err
	}

	records := make([]*Option, 0)
	for _, record := range res.Records() {
		records = append(records, &Option{
			Label: record.Label,
			Value: record.Value,
		})
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Option]{
		Records: records,
	}))
}

func (r *rValues) list(c *fiber.Ctx) error {
	res, err := r.app.List(c.Context())
	if err != nil {
		return err
	}

	records := make([]*Option, 0)
	for _, record := range res.Records() {
		records = append(records, &Option{
			Label: record.Label,
			Value: record.Value,
		})
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Option]{
		Records: records,
	}))
}

func (r *rValues) new(c *fiber.Ctx) error {
	payload := new(ValuesRequest)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := r.validation.Run(payload); err != nil {
		return err
	}

	res, err := r.app.New(c.Context(), values.Args{
		Batch: values.Batch(stn.Key(payload.Batch)),
		Key:   payload.Key,
		Value: payload.Value,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Values]{
		Record: ValuesResponse(res.Record()),
	}))
}

func ValuesRoute(api fiber.Router, app values.App, logger *zap.Logger) {
	v := validation.NewValidation()
	validations.BatchValidation(v.Validate)
	validations.BatchTranslation(v.Validate, v.Translator)

	r := &rValues{
		app:        app,
		validation: v,
		logger:     logger,
	}

	router := api.Group("/values")
	router.Get("", r.listByPage)
	router.Get("/option", r.list)
	router.Get("/batch/:value", r.listByGroup)
	router.Get("/:id", r.one)
	router.Post("", r.new)
}
