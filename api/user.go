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

type Env struct {
	Users interface {
		Add(user models.User) error
		Find(email, password string) (models.User, error)
	}
}

func (e *Env) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading route /api/register")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(utils.CustomError{Message: "method not allowed"})
		return
	}

	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	hash, _ := utils.HashAndSalt([]byte(u.Password))
	u.Password = hash

	e.Users.Add(u)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (e *Env) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading route /api/login")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(utils.CustomError{Message: "method not allowed"})
		return
	}

	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	q, err := e.Users.Find(u.Email, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	res, err := json.Marshal(q)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
