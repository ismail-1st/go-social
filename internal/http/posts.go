package http

import (
	"context"
	"net/http"
	"social/internal/models"

	"github.com/go-chi/chi/v5"
)

type PostStore interface {
	Create(context.Context, *models.Post) error
}

type PostHandler struct {
	Store PostStore
}

func (h PostHandler) PostRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/admin/all", h.ListPosts)
	r.Post("/", h.CreatePost)
	r.Get("/", h.GetPosts)
	r.Get("/{id}", h.GetPostById)
	r.Patch("/{id}", h.UpdatePost)
	r.Delete("/{id}", h.DeletePost)

	return r
}

func (h PostHandler) ListPosts(w http.ResponseWriter, r *http.Request) {

}

func (h PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

}

func (h PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {

}

func (h PostHandler) GetPostById(w http.ResponseWriter, r *http.Request) {

}

func (h PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func (h PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

}
