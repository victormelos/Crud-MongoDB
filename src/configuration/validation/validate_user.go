package validation

import (
	"encoding/json"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en.New()
		unt := ut.New(en.New(), en.New())
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUser(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors
	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	}
