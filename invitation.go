/*
Invitation IDs appear in the url: /games/:gameID/:invitationID and allow players
to join games even though they are not logged in as a user
 */
package main

import (
	"time"
)

// Only the game's admin can invite someone to the game
//TODO: Implement in the database
type Invitation struct {
	ID	 	string
	Expiry 		time.Time
	GameID		int
	UserId		int
}

const (
	invitationLength = 4 * time.Hour
	invitationIDLength = 20
)

func NewInvitation() *Invitation {
	return nil
}

func (invitation *Invitation) Expired() bool {
	return invitation.Expiry.Before(time.Now())
}

func FindInvitation(id string) (*Invitation, error) {
	inv, err := globalInvitationStore.Find(id)
	if err != nil {
		return nil, err
	}

	if inv.Expired() {
		globalInvitationStore.Delete(inv)
		return nil, nil
	}

	return inv, nil
}