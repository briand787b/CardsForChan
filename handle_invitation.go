package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func HandleInvitationConsume(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	invitationID := params.ByName("invitationID")
	invitation, err := FindInvitation(invitationID)
	if err != nil {
		panic(err)
	}

	if invitation == nil {
		http.NotFound(w, r)
		return
	}

	// pass the player with gameID and invitationID.  Once the player is created
	// in the database, the invitationID will no longer be available due to the
	// unique constraint in the database
	player := &Player{
		GameID: invitation.GameID,
		InvitationId: invitation.ID,
	}

	RenderTemplate(w, r, "players/new", map[string]interface{}{
		"player": player,
	})
}