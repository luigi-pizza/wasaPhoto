package database

func (db *appdbimpl) Delete_comment(commentId uint64, photoId uint64) error {
	// Deletes a comment from the database. if the comment was already deleted
	// Implements DELETE /photos/{postID}/comments/

	isCID, _, _, err := db.IsCommentId(commentId)
	if err != nil {
		return err
	}
	if !isCID {
		return nil
	}

	_, err = db.c.Exec(
		`
		BEGIN TRANSACTION;
		DELETE FROM comments WHERE id = ?;
		UPDATE photos SET comments = comments - 1 WHERE id = ?;
		COMMIT;
		`, commentId, photoId)
	return err
}
