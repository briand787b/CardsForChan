package main

//TODO: Implement InvitationID in database schema: invitation_id CHAR(n) UNIQUE REFERENCES invitation(id)
type Player struct {
	ID 	int
	GameID	int
	Name 	string
	UserID 	int
}

