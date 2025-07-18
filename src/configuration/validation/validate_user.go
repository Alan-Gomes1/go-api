package validation

import (
	"encoding/json"
	"errors"

	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translator "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translator.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(err error) *rest_err.Errors {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(err, &jsonValidationErr) {
		details := []rest_err.Details{}
		for _, fieldError := range err.(validator.ValidationErrors) {
			detail := rest_err.Details{
				Field:   fieldError.Field(),
				Message: fieldError.Translate(transl),
			}
			details = append(details, detail)
		}
		return rest_err.NewValidationError("Some fields are invalid", details)
	}
	return rest_err.NewBadRequestError("Error trying to convert fields")
}
