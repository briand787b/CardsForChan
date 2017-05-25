/*
Invitation IDs appear in the url: /games/:gameID/:invitationID and allow players
to join games even though they are not logged in as a user
 */
package main

import (
	"time"
)

// Only the game's admin can invite someone to the game
type Invitation struct {
	ID             string
	Expiry         time.Time
	GameID         int
	SenderUserID   int
	ReceiverUserID int
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

func (invitation *Invitation) IsValidReceiver(user *User) bool {
	if invitation.ReceiverUserID == 0 {
		return true
	}

	if user == nil {
		return false
	}

	return invitation.ReceiverUserID == user.ID
}

