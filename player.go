package main

import "errors"

type Player struct {
	ID 	int
	GameID	int
	Name 	string
	UserID 	int
	InvitationID string
}

// might want to convert the inputs to the fields instead of just the object
func NewPlayer(gameID int, user *User, name, invitationID string, ) (*Player, error) {
	if name == "" {
		return nil, ErrInvalidPlayerName
	}

	// validate gameID and validationID
	game, err := FindGame(gameID)
	if err != nil {
		return nil, err
	}

	if game == nil {
		// if game cannot be found then just redirect home
		return nil, errors.New("Game not found")
	}

	invitation, err := FindInvitation(invitationID)
	if err != nil {
		return nil, err
	}

	if invitation == nil {
		return nil, errors.New("Invalid invitation")
	}

	if !invitation.IsValidReceiver(user) {
		return nil, errors.New("You are not allowed to receive this invitation")
	}

	player := &Player{
		GameID: gameID,
		Name: name,
		InvitationID: invitationID,
	}

	if user == nil {
		return player, globalPlayerStore.SaveWithoutUser(player)
	}

	player.UserID = user.ID
	return player, globalPlayerStore.SaveWithUser(player)
}

func SaveGameAdmin(player *Player) error {
	return globalPlayerStore.SaveAdmin(player)
}