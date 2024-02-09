package database

func (db *appdbimpl) isFollowed_by_uid (possibleFollowedId uint64, possibleFollowerId uint64) (bool, error) {
	// checks if possibleFollowedId is followed by possibleFollowerId

	var isfollowed bool
	err := db.c.QueryRow(`Select 1 from follows where followerId = ? and followedId = ?`, possibleFollowerId, possibleFollowedId).Scan(&isfollowed)
	return isfollowed, err
}

func (db *appdbimpl) followUser (followerId uint64, followedId uint64) error {
	// Inserts a record that followerId has followed followedId, if it didn't already exists.
	// Implements PUT /users/self/followed_users/{userID}

	_, err := db.c.Exec(
		"INSERT OR IGNORE INTO follows (followerId, followedId) VALUES (?, ?);",
		followerId, followedId)

	return err
}

func (db *appdbimpl) unfollowUser (followerId uint64, followedId uint64) error {
	// Removes the record that followerId had followed followedId, if it was present.
	// Implements DELETE /users/self/followed_users/{userID}

	_, err := db.c.Exec(
		"DELETE FROM follows WHERE followerId = ? AND followedId = ?;", 
		followerId, followedId)

	return err
}