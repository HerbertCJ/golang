package helper

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TranslateError(err error, trans ut.Translator) []string {
	if err == nil {
		return nil
	}

	var errorMessages []string
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			errorMessages = append(errorMessages, e.Translate(trans))
		}
	}
	
	return errorMessages
}
