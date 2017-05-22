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
}

const (
	invitationLength = 8 * time.Hour
	invitationIDLength = 20
)

func NewInvitation() *Invitation {
	return nil
}