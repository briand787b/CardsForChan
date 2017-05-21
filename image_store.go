package main

import "database/sql"

const pageSize = 25

var globalImageStore ImageStore

type DBImageStore struct {
	db *sql.DB
}

func NewDBImageStore(db *sql.DB) ImageStore {
	return &DBImageStore{
		db: db,
	}
}

func(store *DBImageStore) Find(id int) (*Image, error) {
	row := store.db.QueryRow(
		`
		SELECT id, md5_sum, location, size, created_at
		FROM images
		WHERE id = ?`,
		id,
	)

	image := Image{}
	err := row.Scan(
		&image.ID,
		&image.MD5Sum,
		&image.Location,
		&image.Size,
		&image.CreatedAt,
	)

	return &image, err
}

func(store *DBImageStore) FindAll(offset int) ([]Image, error) {
	rows, err := store.db.Query(
		`
		SELECT id, md5_sum, location, size, created_at
		FROM images
		ORDER BY created_at DESC
		LIMIT ?
		OFFSET ?
		`,
		pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	images := []Image{}
	for rows.Next() {
		image := Image{}
		err := rows.Scan(
			&image.ID,
			&image.MD5Sum,
			&image.Location,
			&image.Size,
			&image.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, err
}

func(store *DBImageStore) FindAllByUser(user *User, offset int) ([]Image, error) {
	rows, err := store.db.Query(
		`
		SELECT id, md5_sum, location, size, created_at
		FROM images
		WHERE user_id = ?
		ORDER BY created_at DESC
		LIMIT ?
		OFFSET ?
		`,
		user.ID,
		pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	images := []Image{}
	for rows.Next() {
		image := Image{}
		err := rows.Scan(
			&image.ID,
			&image.MD5Sum,
			&image.Location,
			&image.Size,
			&image.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}
