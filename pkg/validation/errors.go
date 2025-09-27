package validation

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// FormatErrors converts validator.ValidationErrors into a map of JSON field -> human-readable message.
func FormatErrors(err error, model any) map[string]string {
	errors := make(map[string]string)

	if errs, ok := err.(validator.ValidationErrors); ok {
		t := reflect.TypeOf(model)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		for _, e := range errs {
			fieldName := e.Field()

			// Try to get JSON tag
			if f, ok := t.FieldByName(e.StructField()); ok {
				if tag := f.Tag.Get("json"); tag != "" && tag != "-" {
					fieldName = strings.Split(tag, ",")[0]
				}
			}

			switch e.Tag() {
			case "required":
				errors[fieldName] = fieldName + " is required"
			case "email":
				errors[fieldName] = "Invalid email format"
			case "min":
				errors[fieldName] = fieldName + " must be at least " + e.Param() + " characters"
			case "max":
				errors[fieldName] = fieldName + " must be at most " + e.Param() + " characters"
			case "unique":
				errors[fieldName] = fieldName + " must be unique"
			default:
				errors[fieldName] = fieldName + " is invalid"
			}
		}
	} else {
		errors["error"] = err.Error()
	}

	return errors
}
