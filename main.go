package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Patrick564/cards-board-golang/api"
	"github.com/Patrick564/cards-board-golang/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file\n")
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	ctx := context.Background()
	mux := http.NewServeMux()

	conn, err := models.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error to connect to database: %s\n", err.Error())
	}
	defer conn.Close()

	log.Println("Connection to database is successfull")

	userEnv := &api.UserEnv{
		Users: models.UserModel{DB: conn, Ctx: ctx},
	}

	// User routes
	mux.HandleFunc("/api/user/register", userEnv.Register)
	mux.HandleFunc("/api/user/login", userEnv.Login)

	log.Printf("Start server in port %s\n", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
