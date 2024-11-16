package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/leekchan/accounting"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/config"
	routes2 "github.com/moriba-cloud/skultem-gateway/api/rest/routes/management"
	"github.com/moriba-cloud/skultem-gateway/app"
	"go.uber.org/zap"
)

type (
	PhoneRequest struct {
		Number int `json:"number" validate:"required,phone"`
	}
	ExperienceRequest struct {
		Company  string `json:"company" validate:"required"`
		Position string `json:"position" validate:"required"`
		Start    string `json:"start" validate:"required,date"`
		End      string `json:"end"`
	}
	EducationRequest struct {
		Qualification string `json:"qualification" validate:"required"`
		School        string `json:"config" validate:"required"`
	}
	Phone struct {
		Id        string    `json:"id"`
		Primary   bool      `json:"primary"`
		Number    int       `json:"number" validate:"required,phone"`
		State     ddd.State `json:"state"`
		CreatedAt string    `json:"createdAt"`
		UpdatedAt string    `json:"updatedAt"`
	}
	Option struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}
	Reference struct {
		Id    string `json:"id"`
		Value string `json:"value"`
	}
	Education struct {
		Id            string    `json:"id"`
		Qualification string    `json:"qualification" validate:"required"`
		School        string    `json:"config" validate:"required"`
		State         ddd.State `json:"state"`
		CreatedAt     string    `json:"createdAt"`
		UpdatedAt     string    `json:"updatedAt"`
	}
	Experience struct {
		Id        string    `json:"id"`
		Company   string    `json:"company" validate:"required"`
		Position  string    `json:"position" validate:"required"`
		Start     string    `json:"start" validate:"required,date"`
		End       string    `json:"end" validate:"date"`
		State     ddd.State `json:"state"`
		CreatedAt string    `json:"createdAt"`
		UpdatedAt string    `json:"updatedAt"`
	}
	Args struct {
		Route  fiber.Router
		Apps   *app.Apps
		Logger *zap.Logger
	}
)

func Currency(amount float64) string {
	ac := accounting.DefaultAccounting("NLE ", 2)
	return ac.FormatMoney(amount)
}

func Routes(args Args) {
	version := config.NewEnvs().EnvStr("API_VERSION")
	apiVersion := fmt.Sprintf("/api/%s", version)
	route := args.Route.Group(apiVersion)

	// routes
	YearRoute(route, args.Apps.Year, args.Logger)
	routes2.ValuesRoute(route, args.Apps.Value, args.Logger)
	FeatureRoute(route, args.Apps.Feature, args.Logger)
	RoleRoute(route, args.Apps.Role, args.Logger)
	PermissionRoute(route, args.Apps.Permission, args.Logger)
	UserRoute(route, args.Apps.User, args.Logger)
	AuthRoute(route, args.Apps.Auth, args.Logger)
	SchoolRoute(route, args.Apps.School, args.Logger)
}
