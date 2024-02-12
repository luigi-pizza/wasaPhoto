package database



func (db *appdbimpl) Insert_photo(authorId uint64, caption string, creation int64) (uint64, error) {
	result, err := db.c.Exec(
		"INSERT INTO photos (authorId, caption, timeOfCreation, likes, comments) VALUES (?, ?, ?, 0, 0)",
		authorId, caption, creation)
	if err != nil {return 0, err}

	id, err := result.LastInsertId()
	if err != nil {return 0, err}

	return uint64(id), nil
}
