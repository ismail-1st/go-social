package http

import (
	"encoding/json"
	"net/http"
	"social/internal/dto"
	"social/internal/services"
	"social/pkg/validation"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	Service *services.UserService
}

func (h AuthHandler) UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/register", h.RegisterUser)

	return r
}

func (h AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// first, we decode the request body into a new user model
	user := dto.UserRegister{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if validation.ValidateAndRespond(w, user) {
		return
	}

	createdUser, err := h.Service.Create(r.Context(), &user)
	if err != nil {
		validation.HandleDBError(w, err)
		return
	}

	// finally, we respond with a success message if all the checks are passed
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": createdUser.ID})
}
