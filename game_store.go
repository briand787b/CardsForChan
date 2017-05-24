package main

import (
	"database/sql"
	"time"
	"fmt"
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
	row := store.db.QueryRow(
		`
		INSERT INTO game
		(name, is_active, created_at)
		VALUES
		($1, $2, $3)
		RETURNING id;`,
		game.Name,
		true,
		time.Now(),
	)

	err := row.Scan(&game.ID)
	if err != nil {
		return err
	}

	return nil
}

func (store *DBGameStore) Find(id int) (*Game, error) {
	row := store.db.QueryRow(`
		SELECT id, name, admin_player_id, is_active, created_at
		FROM game
		WHERE id = $1;`,
		id,
	)

	game := &Game{}
	err := row.Scan(
		&game.ID,
		&game.Name,
		&game.AdminPlayerID,
		&game.IsActive,
		&game.CreatedAt,
	)

	if err.Error() == "sql: no rows in result set" {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return game, nil
}

func (store DBGameStore) FindByGameIDUserID(gameID, userID int) (*Game, error) {
	fmt.Println("gameID: ", gameID, " || userID: ", userID)
	game := &Game{}
	err := store.db.QueryRow(`
		SELECT g.id, g.admin_player_id, g.name, g.is_active, g.created_at
		FROM usr u
		INNER JOIN player p
		ON u.id = p.usr_id
		INNER JOIN game g
		ON g.id = p.game_id
		WHERE g.id = $1 AND u.id = $2;`,
		gameID,
		userID,
	).Scan(
		&game.ID,
		&game.AdminPlayerID,
		&game.Name,
		&game.IsActive,
		&game.CreatedAt,
	)

	return game, err
}

func (store *DBGameStore) Delete(game *Game) error {
	id := game.ID
	var deleted_id int
	err := store.db.QueryRow(`
		DELETE FROM game
		WHERE id = $1
		RETURNING id;
		`,
		id,
	).Scan(&deleted_id)

	if err != nil {
		return err
	}

	return nil
}

func (store *DBGameStore) SetAdmin(gameID, playerID int) error {
	var updatedID int
	err := store.db.QueryRow(`
		UPDATE game
		SET admin_player_id = $1
		WHERE id = $2
		RETURNING id;`,
		playerID,
		gameID,
	).Scan(&updatedID)

	return err
}

