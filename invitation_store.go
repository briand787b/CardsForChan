package main

import "database/sql"

type InvitationStore interface {
	Save(*Invitation) error
	Find(string) (*Invitation, error)
	Delete(string) error
}

var globalInvitationStore InvitationStore

type DBInvitationStore struct {
	db *sql.DB
}

func NewDBInvitationStore(db *sql.DB) *DBInvitationStore {
	return &DBInvitationStore{db: db}
}

func (store *DBInvitationStore) Save(inv *Invitation) error {
	var id string
	err := store.db.QueryRow(`
		INSERT INTO invitation
		(id, expiry)
		VALUES
		($1, $2)
		RETURNING id;`,
		inv.ID,
		inv.Expiry,
	).Scan(&id)

	return err
}

func (store *DBInvitationStore) Find(id string) (*Invitation, error) {
	invitation := &Invitation{}
	err := store.db.QueryRow(`
		SELECT id, expiry
		FROM invitation
		WHERE id = $1;`,
		id,
	).Scan(&invitation.ID, &invitation.Expiry)

	if err.Error() == "sql: no rows in result set" {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return invitation, nil
}

func (store *DBInvitationStore) Delete(id string) error {
	_, err := store.db.Query(`
		DELETE FROM invitation
		WHERE id = $1;`,
		id,
	)

	return err
}