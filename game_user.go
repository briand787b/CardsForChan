package main

type GameUser struct {
	Game Game
	User User
}

func FindGameUser(gameID, userID int) (*GameUser, error) {
	return globalGameUserStore.Find(gameID, userID)
}

