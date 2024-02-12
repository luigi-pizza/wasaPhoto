package database

func (db *appdbimpl) Insert_comment(photoId uint64, authorId uint64, text string, timeOfCreation int64) (uint64, error) {
	// Inserts a new comment in the database.
	// Implements POST /photos/{postID}/comments/

	result, err := db.c.Exec(
		`
		BEGIN TRANSACTION;
		INSERT INTO comments (authorId, photoId, commentText, timeOfCreation) VALUES (?, ?, ?, ?)
		UPDATE photos SET comments = comments + 1 WHERE id = ?;
		COMMIT;
		`,
		authorId, photoId, text, timeOfCreation, photoId)
	if err != nil {return 0, err}
	
	id, err := result.LastInsertId()
	if err != nil {return 0, err}

	return uint64(id), nil
}