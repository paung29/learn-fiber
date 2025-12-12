package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate *validator.Validate
	Translator ut.Translator
)

func InitValidator() {
	Validate := validator.New()

	eng := en.New()
	uni := ut.New(eng, eng)

	Translator, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(Validate, Translator)
}

func TranslateValidationError(err error) []string {
	var errors []string

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			errors = append(errors, fe.Translate(Translator))
		}
	}

	return errors
}