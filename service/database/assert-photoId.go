package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) IsPhotoId(photo_id uint64) (bool, uint64, error) {

	var (
		isPID    bool
		authorId uint64
	)
	err := db.c.QueryRow(
		`SELECT TRUE, authorId FROM photos WHERE id = ?`,
		photo_id).Scan(&isPID, &authorId)

	if errors.Is(err, sql.ErrNoRows) {
		return false, 0, nil
	}
	return isPID, authorId, err
}
