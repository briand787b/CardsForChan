package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"strconv"
	"errors"
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
			"Error": err,
			"GameName": gameName,
			"PlayerName": playerName,
		})
	}

	gameURL := fmt.Sprint("/games/play/", game.ID)
	http.Redirect(w, r, gameURL, http.StatusFound)
}

func HandleGameShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// this handler is for users only
	gameID, err := strconv.Atoi(params.ByName("gameID"))
	if err != nil {
		// TODO: make this a bad request instead of a not found
		http.NotFound(w, r)
		return
	}

	playerIDParam := params.ByName("playerID")
	if playerIDParam == "" {
		user := RequestUser(r)
		if user == nil {
			RenderTemplate(w, r, "games/show", map[string]interface{}{
				"Error": errors.New("Unable to find user"),
				"Game": "Not Found",
			})
		}

		game, err := ShowGameByGameIDUserID(gameID, user.ID)
		RenderTemplate(w, r, "games/show", map[string]interface{}{
			"Game": game,
			"User": user,
			"Error": err,
		})
		return
	}


}


func HandleGameShowWithUser(w http.ResponseWriter, gameID, userID int) {

}

func HandleGameShowWithNonUser(w http.ResponseWriter, gameID, userID int) {

}