package database

func (db *appdbimpl) Delete_like(user_id uint64, photo_id uint64) error {
	// Removes the record that user_id had liked photo_id, if it was present.
	// Implements DELETE /photos/{postID}/likes/self

	_, err := db.c.Exec(
		`
		BEGIN TRANSACTION;
		DELETE FROM likes WHERE userId = ? AND photoId = ?;
		UPDATE photos SET likes = likes - 1 WHERE id = ?;
		COMMIT;
		`,
		user_id, photo_id, photo_id)

	return err
}
