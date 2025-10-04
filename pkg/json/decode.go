package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"social/pkg/response"
	"strings"
)

// decodes the JSON body strictly and sends a JSON error response on failure.
// returns true if decoding succeeded, false if it already handled the error response.
//
// example usage in an HTTP handler:
//
//   var user dto.UserLogin
//   if ok := json.DecodeStrict(w, r, &user); !ok {
//       return
//   }

func DecodeStrict(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil {
		var (
			msg    string
			errMap = make(map[string]string)
		)

		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError

		switch {
		// empty body
		case errors.Is(err, io.EOF):
			msg = "Request body is empty"

		// invalid JSON syntax
		case errors.As(err, &syntaxErr):
			msg = fmt.Sprintf("Invalid JSON syntax at byte %d", syntaxErr.Offset)

		// wrong type for a field (e.g. string instead of number)
		case errors.As(err, &unmarshalTypeErr):
			field := unmarshalTypeErr.Field
			msg = "Invalid type in JSON body"
			errMap[field] = fmt.Sprintf("Expected type %v but got invalid value", unmarshalTypeErr.Type)

		// unknown field (e.g. "username" not defined in struct)
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			re := regexp.MustCompile(`json: unknown field "(.+)"`)
			matches := re.FindStringSubmatch(err.Error())
			if len(matches) == 2 {
				field := matches[1]
				msg = "Invalid field in request body"
				errMap[field] = "This field is not allowed"
			} else {
				msg = "Invalid field in request body"
			}

		// generic fallback
		default:
			msg = "Invalid JSON request body"
		}

		response.Error(w, http.StatusBadRequest, msg, errMap)
		return false
	}

	return true
}
