package main

import (
	"database/sql"
	"time"
)

type GameStore interface {
	Save(*Game) error
	Find(int) (*Game, error)
	Delete(int) error
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
		(name, created_at)
		VALUES
		($1, $2)
		RETURNING id;`,
		game.Name,
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
		SELECT id, name, created_at
		FROM game
		WHERE id = $1;`,
		id,
	)

	game := &Game{}
	err := row.Scan(
		&game.ID,
		&game.Name,
		&game.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (store *DBGameStore) Delete(id int) error {
	var delted_id int
	err := store.db.QueryRow(`
		DELETE FROM game
		WHERE id = $1
		RETURNING id;
		`,
		id,
	).Scan(&delted_id)

	if err != nil {
		return err
	}

	return nil
}