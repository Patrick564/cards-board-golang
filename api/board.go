package api

import (
	"encoding/json"
	"net/http"

	"github.com/Patrick564/cards-board-golang/models"
	"github.com/Patrick564/cards-board-golang/utils"
)

type BoardEnv struct {
	Boards interface {
		Add(name, email string) error
		FindAll(email string) ([]models.Board, error)
		FindOne(email, boardId string) ([]models.CardsBoard, error)
		Update(board models.Board, userId string) error
	}
}

func (env *BoardEnv) GetAllOrCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var email string
		_ = json.NewDecoder(r.Body).Decode(&email)
		boards, _ := env.Boards.FindAll(email)
		json.NewEncoder(w).Encode(boards)
		return
	}

	if r.Method == http.MethodPost {
		var req struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		env.Boards.Add(req.Name, req.Email)

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
