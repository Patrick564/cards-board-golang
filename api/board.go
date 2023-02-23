package api

import (
	"encoding/json"
	"net/http"

	"github.com/Patrick564/cards-board-golang/models"
	"github.com/Patrick564/cards-board-golang/utils"
	"github.com/go-chi/chi/v5"
)

type CreateBoard struct {
	Name string `json:"name"`
}

type BoardEnv struct {
	Boards interface {
		Add(name, email string) error
		Delete(id string) error
		FindAll(username string) ([]models.Board, error)
		FindOne(email, boardId string) ([]models.CardsBoard, error)
		Update(newName, id string) error
	}
}

func (env *BoardEnv) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := chi.URLParam(r, "username")
	board := CreateBoard{}

	err := json.NewDecoder(r.Body).Decode(&board)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	err = env.Boards.Add(board.Name, username)
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

func (env *BoardEnv) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	boardId := chi.URLParam(r, "board_id")
	board := CreateBoard{}

	err := json.NewDecoder(r.Body).Decode(&board)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	err = env.Boards.Update(board.Name, boardId)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (env *BoardEnv) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	boardId := chi.URLParam(r, "board_id")

	err := env.Boards.Delete(boardId)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
