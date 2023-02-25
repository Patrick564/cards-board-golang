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
		Add(name, username string) error
		AddCard(content, username, boardId string) error
		FindAll(username string) ([]models.Board, error)
		FindOne(username, id string) ([]models.CardsBoard, error)
		Update(newName, id string) error
		Delete(id string) error
	}
}

// ShowAccount godoc
//
//	@Summary		Create board
//	@Description    create a new board
//	@Tags			boards
//	@Accept			json
//	@Param			request	body	CreateBoard	true	"query params"
//	@Success		201 {json}
//	@Failure		400 {json}
//	@Failure		502 {json}
//	@Router			/api/{username}/boards  [post]
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

// ShowAccount godoc
//
//	@Summary		Create board
//	@Description	create a new board
//	@Tags			boards
//	@Accept			json
//	@Param			request	body	CreateBoard	true	"query params"
//	@Success		200
//	@Router			/api/{username}/boards/{board_id}  [post]
func (env *BoardEnv) CreateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := chi.URLParam(r, "username")
	boardId := chi.URLParam(r, "board_id")

	var content string
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		return
	}

	err = env.Boards.AddCard(content, username, boardId)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ShowAccount godoc
//
//	@Summary		Get all boards
//	@Description	get all boards by username
//	@Tags			boards
//	@Accept			json
//	@Param			request	body	CreateBoard	true	"query params"
//	@Success		200
//	@Router			/api/{username}/boards [get]
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

// ShowAccount godoc
//
//	@Summary		Get all boards
//	@Description	get all boards by username
//	@Tags			boards
//	@Accept			json
//	@Param			request	body	CreateBoard	true	"query params"
//	@Success		200
//	@Router			/api/{username}/boards/{board_id}  [get]
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

// ShowAccount godoc
//
//	@Summary		Get all boards
//	@Description	get all boards by username
//	@Tags			boards
//	@Accept			json
//	@Param			request	body	CreateBoard	true	"query params"
//	@Success		200
//	@Router			/api/{username}/boards/{board_id}  [patch]
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

// ShowAccount godoc
//
//	@Summary		Get all boards
//	@Description	get all boards by username
//	@Tags			boards
//	@Accept			json
//	@Param			request	body	CreateBoard	true	"query params"
//	@Success		200
//	@Router			/api/{username}/boards/{board_id}  [delete]
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
