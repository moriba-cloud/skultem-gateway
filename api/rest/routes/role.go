package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-cloud/skultem-gateway/domain/role"
	"go.uber.org/zap"
	"time"
)

type (
	apiRole struct {
		validation *validation.Validation
		app        role.App
		logger     *zap.Logger
	}
	Role struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		State       ddd.State `json:"state"`
		CreatedAt   string    `json:"createdAt"`
		UpdatedAt   string    `json:"updatedAt"`
	}
	RoleRequest struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
)

func RoleResponse(o *role.Domain) *Role {
	return &Role{
		Id:          o.ID(),
		Name:        o.Name(),
		Description: o.Description(),
		State:       o.State(),
		CreatedAt:   o.CreatedAt().Format(time.RFC850),
		UpdatedAt:   o.UpdatedAt().Format(time.RFC850),
	}
}

func (a apiRole) listByPage(c *fiber.Ctx) error {
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

	records := make([]*Role, 0)
	for _, record := range res.Records() {
		records = append(records, RoleResponse(record))
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Role]{
		Pagination: &dto.Pagination{
			Limit: uint32(res.Pagination.Limit()),
			Page:  uint64(res.Pagination.Page()),
			Pages: uint64(res.Pagination.Pages()),
			Size:  uint64(res.Pagination.Size()),
		},
		Records: records,
	}))
}

func (a apiRole) list(c *fiber.Ctx) error {
	res, err := a.app.List(c.Context())
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

func (a apiRole) new(c *fiber.Ctx) error {
	payload := new(FeatureRequest)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.New(c.Context(), role.Args{
		Name:        payload.Name,
		Description: payload.Description,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Role]{
		Record: RoleResponse(res.Record()),
	}))
}

func (a apiRole) update(c *fiber.Ctx) error {
	byId := new(dto.ById)
	if err := c.ParamsParser(byId); err != nil {
		return err
	}
	if err := a.validation.Run(byId); err != nil {
		return err
	}

	payload := new(RoleRequest)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.Update(c.Context(), role.Args{
		Aggregation: ddd.AggregationArgs{
			Id: byId.Id,
		},
		Name:        payload.Name,
		Description: payload.Description,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Role]{
		Record: RoleResponse(res.Record()),
	}))
}

func (a apiRole) remove(c *fiber.Ctx) error {
	payload := new(dto.ById)
	if err := c.ParamsParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.Remove(c.Context(), payload.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Role]{
		Record: RoleResponse(res.Record()),
	}))
}

func RoleRoute(api fiber.Router, app role.App, logger *zap.Logger) {
	r := &apiRole{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	router := api.Group("/role")
	router.Get("", r.listByPage)
	router.Get("/option", r.list)
	router.Post("", r.new)
	router.Patch("/:id", r.update)
	router.Delete("/:id", r.remove)
}
