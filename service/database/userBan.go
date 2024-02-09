package database

func (db *appdbimpl) isBanned_by_uid (possibleBannedId uint64, possibleBannerId uint64) (bool, error) {
	// checks if possibleBannedId was banned by possibleBannerId

	var isbanned bool
	err := db.c.QueryRow(`Select 1 from bans where bannerId = ? and bannedId = ?`, possibleBannerId, possibleBannedId).Scan(&isbanned)
	return isbanned, err
}


func (db *appdbimpl) banUser (bannerId uint64, bannedId uint64) error {
	// Inserts a record that bannerId has banned bannedId, if it didn't already exists.
	// Implements PUT /users/self/banned_users/{userID}

	_, err := db.c.Exec(
		`
		BEGIN TRANSACTION;
		INSERT OR IGNORE INTO bans (bannerId, bannedId) VALUES (?, ?);
		DELETE FROM follows WHERE followerId = ? AND followedId = ?;
		COMMIT;
		`,
		bannerId, bannedId,
		bannedId, bannerId,
	)

	return err
}

func (db *appdbimpl) pardonUser (pardonerId uint64, pardonedId uint64) error {
	// Removes the record that pardonerId had banned pardonedId, if it was present.
	// Implements DELETE /users/self/banned_users/{userID}

	_, err := db.c.Exec(
		"DELETE FROM bans WHERE bannerId = ? AND bannedId = ?;", 
		pardonerId, pardonedId)

	return err
}