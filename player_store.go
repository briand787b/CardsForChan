package main

import "database/sql"

type PlayerStore interface {
	Save(*Player) error
	Find(int) (*Player, error)
}

type DBPlayerStore struct {
	db *sql.DB
}

var globalPlayerStore PlayerStore

func NewDBPlayerStore(db *sql.DB) (*DBPlayerStore) {
	return &DBPlayerStore{db: db}
}

func (store *DBPlayerStore) Save(player *Player) error {
	err := store.db.QueryRow(`
		INSERT INTO player
		(name, game_id, usr_id)
		VALUES
		($1, $2, $3)
		RETURNING id;`,
		player.Name,
		player.GameID,
		player.UserID,
	).Scan(&player.ID)

	if err != nil {
		return err
	}

	return nil
}

func (store *DBPlayerStore) Find(id int) (*Player, error) {
	player := &Player{}
	err := store.db.QueryRow(`
		SELECT id, name, game_id, usr_id
		FROM player
		WHERE id = $1;`,
		id,
	).Scan(&player)

	if err != nil {
		return nil, err
	}

	return player, nil
}