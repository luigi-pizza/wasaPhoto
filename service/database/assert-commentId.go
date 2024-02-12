package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) IsCommentId(comment_id uint64) (bool, uint64, uint64, error) {

	var (
		isCID bool
		authorId uint64
		photoId  uint64
	)
	err := db.c.QueryRow(
		`SELECT TRUE, authorId, photoId FROM comments WHERE id = ?`, 
		comment_id ).Scan(&isCID, &authorId, &photoId)
	
	if errors.Is(err, sql.ErrNoRows) {return false, 0, 0, nil}
	return isCID, authorId, photoId, err
}
