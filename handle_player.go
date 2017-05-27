package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"fmt"
	"github.com/briand787b/validation"
)

// Probably shouldn't even expose this form since its impossible to
// successfully create a player without satisfying the hidden fields
func HandlePlayerNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "players/new", nil)
}
// /players/create/:invitationID
func HandlePlayerCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// create player
	// redirect player to relevant game url
	// i.e. /games/:gameID for users,
	// 	/games/:gameID/:playerID for non-users
	gameID, err := strconv.Atoi(r.FormValue("gameID"))
	if err != nil {
		// TODO: Make this a bad request instead of 404
		http.NotFound(w, r)
		return
	}

	user := RequestUser(r)
	player, err := NewPlayer(gameID, user, r.FormValue("name"), r.FormValue("invitationID"))
	if validation.IsValidationError(err) {
		RenderTemplate(w, r, "players/new", map[string]interface{}{
			"player": player,
		})
		return
	}

	gameURL := fmt.Sprint("/games/", player.GameID)
	if player.UserID != 0 {
		gameURL += fmt.Sprint("/", player.ID)
	}

	http.Redirect(w, r, gameURL, http.StatusFound)
}
