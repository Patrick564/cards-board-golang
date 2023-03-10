package api

import (
	"encoding/json"
	"net/http"

	"github.com/Patrick564/cards-board-golang/models"
	"github.com/Patrick564/cards-board-golang/utils"
)

type UserEnv struct {
	Users interface {
		Add(user models.User) error
		Find(email, password string) (models.User, error)
	}
}

// ShowAccount godoc
//
//	@Summary		Register an account
//	@Description	post a new account
//	@Tags			accounts
//	@Accept			json
//	@Param			request	body	models.User	true	"query params"
//	@Success		200
//	@Router			/api/auth/register  [post]
func (env *UserEnv) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	hash, err := utils.HashAndSalt([]byte(user.Password))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}
	user.Password = hash

	err = env.Users.Add(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ShowAccount godoc
//
//	@Summary		Login with email and password
//	@Description	login
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.User	true	"query params"
//	@Success		200		{object}	models.User
//	@Router			/api/auth/login  [post]
func (env *UserEnv) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	query, err := env.Users.Find(user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	res, err := json.Marshal(query)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
