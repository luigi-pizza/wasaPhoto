package database

func (db *appdbimpl) isLiked_by_uid (user_id uint64, post_id uint64) (bool, error) {
	// checks if post_id is liked by user_id

	var isliked bool
	err := db.c.QueryRow(`Select 1 from likes where photoId = ? and userId = ?`, post_id, user_id).Scan(&isliked)
	return isliked, err
}

func (db *appdbimpl) likePhoto (user_id uint64, post_id uint64) error {
	// Inserts a record that user_id has liked post_id, if it didn't already exists.
	// Implements PUT /photos/{postID}/likes/self

	_, err := db.c.Exec(
		"INSERT OR IGNORE INTO likes (user_id, post_id) VALUES (?, ?);",
		user_id, post_id)

	return err
}

func (db *appdbimpl) unlikePhoto (user_id uint64, post_id uint64) error {
	// Removes the record that user_id had liked post_id, if it was present.
	// Implements DELETE /photos/{postID}/likes/self

	_, err := db.c.Exec(
		"DELETE FROM likes WHERE user_id = ? AND post_id = ?;", 
		user_id, post_id)

	return err
}