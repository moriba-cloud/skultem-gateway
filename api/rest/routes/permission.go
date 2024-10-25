package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-cloud/skultem-gateway/domain/permission"
	"go.uber.org/zap"
	"time"
)

type (
	apiPermission struct {
		validation *validation.Validation
		app        permission.App
		logger     *zap.Logger
	}
	Permission struct {
		Id        string    `json:"id"`
		Feature   string    `json:"feature"`
		Create    bool      `json:"create"`
		ReadAll   bool      `json:"readAll"`
		Read      bool      `json:"read"`
		Edit      bool      `json:"edit"`
		Delete    bool      `json:"delete"`
		State     ddd.State `json:"state"`
		CreatedAt string    `json:"createdAt"`
		UpdatedAt string    `json:"updatedAt"`
	}
	One struct {
		Feature string `json:"feature" validate:"required"`
		Create  bool   `json:"create" validate:"required"`
		ReadAll bool   `json:"readAll" validate:"required"`
		Read    bool   `json:"read" validate:"required"`
		Edit    bool   `json:"edit" validate:"required"`
		Delete  bool   `json:"delete" validate:"required"`
	}
	PermissionRequest struct {
		Permissions []One `json:"permissions" validate:"required,dive,required"`
	}
)

func PermissionResponse(o *permission.Domain) *Permission {
	return &Permission{
		Id:        o.ID(),
		Feature:   o.Feature(),
		Create:    o.Create(),
		Read:      o.Read(),
		ReadAll:   o.ReadAll(),
		Edit:      o.Edit(),
		Delete:    o.Delete(),
		State:     o.State(),
		CreatedAt: o.CreatedAt().Format(time.RFC850),
		UpdatedAt: o.UpdatedAt().Format(time.RFC850),
	}
}

func (a apiPermission) new(c *fiber.Ctx) error {
	payload := new(PermissionRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := a.validation.Run(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	byId := new(dto.ById)
	if err := c.ParamsParser(byId); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := a.validation.Run(byId); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	permissions := make([]*permission.Args, len(payload.Permissions))
	for i, one := range payload.Permissions {
		permissions[i] = &permission.Args{
			Feature: one.Feature,
			Create:  one.Create,
			Read:    one.Read,
			ReadAll: one.ReadAll,
			Edit:    one.Edit,
			Delete:  one.Delete,
		}
	}

	res, err := a.app.Update(c.Context(), permissions, byId.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	records := make([]*Permission, 0)
	for _, record := range res.Records() {
		records = append(records, PermissionResponse(record))
	}
	return c.JSON(dto.NewResponse(dto.ResponseArgs[Permission]{
		Records: records,
	}))
}

//
//func (a apiRole) update(c *fiber.Ctx) error {
//	byId := new(dto.ById)
//	if err := c.ParamsParser(byId); err != nil {
//		return err
//	}
//	if err := a.validation.Run(byId); err != nil {
//		return err
//	}
//
//	payload := new(RoleRequest)
//	if err := c.BodyParser(payload); err != nil {
//		return err
//	}
//	if err := a.validation.Run(payload); err != nil {
//		return err
//	}
//
//	res, err := a.app.Update(c.Context(), role.Args{
//		Aggregation: ddd.AggregationArgs{
//			Id: byId.Id,
//		},
//		Name:        payload.Name,
//		Description: payload.Description,
//	})
//	if err != nil {
//		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
//	}
//
//	return c.JSON(dto.NewResponse(dto.ResponseArgs[Role]{
//		Record: RoleResponse(res.Record()),
//	}))
//}
//
//func (a apiRole) remove(c *fiber.Ctx) error {
//	payload := new(dto.ById)
//	if err := c.ParamsParser(payload); err != nil {
//		return err
//	}
//	if err := a.validation.Run(payload); err != nil {
//		return err
//	}
//
//	res, err := a.app.Remove(c.Context(), payload.Id)
//	if err != nil {
//		return fiber.NewError(fiber.StatusNotFound, err.Error())
//	}
//
//	return c.JSON(dto.NewResponse(dto.ResponseArgs[Role]{
//		Record: RoleResponse(res.Record()),
//	}))
//}

func PermissionRoute(api fiber.Router, app permission.App, logger *zap.Logger) {
	r := &apiPermission{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	router := api.Group("/permission")
	router.Post("/:id", r.new)
}
