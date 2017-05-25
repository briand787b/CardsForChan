package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
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
		InvitationID: invitation.ID,
	}

	RenderTemplate(w, r, "players/new", map[string]interface{}{
		"Player": player,
	})
}

//TODO: Make the invitation creation process be controlled through javascript
// /invitations/new/:gameID GET
func HandleInvitationNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	gameID, err := strconv.Atoi(params.ByName("gameID"))
	if err != nil {
		http.Redirect(w, r, "/?flash=invalid+game+ID", http.StatusBadRequest)
		return
	}

	game, err := ShowGameByGameIDUserID(gameID, RequestUser(r).ID)
	if err != nil {
		http.Redirect(w, r, "/?flash=unauthorized+to+create+invitations", http.StatusBadRequest)
		return
	}

	RenderTemplate(w, r, "invitations/new", map[string]interface{}{
		"Game": game,
	})
}

// /invitations/new/:gameID POST
func HandleInvitationCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}