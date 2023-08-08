package authcontrollers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/deanrtaylor1/backend-go/internal/controllers/basecontrollers"
	db "github.com/deanrtaylor1/backend-go/internal/db/sqlc"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type createUserRequest struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func NewAuthController(baseController basecontrollers.BaseController) basecontrollers.Controller {
	return &AuthController{
		BaseController: baseController,
	}
}

func (a *AuthController) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/signup", a.SignUp)
	r.Post("/login", a.Login)

	return r
}

type SignUpParams struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

func (a *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var params createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// err, res := util.ValidateStruct(a.Validator, params, a.Config)
	// if err != nil {
	// 	util.SendResponse(w, res.Code, res)
	// }

	createUser := db.CreateUserParams{
		Email:          params.Email,
		Username:       params.Username,
		HashedPassword: params.Password,
	}
	user, err := a.Store.CreateUser(context.Background(), createUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// Implement login logic
}
