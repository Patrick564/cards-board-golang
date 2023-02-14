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

	// User routes
	r.HandleFunc("/api/users/register", userEnv.Register)
	r.HandleFunc("/api/users/login", userEnv.Login)

	// Board routes
	r.HandleFunc("/api/boards/{username}", boardEnv.GetAllOrCreate)
	r.HandleFunc("/api/boards/{username}/{board_id}", boardEnv.FindById)

	log.Printf("Start server in port %s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}
