package main

import (
	"time"
)

type Game struct {
	ID 		int
	Name 		string
	AdminPlayerID	int
	IsActive		bool
	CreatedAt 	time.Time
}

// Must have a user account to start a game
// But you do not need one to join
func NewGame(name string, userId int) (*Game, error) {
	if name == "" {
		return nil, ErrNoGameName
	}

	user, err := globalUserStore.Find(userId)
	if err != nil {
		return nil, err
	}

	
}

