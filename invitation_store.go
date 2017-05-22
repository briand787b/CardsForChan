package main

import "database/sql"

type InvitationStore interface {
	Save(*Invitation) error
	Find(string) (*Invitation, error)
	Revoke(string) error
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
		(id, is_revoked, expiry)
		VALUES
		($1, $2, $3)
		RETURNING id;`,
		inv.ID,
		inv.IsRevoked,
		inv.Expiry,
	).Scan(&id)

	return err
}

func (store *DBInvitationStore) Find(id string) (*Invitation, error) {
	invitation := &Invitation{}
	err := store.db.QueryRow(`
		SELECT id, is_revoked, expiry
		FROM invitation
		WHERE id = $1;`,
		id,
	).Scan(&invitation)

	if err.Error() == "sql: no rows in result set" {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return invitation, nil
}

func (store *DBInvitationStore) Revoke(id string) error {
	return nil
}



