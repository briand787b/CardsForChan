package main

import (
	"time"
	"errors"
)

type Game struct {
	ID 		int
	Name 		string
	CreatedAt 	time.Time
}

func NewGame(name string, userId int) (*Game, error) {
	if name == "" {
		return nil, ErrNoGameName
	}

	user, err := globalUserStore.Find(userId)
	if err != nil {
		return nil, errors.New("cannot find user by ID")
	}

	
}

