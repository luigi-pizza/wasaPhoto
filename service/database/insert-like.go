package database

func (db *appdbimpl) Insert_like(user_id uint64, photo_id uint64) error {
	// Inserts a record that user_id has liked photo_id, if it didn't already exists.
	// Implements PUT /photos/{postID}/likes/self

	_, err := db.c.Exec(
		`
		BEGIN TRANSACTION;
		INSERT OR IGNORE INTO likes (userId, photoId) VALUES (?, ?);
		UPDATE photos SET likes = likes + 1 WHERE id = ?
		COMMIT;
		`,
		user_id, photo_id, photo_id)

	return err
}
