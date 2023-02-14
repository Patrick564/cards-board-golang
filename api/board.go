package api

import (
	"encoding/json"
	"net/http"

	"github.com/Patrick564/cards-board-golang/models"
	"github.com/Patrick564/cards-board-golang/utils"
	"github.com/go-chi/chi/v5"
)

type BoardEnv struct {
	Boards interface {
		Add(name, email string) error
		FindAll(username string) ([]models.Board, error)
		FindOne(email, boardId string) ([]models.CardsBoard, error)
		Update(board models.Board, userId string) error
	}
}

func (env *BoardEnv) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := chi.URLParam(r, "username")

	var boardName string
	err := json.NewDecoder(r.Body).Decode(&boardName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	err = env.Boards.Add(boardName, username)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (env *BoardEnv) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := chi.URLParam(r, "username")
	boards, err := env.Boards.FindAll(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(boards)
}

func (env *BoardEnv) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := chi.URLParam(r, "username")
	boardId := chi.URLParam(r, "board_id")

	cards, err := env.Boards.FindOne(username, boardId)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cards)
}
