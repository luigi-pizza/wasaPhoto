package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) IsLiked(user_id uint64, photo_id uint64) (bool, error) {
	// checks if photo_id is liked by user_id

	var isliked bool
	err := db.c.QueryRow(
		`SELECT TRUE FROM likes WHERE photoId = ? AND userId = ?`, 
		photo_id, user_id).Scan(&isliked)

	if errors.Is(err, sql.ErrNoRows) {return false, nil}
	return isliked, err
}
