package database

func (db *appdbimpl) Delete_comment(commentId uint64, photoId uint64) error {
	// Deletes a comment from the database. if the comment was already deleted
	// Implements DELETE /photos/{postID}/comments/
	_, err := db.c.Exec(
		`
		BEGIN TRANSACTION;
		DELETE FROM comments WHERE id = ?;
		UPDATE photos SET comments = comments - 1 WHERE id = ?;
		COMMIT;
		`, commentId, photoId)
	return err
}
