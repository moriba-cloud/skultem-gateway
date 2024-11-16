package authorization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-cloud/skultem-gateway/api/rest/middlewares"
	core2 "github.com/moriba-cloud/skultem-gateway/api/rest/routes"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"go.uber.org/zap"
	"time"
)

type (
	apiUser struct {
		validation *validation.Validation
		app        user.App
		logger     *zap.Logger
	}
	User struct {
		Id          string          `json:"id"`
		GivenNames  string          `json:"givenNames"`
		FamilyName  string          `json:"familyName"`
		Email       string          `json:"email"`
		Phone       int             `json:"phone"`
		Role        core2.Reference `json:"role"`
		PasswordTxt string          `json:"password"`
		State       ddd.State       `json:"state"`
		CreatedAt   string          `json:"createdAt"`
		UpdatedAt   string          `json:"updatedAt"`
	}
	UserRequest struct {
		GivenNames string `json:"givenNames" validate:"required"`
		FamilyName string `json:"familyName" validate:"required"`
		Phone      int    `json:"phone" validate:"required"`
		Email      string `json:"email" validate:"email"`
		Role       string `json:"role" validate:"required"`
	}
)

func UserResponse(o *user.Domain) *User {
	return &User{
		Id:         o.ID(),
		GivenNames: o.GivenNames(),
		FamilyName: o.FamilyName(),
		Email:      o.Email(),
		Role: core2.Reference{
			Id:    o.Role().Id,
			Value: o.Role().Value,
		},
		PasswordTxt: o.Password().Value,
		Phone:       o.Phone(),
		State:       o.State(),
		CreatedAt:   o.CreatedAt().Format(time.RFC850),
		UpdatedAt:   o.UpdatedAt().Format(time.RFC850),
	}
}

func (a apiUser) listByPage(c *fiber.Ctx) error {
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

	records := make([]*User, 0)
	for _, record := range res.Records() {
		records = append(records, UserResponse(record))
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[User]{
		Pagination: &dto.Pagination{
			Limit: uint32(res.Pagination.Limit()),
			Page:  uint64(res.Pagination.Page()),
			Pages: uint64(res.Pagination.Pages()),
			Size:  uint64(res.Pagination.Size()),
		},
		Records: records,
	}))
}

func (a apiUser) list(c *fiber.Ctx) error {
	res, err := a.app.List(c.Context())
	if err != nil {
		return err
	}

	records := make([]*core2.Option, 0)
	for _, record := range res.Records() {
		records = append(records, &core2.Option{
			Label: record.Label,
			Value: record.Value,
		})
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[core2.Option]{
		Records: records,
	}))
}

func (a apiUser) new(c *fiber.Ctx) error {
	payload := new(UserRequest)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.New(c.Context(), user.Args{
		GivenNames: payload.GivenNames,
		FamilyName: payload.FamilyName,
		Phone:      payload.Phone,
		Email:      payload.Email,
		Role: core.Reference{
			Id: payload.Role,
		},
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[User]{
		Record: UserResponse(res.Record()),
	}))
}

func (a apiUser) update(c *fiber.Ctx) error {
	byId := new(dto.ById)
	if err := c.ParamsParser(byId); err != nil {
		return err
	}
	if err := a.validation.Run(byId); err != nil {
		return err
	}

	payload := new(UserRequest)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.Update(c.Context(), user.Args{
		Aggregation: ddd.AggregationArgs{
			Id: byId.Id,
		},
		GivenNames: payload.GivenNames,
		FamilyName: payload.FamilyName,
		Phone:      payload.Phone,
		Email:      payload.Email,
		Role: core.Reference{
			Id: payload.Role,
		},
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[User]{
		Record: UserResponse(res.Record()),
	}))
}

func (a apiUser) findById(c *fiber.Ctx) error {
	payload := new(dto.ById)
	if err := c.ParamsParser(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := a.validation.Run(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := a.app.FindById(c.Context(), payload.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[User]{
		Record: UserResponse(res.Record()),
	}))
}

func (a apiUser) remove(c *fiber.Ctx) error {
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

	return c.JSON(dto.NewResponse(dto.ResponseArgs[User]{
		Record: UserResponse(res.Record()),
	}))
}

func UserRoute(api fiber.Router, app user.App, logger *zap.Logger) {
	r := &apiUser{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	api.Group("/user", middlewares.AccessTokenGuard).
		Get("", r.listByPage).
		Get("/option", r.list).
		Get("/:id", r.findById).
		Post("", r.new).
		Patch("/:id", r.update).
		Delete("/:id", r.remove)
}
