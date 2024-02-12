package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) IsFollowed (followerId uint64, followedId uint64) (bool, error) {
	// checks if followedId is followed by followerId

	var isfollowed bool
	err := db.c.QueryRow(
		`SELECT TRUE FROM follows WHERE followerId = ? AND followedId = ?`, 
		followerId, followedId).Scan(&isfollowed)
	
	if errors.Is(err, sql.ErrNoRows) {return false, nil}
	return isfollowed, err
}