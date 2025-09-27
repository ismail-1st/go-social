package validation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

var Validate = validator.New()

// ValidateAndRespond validates a struct and writes JSON errors if invalid.
// Returns true if errors were written (so handler should return).
func ValidateAndRespond(w http.ResponseWriter, data any) bool {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println("Validating struct:", string(b))

	if err := Validate.Struct(data); err != nil {
		SendValidationErrors(w, FormatErrors(err, data))
		return true
	}
	return false
}
