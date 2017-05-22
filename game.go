package main

import (
	"time"
	"errors"
)

type Game struct {
	ID 		int
	Name 		string
	AdminPlayerID	int
	IsActive	bool
	CreatedAt 	time.Time
}

// Need to wrap this whole thing in a transaction
func NewGame(gameName, playerName string, userId int) (*Game, error) {
	if gameName == "" {
		return nil, ErrNoGameName
	}

	user, err := globalUserStore.Find(userId)
	if err != nil {
		return nil, errors.New("cannot find user by ID")
	}

	game := &Game{
		Name:     gameName,
		IsActive: true,
	}

	err = globalGameStore.Save(game)
	if err != nil {
		return nil, err
	}

	player := &Player{
		Name: playerName,
		GameID: game.ID,
		UserID: user.ID,
	}

	err = globalPlayerStore.Save(player)
	if err != nil {
		return nil, err
	}

	err = globalGameStore.SetAdmin(game.ID, player.ID)
	if err != nil {
		return nil, err
	}

	return game, nil
}

