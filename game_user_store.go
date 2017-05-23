package main

import "database/sql"

type GameUserStore interface {
	Find(int, int) (*GameUser, error)
}

type DBGameUserStore struct {
	db *sql.DB
}

var globalGameUserStore GameUserStore

func NewDBGameUserStore(db *sql.DB) *DBGameUserStore {
	return &DBGameUserStore{
		db: db,
	}
}

func (store DBGameUserStore) Find(gameID, userID int) (*GameUser, error) {
	gameUser := &GameUser{}
	err := store.db.QueryRow(`
		SELECT g.id, g.name, g.admin_player_id, g.is_active, g.created_at,
			u.id, u.email, u.username, u.created_at, u.modified_at
		FROM user u
		INNER JOIN player p
		ON u.id = p.usr_id
		INNER JOIN game g
		ON g.id = p.game_id
		WHERE g.id = $1 AND p.id = $2;`,
		gameID,
		userID,
	).Scan(
		&gameUser.Game.ID,
		&gameUser.Game.Name,
		&gameUser.Game.IsActive,
		&gameUser.Game.CreatedAt,
		&gameUser.User.ID,
		&gameUser.User.Email,
		&gameUser.User.Username,
		&gameUser.User.CreatedAt,
		&gameUser.User.ModifiedAt,
	)

	return gameUser, err
}