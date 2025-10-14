package json

import (
	"net/http"
	"social/pkg/validation"
)

// ParseAndValidate decodes the JSON body strictly into v and validates it.
// returns false if decoding or validation fails (and writes response errors automatically).
func ParseAndValidate(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if ok := DecodeStrict(w, r, v); !ok {
		return false
	}

	if validation.ValidateAndRespond(w, v) {
		return false
	}

	return true
}
