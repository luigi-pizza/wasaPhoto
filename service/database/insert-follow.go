package database

func (db *appdbimpl) Insert_follow(followerId uint64, followedId uint64) error {
	// Inserts a record that followerId has followed followedId, if it didn't already exists.
	// Implements PUT /users/self/followed_users/{userID}

	_, err := db.c.Exec(
		"INSERT OR IGNORE INTO follows (followerId, followedId) VALUES (?, ?);",
		followerId, followedId)

	return err
}
