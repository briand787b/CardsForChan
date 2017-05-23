package main

import "net/http"

//TODO: Implement InvitationID in database schema: invitation_id CHAR(n) UNIQUE REFERENCES invitation(id)
type Player struct {
	ID 	int
	GameID	int
	Name 	string
	UserID 	int
	InvitationId string
}

func NewPlayer(player *Player) error {
	// validate gameID and validationID
	game, err := FindGame(player.GameID)
	if err != nil {
		return err
	}

	if game == nil {
		// if game cannot be found then just redirect home
		return ErrGameNotFound
	}

	invitation, err := FindInvitation(player.InvitationId)
	if err != nil {
		return err
	}

	if invitation == nil {
		return Err
	}
}

