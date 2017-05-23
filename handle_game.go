package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func HandleGameNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "games/new", nil)
}

func HandleGameCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	gameName := r.FormValue("gameName")
	playerName := r.FormValue("playerName")
	user := RequestUser(r)
	game, err := NewGame(gameName, playerName, user.ID)
	if err != nil {
		RenderTemplate(w, r, "games/new", map[string]interface{}{
			"gameName": gameName,
			"playerName": playerName,
		})
	}

	gameURL := fmt.Sprint("/games/", game.ID)
	http.Redirect(w, r, gameURL, http.StatusFound)
}

func HandleGameShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// this handler is for users only
}

func HandleGameAnonPlayerShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// this handler extracts the player id from the url, so non-users can play
}