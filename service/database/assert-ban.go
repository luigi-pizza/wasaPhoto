package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) IsBanned(bannerId uint64, bannedId uint64) (bool, error) {
	// checks if bannedId was banned by bannerId

	var isbanned bool
	err := db.c.QueryRow(
		"SELECT TRUE FROM bans WHERE bannerId = ? AND bannedId = ?",
		bannerId, bannedId).Scan(&isbanned)
	
	if errors.Is(err, sql.ErrNoRows) {return false, nil}
	return isbanned, err
}
