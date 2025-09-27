package http

import (
	"context"
	"encoding/json"
	"net/http"
	"social/internal/models"
	"social/pkg/validation"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	Create(context.Context, *models.User) error
}

type UserHandler struct {
	Store UserStore
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

	// second, we hash the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
	}
	user.Password = string(hashedPassword)

	// third, call the store's create method
	err = h.Store.Create(r.Context(), &user)
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
