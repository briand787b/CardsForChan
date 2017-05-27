package main

import (
	"database/sql"
	"time"
)

type GameStore interface {
	Save(*Game) error
	Find(int) (*Game, error)
	FindByGameIDUserID(int, int) (*Game, error)
	Delete(*Game) error
	SetAdmin(int, int) error
}

type DBGameStore struct {
	db *sql.DB
}

var globalGameStore GameStore

func NewDBGameStore(db *sql.DB) *DBGameStore {
	return &DBGameStore{db: db}
}

func (store *DBGameStore) Save(game *Game) error {
	return store.db.QueryRow(
		`
		INSERT INTO game
		(name, is_active, created_at)
		VALUES
		($1, $2, $3)
		RETURNING id;
		`,
		game.Name,
		true,
		time.Now(),
	).Scan(&game.ID)
}

func (store *DBGameStore) Find(id int) (*Game, error) {
	game := &Game{}
	err := store.db.QueryRow(`
		SELECT id, name, admin_player_id, is_active, created_at
		FROM game
		WHERE id = $1;
		`,
		id,
	).Scan(
		&game.ID,
		&game.Name,
		&game.AdminPlayerID,
		&game.IsActive,
		&game.CreatedAt,
	)

	// TODO: Find a cleaner way to handle empty result sets
	if err.Error() == "sql: no rows in result set" {
		return nil, nil
	}

	return game, nil
}

func (store DBGameStore) FindByGameIDUserID(gameID, userID int) (*Game, error) {
	game := &Game{}
	err := store.db.QueryRow(`
		SELECT g.id, g.admin_player_id, g.name, g.is_active, g.created_at
		FROM usr u
		INNER JOIN player p
		ON u.id = p.usr_id
		INNER JOIN game g
		ON g.id = p.game_id
		WHERE g.id = $1 AND u.id = $2;
		`,
		gameID,
		userID,
	).Scan(
		&game.ID,
		&game.AdminPlayerID,
		&game.Name,
		&game.IsActive,
		&game.CreatedAt,
	)

	// TODO: Find a cleaner way to handle empty result sets
	// if err.Error() == "sql: no rows in result set" {
	//	return nil, nil
	//}

	return game, err
}

func (store *DBGameStore) Delete(game *Game) error {
	id := game.ID
	var deleted_id int
	return store.db.QueryRow(`
		DELETE FROM game
		WHERE id = $1
		RETURNING id;
		`,
		id,
	).Scan(&deleted_id)
}

func (store *DBGameStore) SetAdmin(gameID, playerID int) error {
	var updatedID int
 	return store.db.QueryRow(`
		UPDATE game
		SET admin_player_id = $1
		WHERE id = $2
		RETURNING id;
		`,
		playerID,
		gameID,
	).Scan(&updatedID)
}
