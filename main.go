package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Patrick564/cards-board-golang/api"
	"github.com/Patrick564/cards-board-golang/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file\n")
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	ctx := context.Background()
	r := chi.NewRouter()

	conn, err := models.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error to connect to database: %s\n", err.Error())
	}
	defer conn.Close()

	log.Println("Connection to database is successfull")

	userEnv := &api.UserEnv{
		Users: models.UserModel{DB: conn, Ctx: ctx},
	}
	boardEnv := &api.BoardEnv{
		Boards: models.BoardModel{DB: conn, Ctx: ctx},
	}

	// Middlewares
	r.Use(middleware.Logger)

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/login", userEnv.Login)

		r.Post("/register", userEnv.Register)
	})

	r.Route("/api/boards", func(r chi.Router) {
		r.Get("/{username}", boardEnv.GetAll)
		r.Post("/{username}", boardEnv.Create)

		r.Get("/{username}/{board_id}", boardEnv.GetOne)
		r.Post("/{username}/{board_id}", boardEnv.GetOne)
	})

	log.Printf("Start server in port %s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}
