package main

import (
	"github.com/briand787b/pgInit"
	"database/sql"
	"net/http"
	"github.com/briand787b/middleware"
	"log"
	"github.com/julienschmidt/httprouter"
)

var globalPgDB *sql.DB

func init() {
	var err error
	globalPgDB, err = pgInit.ConnectDefault("CardsForChan")
	if err != nil {
		panic(err)
	}

	// User store initialization
	globalUserStore = NewDBUserStore(globalPgDB)

	// Session store initialization
	globalSessionStore = NewDBSessionStore(globalPgDB)

	// Image store initialization
	globalImageStore = NewDBImageStore(globalPgDB)

	// Game store initialization
	globalGameStore = NewDBGameStore(globalPgDB)

	// Player store initialization
	globalPlayerStore = NewDBPlayerStore(globalPgDB)

	// Invitation store initialization
	globalInvitationStore = NewDBInvitationStore(globalPgDB)
}

func main() {
	defer globalPgDB.Close()
	router := NewRouter()

	router.Handle("GET", "/", HandleHome)
	router.Handle("GET", "/register", HandleUserNew)
	router.Handle("POST", "/register", HandleUserCreate)
	router.Handle("GET", "/login", HandleSessionNew)
	router.Handle("POST", "/login", HandleSessionCreate)
	router.Handle("GET", "/image/:imageID", HandleImageShow)
	router.Handle("GET", "/user/:userID", HandleUserShow)
	router.Handle("GET", "/games/:gameID", HandleGameShow)
	router.Handle("GET", "/games/:gameID/:playerID", HandleGameAnonPlayerShow)
	router.Handle("POST", "/players/create/:invitationID", HandlePlayerCreate)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)

	router.ServeFiles(
		"/im/*filepath",
		http.Dir("data/images/"),
	)

	secureRouter := NewRouter()
	secureRouter.Handle("GET", "/sign-out", HandleSessionDestroy)
	secureRouter.Handle("GET", "/account", HandleUserEdit)
	secureRouter.Handle("POST", "/account", HandleUserUpdate)
	secureRouter.Handle("GET", "/games/new", HandleGameNew)
	secureRouter.Handle("POST", "/games/new", HandleGameCreate)

	middleware := middleware.Middleware{}
	middleware.Add(router)
	middleware.Add(http.HandlerFunc(RequireLogin))
	middleware.Add(secureRouter)

	log.Fatal(http.ListenAndServe(":3000", middleware))
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	return router
}

