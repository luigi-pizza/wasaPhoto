package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) IsUserId(user_id uint64) (bool, error) {

	var isUID bool
	err := db.c.QueryRow(
		`SELECT TRUE FROM users WHERE id = ?`,
		user_id).Scan(&isUID)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return isUID, err
}
