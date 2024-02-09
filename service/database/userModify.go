package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) modifyUsername(userId uint64, newUsername string) (error) {
	/* modifyUsername modifies the username of the record with given 
	   userId to newUsername, if newUsername wasn't already taken.
	   If newUsername was taken, it raises UsernameFoundError */

	// Implements PUT /users/self

	var user ReducedUser
	err := db.c.QueryRow("SELECT * FROM users WHERE username=?", newUsername).Scan(&user)

	if !errors.Is(err, sql.ErrNoRows) {  // if there is a row or there was an error
		if err != nil { return err } else { return &UsernameFoundError{} }
	}

	_, err = db.c.Exec(`UPDATE users SET username=? WHERE id = ?;`, newUsername, userId)

	return err
}
