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

	var user schema.ReducedUser
	err := db.c.QueryRow("SELECT * FROM users WHERE username = ?", newUsername).Scan(&user)

	if !errors.Is(err, sql.ErrNoRows) { // if there is a row or there was an error
		if err != nil {
			return err
		} else {
			return schema.ErrUsernameAlreadyInUse
		}
	}

	_, err = db.c.Exec(`UPDATE users SET username = ? WHERE id = ?;`, newUsername, userId)

	return err
}
