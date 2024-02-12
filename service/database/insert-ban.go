package database

func (db *appdbimpl) Insert_ban(bannerId uint64, bannedId uint64) error {
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
