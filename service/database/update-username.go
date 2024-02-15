package database

import (
	"database/sql"
	"errors"

	"github.com/luigi-pizza/wasaPhoto/service/components/schema"
)

func (db *appdbimpl) Update_username(userId uint64, newUsername string) error {
	/* updateUsername modifies the username of the record with given
	   userId to newUsername, if newUsername wasn't already taken.
	   If newUsername was taken, it raises UsernameFoundError */

	// Implements PUT /users/self

	var uid uint64
	err := db.c.QueryRow("SELECT users.id FROM users WHERE username = ?", newUsername).Scan(&uid)

	if !errors.Is(err, sql.ErrNoRows) { // if there is a row or there was an error
		if err != nil {
			return err
		} else {
			return schema.ErrUsernameAlreadyInUse
		}
	}

	_, err = db.c.Exec(`UPDATE users SET username = ? WHERE id = ?;`, newUsername, userId)

	if err != nil {
		return err
	}

	return err
}
