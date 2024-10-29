package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-build/ose/domain"
	"github.com/moriba-cloud/skultem-gateway/domain/school"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"go.uber.org/zap"
	"time"
)

type (
	apiSchool struct {
		validation *validation.Validation
		app        school.App
		logger     *zap.Logger
	}
	Owner struct {
		Id         string `json:"id"`
		GivenNames string `json:"givenNames"`
		FamilyName string `json:"familyName"`
		Email      string `json:"email"`
		Phone      int    `json:"phone"`
	}
	School struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		Domain    string    `json:"domain"`
		Email     string    `json:"email"`
		Region    string    `json:"region"`
		District  string    `json:"district"`
		Chiefdom  string    `json:"chiefdom"`
		City      string    `json:"city"`
		Street    string    `json:"street"`
		Phones    []int     `json:"phones"`
		Owner     Owner     `json:"owner"`
		State     ddd.State `json:"state"`
		CreatedAt string    `json:"createdAt"`
		UpdatedAt string    `json:"updatedAt"`
	}
	OwnerRequest struct {
		GivenNames string `json:"givenNames" validate:"required"`
		FamilyName string `json:"familyName" validate:"required"`
		Phone      int    `json:"phone" validate:"required"`
		Email      string `json:"email" validate:"email"`
	}
	SchoolRequest struct {
		Name     string `json:"name" validate:"required"`
		Domain   string `json:"domain" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Region   string `json:"region" validate:"required"`
		District string `json:"district" validate:"required"`
		Chiefdom string `json:"chiefdom" validate:"required"`
		City     string `json:"city" validate:"required"`
		Street   string `json:"street" validate:"required"`
		Phones   []int  `json:"phones" validate:"required"`
		Owner    Owner  `json:"owner" validate:"required"`
	}
)

func SchoolResponse(o *school.Domain) *School {
	phones := make([]int, len(o.Phones()))
	for i, phone := range o.Phones() {
		phones[i] = phone.Number()
	}

	return &School{
		Id:       o.ID(),
		Name:     o.Name(),
		Domain:   o.Domain(),
		Region:   o.Region(),
		District: o.District(),
		Chiefdom: o.Chiefdom(),
		City:     o.City(),
		Phones:   phones,
		Email:    o.Email(),
		Owner: Owner{
			Id:         o.ID(),
			GivenNames: o.Owner().GivenNames(),
			FamilyName: o.Owner().FamilyName(),
			Email:      o.Owner().Email(),
			Phone:      o.Owner().Phone(),
		},
		State:     o.State(),
		CreatedAt: o.CreatedAt().Format(time.RFC850),
		UpdatedAt: o.UpdatedAt().Format(time.RFC850),
	}
}

func (a apiSchool) listByPage(c *fiber.Ctx) error {
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

	records := make([]*School, 0)
	for _, record := range res.Records() {
		records = append(records, SchoolResponse(record))
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[School]{
		Pagination: &dto.Pagination{
			Limit: uint32(res.Pagination.Limit()),
			Page:  uint64(res.Pagination.Page()),
			Pages: uint64(res.Pagination.Pages()),
			Size:  uint64(res.Pagination.Size()),
		},
		Records: records,
	}))
}

func (a apiSchool) list(c *fiber.Ctx) error {
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

func (a apiSchool) new(c *fiber.Ctx) error {
	payload := new(SchoolRequest)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	phones := make([]domain.PhoneArgs, len(payload.Phones))
	for i, phone := range payload.Phones {
		phones[i] = domain.PhoneArgs{
			Number: phone,
		}
	}
	owner := user.Args{
		GivenNames: payload.Owner.GivenNames,
		FamilyName: payload.Owner.FamilyName,
		Phone:      payload.Owner.Phone,
		Email:      payload.Owner.Email,
	}

	res, err := a.app.New(c.Context(), school.Args{
		Name:     payload.Name,
		Domain:   payload.Domain,
		Email:    payload.Email,
		Region:   payload.Region,
		Chiefdom: payload.Chiefdom,
		District: payload.District,
		City:     payload.City,
		Street:   payload.Street,
		Phones:   phones,
		Owner:    owner,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	return c.JSON(dto.NewResponse(dto.ResponseArgs[School]{
		Record: SchoolResponse(res.Record()),
	}))
}

func (a apiSchool) update(c *fiber.Ctx) error {
	//byId := new(dto.ById)
	//if err := c.ParamsParser(byId); err != nil {
	//	return err
	//}
	//if err := a.validation.Run(byId); err != nil {
	//	return err
	//}
	//
	//payload := new(UserRequest)
	//if err := c.BodyParser(payload); err != nil {
	//	return err
	//}
	//if err := a.validation.Run(payload); err != nil {
	//	return err
	//}
	//
	//phones := make([]domain.PhoneArgs, len(payload.Phones))
	//for i, p := range payload.Phones {
	//	phone, err := strconv.Atoi(p)
	//	if err != nil {
	//		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	//	}
	//	phones[i] = domain.PhoneArgs{
	//		Number: phone,
	//	}
	//}
	//owner := user.Args{
	//	GivenNames: payload.Owner.GivenNames,
	//	FamilyName: payload.Owner.FamilyName,
	//	Phone:      payload.Owner.Phone,
	//	Email:      payload.Owner.Email,
	//}
	//
	//res, err := a.app.New(c.Context(), school.Args{
	//	Name:     payload.Name,
	//	Domain:   payload.Domain,
	//	Email:    payload.Email,
	//	Region:   payload.Region,
	//	Chiefdom: payload.Chiefdom,
	//	District: payload.District,
	//	City:     payload.City,
	//	Street:   payload.Street,
	//	Phones:   phones,
	//	Owner:    owner,
	//})
	//
	//if err != nil {
	//	return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	//}
	//
	//return c.JSON(dto.NewResponse(dto.ResponseArgs[School]{
	//	Record: SchoolResponse(res.Record()),
	//}))
	return nil
}

func (a apiSchool) remove(c *fiber.Ctx) error {
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

	return c.JSON(dto.NewResponse(dto.ResponseArgs[School]{
		Record: SchoolResponse(res.Record()),
	}))
}

func SchoolRoute(api fiber.Router, app school.App, logger *zap.Logger) {
	r := &apiSchool{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	api.Group("/school").
		Get("", r.listByPage).
		Get("/option", r.list).
		Post("", r.new).
		Patch("/:id", r.update).
		Delete("/:id", r.remove)
}
