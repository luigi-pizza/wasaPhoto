package database

func (db *appdbimpl) getComments_by_postId (user_id uint64, post_id uint64) ([]CompleteComment, error) {
	// checks if post_id is liked by user_id

	var comments []CompleteComment
	rows, err := db.c.Query(`
		SELECT * FROM comments WHERE photoId = ? AND
		NOT EXISTS (SELECT 1 FROM bans WHERE 
			bannerId = comments.authorId AND bannedId = ?)
		LIMIT 24`, post_id, user_id)
	if (err != nil) {return nil, err}

	defer rows.Close()
	for rows.Next() {
		var comment CompleteComment

		if err := rows.Scan(&comment.id, &comment.username); err != nil {return nil, err}
		result = append(result, user)
	}
	return result, err
}

func (db *appdbimpl) likePhoto (user_id uint64, post_id uint64) error {
	// Inserts a record that user_id has liked post_id, if it didn't already exists.
	// Implements PUT /photos/{postID}/likes/self

	_, err := db.c.Exec(
		"INSERT OR IGNORE INTO likes (user_id, post_id) VALUES (?, ?);",
		user_id, post_id)
s
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