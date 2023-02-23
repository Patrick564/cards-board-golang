package api

import (
	"encoding/json"
	"net/http"

	"github.com/Patrick564/cards-board-golang/models"
)

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

	env.Cards.Add("testing", "9f4be1e0-96cc-4d97-b478-a7eaa7e917f1", "card testing")
}

func (env *CardEnv) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cards, err := env.Cards.FindAllByUsername("testing")
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cards)
}
