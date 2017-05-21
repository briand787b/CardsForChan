package main

import (
	"database/sql"
)

type SessionStore interface {
	Find(string) (*Session, error)
	Save(*Session) error
	Delete(*Session) error
}

type DBSessionStore struct {
	db *sql.DB
}

var globalSessionStore SessionStore

func NewDBSessionStore(db *sql.DB) (*DBSessionStore) {
	store := &DBSessionStore{
		db: db,
	}

	return store
}

func (store *DBSessionStore) Find(id string) (*Session, error) {
	var sess Session
	rows, err := store.db.Query(
		`SELECT id, usr_id, expiry
		FROM session
		WHERE id = $1;`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&sess.ID, &sess.UserID, &sess.Expiry)
		if err != nil {
			return nil, err
		}
	}

	return &sess, nil
}

func (store *DBSessionStore) Save(session *Session) error {
	_, err := store.db.Query(
		`INSERT INTO session
		(id, usr_id, expiry)
		VALUES
		($1, $2, $3);`,
		session.ID,
		session.UserID,
		session.Expiry,
	)
	if err != nil {
		return err
	}

	return nil
}

func (store *DBSessionStore) Delete(session *Session) error {
	_, err := store.db.Query(
		`DELETE FROM session
		WHERE id = $1;`,
		session.ID,
	)

	return err
}
