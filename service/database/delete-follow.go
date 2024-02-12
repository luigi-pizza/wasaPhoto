package database

func (db *appdbimpl) Delete_follow(followerId uint64, followedId uint64) error {
	// Removes the record that followerId had followed followedId, if it was present.
	// Implements DELETE /users/self/followed_users/{userID}

	_, err := db.c.Exec(
		"DELETE FROM follows WHERE followerId = ? AND followedId = ?;",
		followerId, followedId)

	return err
}
