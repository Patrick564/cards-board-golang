package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Patrick564/cards-board-golang/api"
	"github.com/Patrick564/cards-board-golang/models"

	_ "github.com/Patrick564/cards-board-golang/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Swagger Example API
//	@version		0.5
//	@description	API documentation.
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:5555
//	@BasePath	/api

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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
	cardEnv := &api.CardEnv{
		Cards: models.CardModel{DB: conn, Ctx: ctx},
	}

	// Middlewares
	r.Use(middleware.Logger)

	// Swagger documentation
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5555/swagger/doc.json"),
	))

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", userEnv.Login)

		r.Post("/register", userEnv.Register)
	})

	// r.Route("/api/{username}", func(r chi.Router) {})

	r.Route("/api/{username}/boards", func(r chi.Router) {
		r.Get("/", boardEnv.GetAll)
		r.Post("/", boardEnv.Create)

		r.Get("/{board_id}", boardEnv.GetOne)
		r.Post("/{board_id}", boardEnv.CreateCard)
		r.Patch("/{board_id}", boardEnv.Update)
		r.Delete("/{board_id}", boardEnv.Delete)
	})

	// private all
	r.Route("/api/{username}/cards", func(r chi.Router) {
		r.Get("/", cardEnv.GetAll)

		r.Get("/{card_id}", cardEnv.GetOne)
		r.Patch("/{card_id}", cardEnv.Update)
		r.Delete("/{card_id}", cardEnv.Delete)
	})

	log.Printf("Start server in port %s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}
