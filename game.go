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

// Must have a user account to start a game
// But you do not need one to join
// Need to wrap this whole thing in a transaction
func NewGame(gameName, playerName string, userId int) (*Game, error) {
	if gameName == "" {
		return nil, ErrNoGameName
	}

	user, err := globalUserStore.Find(userId)
	if err != nil {
		return nil, err
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

	err = SaveGameAdmin(player)
	if err != nil {
		return nil, err
	}


	err = globalGameStore.SetAdmin(game.ID, player.ID)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func FindGame(id int) (*Game, error) {
	game, err := globalGameStore.Find(id)
	if err != nil {
		return nil, err
	}

	if game.IsActive {
		return game, nil
	}

	err = DeleteGame(game)
	if err != nil {
		panic(err)
	}

	return nil, errors.New("Game is inactive")
}

func DeleteGame(game *Game) error {
	// This should probably take an id instead of an actual game
	// It's redundant to have two deletes that take a pointer to
	// a game.
	return globalGameStore.Delete(game)
}

//func (game *Game) ValidateUser() {
//
//}
//
//func (game *Game) IsValidUser() bool {
//
//}
//
//func (game *Game) IsValidPlayer() bool {
//
//}

func ShowGameByGameIDUserID(gameID, userID int) (*Game, error){
	return globalGameStore.FindByGameIDUserID(gameID, userID)
}