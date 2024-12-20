package authorization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-cloud/skultem-gateway/api/rest/middlewares"
	core2 "github.com/moriba-cloud/skultem-gateway/api/rest/routes"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/permission"
	"go.uber.org/zap"
	"time"
)

type (
	argsPermissions struct {
		validation *validation.Validation
		app        permission.App
		logger     *zap.Logger
	}
	Permission struct {
		Id        string          `json:"id"`
		Feature   core2.Reference `json:"feature"`
		Create    bool            `json:"create"`
		ReadAll   bool            `json:"readAll"`
		Read      bool            `json:"read"`
		Edit      bool            `json:"edit"`
		Delete    bool            `json:"delete"`
		State     ddd.State       `json:"state"`
		CreatedAt string          `json:"createdAt"`
		UpdatedAt string          `json:"updatedAt"`
	}
	One struct {
		Feature string `json:"feature" validate:"required"`
		Create  bool   `json:"create"`
		ReadAll bool   `json:"readAll"`
		Read    bool   `json:"read"`
		Edit    bool   `json:"edit"`
		Delete  bool   `json:"delete"`
	}
	PermissionRequest struct {
		Permissions []One `json:"permissions" validate:"required,dive,required"`
	}
)

func PermissionResponse(o *permission.Domain) *Permission {
	return &Permission{
		Id: o.ID(),
		Feature: core2.Reference{
			Id:    o.Feature().Id,
			Value: o.Feature().Value,
		},
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

func (a argsPermissions) new(c *fiber.Ctx) error {
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
			Feature: core.Reference{
				Id: one.Feature,
			},
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

func PermissionRoute(api fiber.Router, app permission.App, logger *zap.Logger) {
	r := &argsPermissions{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	router := api.Group("/permission", middlewares.AccessTokenGuard)
	router.Post("/:id", r.new)
}
