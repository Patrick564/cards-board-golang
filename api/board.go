package api

import "github.com/Patrick564/cards-board-golang/models"

type BoardEnv struct {
	Boards interface {
		Add(user models.User) error
		Find(email, password string) (models.User, error)
	}
}

func (env *BoardEnv) Add() {}

func (env *BoardEnv) Find() {}
