package database

func (db *appdbimpl) Delete_photo(photoId uint64) error {
	_, err := db.c.Exec(
		`
		BEGIN TRANSACTION;
		DELETE FROM photos WHERE id = ?;
		DELETE FROM comments WHERE photoId = ?;
		DELETE FROM likes WHERE photoId= ?;
		COMMIT;
		`,
		photoId, photoId, photoId)
	return err
}
