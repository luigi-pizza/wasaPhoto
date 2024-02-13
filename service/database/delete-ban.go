package database

func (db *appdbimpl) Delete_ban(pardonerId uint64, pardonedId uint64) error {
	// Removes the record that pardonerId had banned pardonedId, if it was present.
	// Implements DELETE /users/self/banned_users/{userID}

	_, err := db.c.Exec(
		"DELETE FROM bans WHERE bannerId = ? AND bannedId = ?;",
		pardonerId, pardonedId)

	return err
}
