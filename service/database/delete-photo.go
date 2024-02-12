package database

func (db *appdbimpl) Delete_photo (photoId uint64) error {
	_, err := db.c.Exec(`
	DELETE FROM photos WHERE id = ?
	DELETE FROM comments WHERE photoId = ?
	DELETE FROM likes WHERE photoId= ?
	`, photoId, photoId, photoId)
	return err
}
