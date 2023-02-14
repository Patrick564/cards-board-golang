package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Patrick564/cards-board-golang/models"
	"github.com/Patrick564/cards-board-golang/utils"
)

type BoardEnv struct {
	Boards interface {
		Add(name, email string) error
		FindAll(username string) ([]models.Board, error)
		FindOne(email, boardId string) ([]models.CardsBoard, error)
		Update(board models.Board, userId string) error
	}
}

func (env *BoardEnv) GetAllOrCreate(w http.ResponseWriter, r *http.Request) {
	log.Print("Loading route /api/boards/{username} \n")

	path := strings.TrimPrefix(r.URL.Path, "/api/boards/")

	if r.Method == http.MethodGet {
		// username := strings.TrimSuffix(path, "/")
		username := r.URL.Path[len("/api/boards/"):]
		fmt.Println(username)

		boards, err := env.Boards.FindAll("testing")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(boards)
		return
	}

	if r.Method == http.MethodPost {
		params := strings.Split(strings.TrimSuffix(path, "/"), "/")

		fmt.Println(params)

		// env.Boards.Add(params[1], params[0])

		w.WriteHeader(http.StatusCreated)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(utils.CustomError{Message: "method not allowed"})
}

func (env *BoardEnv) FindById(w http.ResponseWriter, r *http.Request) {
	cards, _ := env.Boards.FindOne("testing@gmail.com", "20b5d21f-e292-4db2-8baa-b7e850854fae")

	json.NewEncoder(w).Encode(cards)
}
