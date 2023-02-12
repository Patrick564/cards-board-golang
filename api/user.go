package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Patrick564/cards-board-golang/models"
	"github.com/Patrick564/cards-board-golang/utils"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Range string `json:"range"`
	jwt.RegisteredClaims
}

type UserEnv struct {
	Users interface {
		Add(user models.User) error
		Find(email, password string) (models.User, error)
	}
}

func (env *UserEnv) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading route /api/register")
	w.Header().Set("Content-Type", "application/json")

	// Only POST method allowed
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(utils.CustomError{Message: "method not allowed"})
		return
	}

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

func (env *UserEnv) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading route /api/login")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(utils.CustomError{Message: "method not allowed"})
		return
	}

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
