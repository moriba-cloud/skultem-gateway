package management

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	ose "github.com/moriba-build/ose/domain"
	"github.com/moriba-cloud/skultem-gateway/api/rest/routes/core"
	"github.com/moriba-cloud/skultem-gateway/domain/guardian"
	"go.uber.org/zap"
	"time"
)

type (
	rGuardian struct {
		validation *validation.Validation
		app        guardian.App
		logger     *zap.Logger
	}
	Guardian struct {
		Id         string       `json:"id"`
		GivenNames string       `json:"givenNames"`
		FamilyName string       `json:"familyName"`
		Profession string       `json:"profession"`
		Email      string       `json:"email"`
		Region     string       `json:"region"`
		Chiefdom   string       `json:"chiefdom"`
		District   string       `json:"district"`
		City       string       `json:"city"`
		Street     string       `json:"street"`
		Phones     []core.Phone `json:"phones"`
		State      ddd.State    `json:"state"`
		CreatedAt  string       `json:"createdAt"`
		UpdatedAt  string       `json:"updatedAt"`
	}
	GuardianReference struct {
		Id         string       `json:"id"`
		GivenNames string       `json:"givenNames"`
		FamilyName string       `json:"familyName"`
		Profession string       `json:"profession"`
		Email      string       `json:"email"`
		Region     string       `json:"region"`
		Chiefdom   string       `json:"chiefdom"`
		District   string       `json:"district"`
		City       string       `json:"city"`
		Street     string       `json:"street"`
		Phones     []core.Phone `json:"phones"`
	}
	GuardianRequest struct {
		GivenNames string              `json:"givenNames" validate:"required"`
		FamilyName string              `json:"familyName" validate:"required"`
		Profession string              `json:"profession" validate:"required"`
		Email      string              `json:"email"`
		Region     string              `json:"region" validate:"required"`
		Chiefdom   string              `json:"chiefdom" validate:"required"`
		District   string              `json:"district" validate:"required"`
		City       string              `json:"city" validate:"required"`
		Street     string              `json:"street" validate:"required"`
		Phones     []core.PhoneRequest `json:"phones" validate:"required,dive,required"`
	}
)

func GuardianResponse(o *guardian.Domain) *Guardian {
	phones := make([]core.Phone, 0)
	for _, phone := range o.Phones() {
		phones = append(phones, core.Phone{
			Id:        phone.ID(),
			Primary:   phone.Primary(),
			Number:    phone.Number(),
			State:     phone.State(),
			CreatedAt: phone.CreatedAt().Format(time.RFC850),
			UpdatedAt: phone.UpdatedAt().Format(time.RFC850),
		})
	}

	return &Guardian{
		Id:         o.ID(),
		GivenNames: o.GivenNames(),
		FamilyName: o.FamilyName(),
		Profession: o.Profession(),
		Email:      o.Email(),
		Region:     o.Region(),
		Chiefdom:   o.Chiefdom(),
		District:   o.District(),
		City:       o.City(),
		Street:     o.Street(),
		Phones:     phones,
		State:      o.State(),
		CreatedAt:  o.CreatedAt().Format(time.RFC850),
		UpdatedAt:  o.UpdatedAt().Format(time.RFC850),
	}
}

func (r rGuardian) one(c *fiber.Ctx) error {
	payload := new(dto.ById)
	if err := c.ParamsParser(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := r.validation.Run(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := r.app.One(c.Context(), payload.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Guardian]{
		Record: GuardianResponse(res.Record()),
	}))
}

func (r rGuardian) list(c *fiber.Ctx) error {
	res, err := r.app.List(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	records := make([]*core.Option, 0)
	for _, record := range res.Records() {
		records = append(records, &core.Option{
			Label: record.Label,
			Value: record.Value,
		})
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[core.Option]{
		Pagination: &dto.Pagination{
			Limit: uint32(res.Pagination.Limit()),
			Page:  uint64(res.Pagination.Page()),
			Pages: uint64(res.Pagination.Pages()),
			Size:  uint64(res.Pagination.Size()),
		},
		Records: records,
	}))
}

func (r rGuardian) listByPage(c *fiber.Ctx) error {
	payload := new(dto.Pagination)
	if err := c.QueryParser(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := r.validation.Run(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := r.app.ListByPage(c.Context(), ddd.PaginationArgs{
		Limit: int(payload.Limit),
		Page:  int(payload.Page),
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	records := make([]*Guardian, 0)
	for _, record := range res.Records() {
		records = append(records, GuardianResponse(record))
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Guardian]{
		Pagination: &dto.Pagination{
			Limit: uint32(res.Pagination.Limit()),
			Page:  uint64(res.Pagination.Page()),
			Pages: uint64(res.Pagination.Pages()),
			Size:  uint64(res.Pagination.Size()),
		},
		Records: records,
	}))
}

func (r rGuardian) new(c *fiber.Ctx) error {
	payload := new(GuardianRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := r.validation.Run(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	phones := make([]ose.PhoneArgs, 0)
	for _, phone := range payload.Phones {
		phones = append(phones, ose.PhoneArgs{
			Number: phone.Number,
		})
	}

	res, err := r.app.New(c.Context(), guardian.Args{
		GivenNames: payload.GivenNames,
		FamilyName: payload.FamilyName,
		Profession: payload.Profession,
		Phones:     phones,
		Email:      payload.Email,
		Region:     payload.Region,
		Chiefdom:   payload.Chiefdom,
		District:   payload.District,
		City:       payload.City,
		Street:     payload.Street,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Guardian]{
		Record: GuardianResponse(res.Record()),
	}))
}

func GuardianRoute(api fiber.Router, app guardian.App, logger *zap.Logger) {
	r := &rGuardian{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	router := api.Group("/guardian")
	router.Get("/option", r.list)
	router.Get("", r.listByPage)
	router.Get("/:id", r.one)
	router.Post("", r.new)
}
