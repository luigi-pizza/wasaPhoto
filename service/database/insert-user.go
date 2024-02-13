package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) UserId_by_username(username string) (uint64, error) {
	// userId_by_username returns the ReducedUser with the given username
	var id uint64
	err := db.c.QueryRow("SELECT id FROM users WHERE username=?", username).Scan(&id)
	return id, err
}

func (db *appdbimpl) Insert_user(username string) (uint64, bool, error) {
	// InsertUser inserts with autoincremented id a new user record with username `username`.
	// before inserting the user, it checks that the username isn't already in use.
	// If it is, returns the identifier in the stored record and doesn't insert a new record.
	// Otherwise, it inserts the new record and returns the new user_id.

	id, err := db.UserId_by_username(username)
	if errors.Is(err, sql.ErrNoRows) {

		result, err := db.c.Exec(`INSERT INTO users(username) VALUES (?);`, username)
		if err != nil {
			return id, false, err
		}

		res, err := result.LastInsertId()
		if err != nil {
			return id, false, err
		}

		return uint64(res), true, nil
	}
	if err != nil {
		return id, false, err
	}
	return id, false, nil

}
