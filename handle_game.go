package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"strconv"
	"errors"
)

// /games/new
func HandleGameNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "games/new", nil)
}

// /games/new
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

// games/play/:gameID/:playerID
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
		fmt.Println("PlayerID is an empty string")
		user := RequestUser(r)
		if user == nil {
			fmt.Println("User data is unable to be obtained")
			RenderTemplate(w, r, "games/show", map[string]interface{}{
				"Error": errors.New("Unable to find user"),
				"Game": "Not Found",
			})
			return
		}

		fmt.Println("User data has been obtained, user is not null")
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