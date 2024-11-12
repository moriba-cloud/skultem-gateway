package validations

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/moriba-build/ose/ddd/utils/stn"
	"github.com/moriba-cloud/skultem-management/domain/values"
	"slices"
)

func BatchValidation(validate *validator.Validate) {
	err := validate.RegisterValidation("batch", func(fl validator.FieldLevel) bool {
		status := make([]string, 0)
		status = append(status, stn.Key(string(values.DESIGNATION)))
		status = append(status, stn.Key(string(values.SUBJECT)))
		status = append(status, stn.Key(string(values.SECTION)))
		status = append(status, stn.Key(string(values.RELIGION)))
		status = append(status, stn.Key(string(values.PAYMENT_PLAN)))

		return slices.Contains(status, stn.Key(fl.Field().String()))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func BatchTranslation(validate *validator.Validate, translator ut.Translator) {
	err := validate.RegisterTranslation("batch", translator, func(ut ut.Translator) error {
		return ut.Add("batch", "{0} must be either PAYMENT_PLAN | DESIGNATION | SUBJECT | SECTION", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("batch", fmt.Sprintf("%s '%s'", fe.Field(), fe.Value().(string)))
		return t
	})
	if err != nil {
		log.Fatal(err)
	}
}
