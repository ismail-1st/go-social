package http

import (
	"net/http"
	"social/internal/services"

	"github.com/go-chi/chi/v5"
)

// type UserStore interface {
// 	Create(context.Context, *models.User) error
// 	// Register(context.Context, *models.User) error
// }

type UserHandler struct {
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
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
