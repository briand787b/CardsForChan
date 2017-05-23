package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

// Probably shouldn't even expose this form since its impossible to
// successfully create a player without satisfying the hidden fields
func HandlePlayerNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "players/new", nil)
}

func HandlePlayerCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// create player
	// redirect player to relevant game url
	// i.e. /game/:gameID for users,
	// 	/game/:gameID/:playerID for non-users
	gameID, err := strconv.Atoi(r.FormValue("gameID"))
	if err != nil {
		// TODO: Make this a bad request instead of 404
		http.NotFound(w, r)
		return
	}

	player := &Player{
		Name: r.FormValue("playerName"),
		GameID: gameID,
		InvitationId: r.FormValue("invitationID"),
	}


}
