package api

import (
	"encoding/json"
	"net/http"

	"github.com/Patrick564/cards-board-golang/models"
	"github.com/go-chi/chi/v5"
)

type CreateCard struct {
	Content string `json:"content"`
	BoardId string `json:"board_id"`
}

type CardEnv struct {
	Cards interface {
		Add(username, boardId, content string) error
		FindAllByUsername(username string) ([]models.Card, error)
		FindAllByBoardId(boardId string) ([]models.Card, error)
		FindOne(id string) (models.Card, error)
		Update(id string) error
		Delete(id string) error
	}
}

func (env *CardEnv) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := chi.URLParam(r, "username")
	card := CreateCard{}

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		return
	}

	err = env.Cards.Add(username, card.BoardId, card.Content)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (env *CardEnv) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := chi.URLParam(r, "username")

	cards, err := env.Cards.FindAllByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cards)
}

func (env *CardEnv) GetAllBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	boardId := chi.URLParam(r, "board_id")

	cards, err := env.Cards.FindAllByBoardId(boardId)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cards)
}

// TODO
func (env *CardEnv) GetOne(w http.ResponseWriter, r *http.Request) {}

func (env *CardEnv) Update(w http.ResponseWriter, r *http.Request) {}

func (env *CardEnv) Delete(w http.ResponseWriter, r *http.Request) {}
