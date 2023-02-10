package main

import (
	"log"
	"net/http"

	"github.com/Patrick564/cards-board-golang/api"
	"github.com/Patrick564/cards-board-golang/models"
)

type Env struct {
	users interface {
		Register() (models.User, error)
	}
}

func main() {
	mux := http.NewServeMux()
	conn, err := models.Connect()
	if err != nil {
		log.Fatal(err)
	}
	env := &Env{
		users: models.UserModel{DB: conn.DB},
	}

	mux.HandleFunc("/api/register", api.RegisterRoute)
	// mux.HandleFunc("/:link", resolveLinkRoute)

	err = http.ListenAndServe(":5555", mux)
	if err != nil {
		log.Fatal(err)
	}
}
