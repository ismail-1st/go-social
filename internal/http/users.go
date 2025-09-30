package http

import (
	"encoding/json"
	"net/http"
	"social/internal/models"
	"social/internal/services"
	"social/pkg/validation"

	"github.com/go-chi/chi/v5"
)

// type UserStore interface {
// 	Create(context.Context, *models.User) error
// 	// Register(context.Context, *models.User) error
// }

type UserHandler struct {
	// Store UserStore
	Service *services.UserService
}

func (h UserHandler) UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ListUsers)
	r.Post("/", h.CreateUser)
	r.Get("/{id}", h.ListUsers)
	r.Delete("/{id}", h.DeleteUser)
	r.Patch("/{id}", h.UpdateUser)

	return r
}

func (h UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// first, we decode the request body into a new user model
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if validation.ValidateAndRespond(w, user) {
		return
	}

	err = h.Service.Create(r.Context(), &user)
	if err != nil {
		validation.HandleDBError(w, err)
		return
	}

	// finally, we respond with a success message if all the checks are passed
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": user.ID})
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
