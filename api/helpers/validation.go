package helpers

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorMessages(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "string":
		return "This field must be a string"
	}
	return ""
}

func ConfigureValidationJsonFields(v *validator.Validate) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
