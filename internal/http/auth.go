package http

import (
	"net/http"
	"social/internal/dto"
	"social/internal/services"
	"social/pkg/json"
	"social/pkg/response"
	"social/pkg/validation"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	Service *services.AuthService
}

func (h AuthHandler) AuthRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/register", h.RegisterUser)
	r.Post("/login", h.LoginUser)

	return r
}

func (h AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// first, we decode the request body into a new user model
	user := dto.UserRegister{}
	if !json.ParseAndValidate(w, r, &user) {
		return
	}

	createdUser, err := h.Service.Register(r.Context(), &user)
	if err != nil {
		validation.HandleDBError(w, err)
		return
	}

	// finally, we respond with a success message if all the checks are passed
	response.Created(w, "User registered successfully", map[string]interface{}{
		"user": dto.UserRegisterResponse{
			ID:       createdUser.ID,
			Username: createdUser.Username,
			Email:    createdUser.Email,
		},
	})
}

func (h AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	user := dto.UserLogin{}
	if !json.ParseAndValidate(w, r, &user) {
		return
	}

	loggedInUser, err := h.Service.Login(r.Context(), &user)

	if err != nil {
		if err.Error() == "invalid credentials" {
			response.Error(w, http.StatusUnauthorized, "Invalid credentials", map[string]string{
				"db": "invalid credentials",
			})
			return
		}
		validation.HandleDBError(w, err)
		return
	}

	response.OK(w, "User logged in successfully", map[string]interface{}{
		"user": dto.UserLoginResponse{
			ID:           loggedInUser.ID,
			Email:        loggedInUser.Email,
			AccessToken:  loggedInUser.AccessToken,
			RefreshToken: loggedInUser.RefreshToken,
		},
	})
}

func (h UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	user := dto.UserLoginDetail{}

	if !json.ParseAndValidate(w, r, &user) {
		return
	}

	// gottenUser, err := h.Service.GetUserByEmail()
}

// func (h AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
// 	user := dto.ForgetPassword{}
// 	if !json.ParseAndValidate(w, r, &user) {
// 		return
// 	}

// 	err, gottenUser := h.Service.GetUserByEmail(user.Email)

// }
