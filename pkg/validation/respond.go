package validation

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lib/pq"
)

// WriteJSON sends a JSON response with given status.
func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// ErrorResponse standard shape for errors.
type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// SendValidationErrors sends 400 with formatted validation errors.
func SendValidationErrors(w http.ResponseWriter, errs map[string]string) {
	WriteJSON(w, http.StatusBadRequest, ErrorResponse{Errors: errs})
}

// SendDBError sends 400 or 500 depending on DB error.
func SendDBError(w http.ResponseWriter, field, message string) {
	WriteJSON(w, http.StatusBadRequest, ErrorResponse{
		Errors: map[string]string{
			field: message,
		},
	})
}

// SendServerError sends 500 with generic error message.
func SendServerError(w http.ResponseWriter, msg string) {
	WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": msg})
}

// HandleDBError inspects a DB error and sends a proper JSON response.
func HandleDBError(w http.ResponseWriter, err error) {
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "23505": // unique_violation
			// Extract column name from constraint
			field := parseConstraint(pqErr.Constraint)
			SendDBError(w, field, strings.Title(field)+" already exists")
			return

		case "23503": // foreign_key_violation
			field := parseConstraint(pqErr.Constraint)
			SendDBError(w, field, strings.Title(field)+" does not exist")
			return
		}
	}

	// fallback: unexpected DB error
	SendServerError(w, "Database error: "+err.Error())
}

// parseConstraint extracts field name from constraint string
// e.g. "users_email_key" -> "email"
func parseConstraint(constraint string) string {
	parts := strings.Split(constraint, "_")
	if len(parts) > 1 {
		return parts[1] // assumes users_email_key format
	}
	return constraint
}
