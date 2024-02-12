package database

import "github.com/luigi-pizza/wasaPhoto/service/components/schema"

func (db *appdbimpl) Select_commentList(user_id uint64, post_id uint64, page_numb uint64) (schema.CommentList, error) {
	// return the comments under a post.
	// Implements GET /photos/{postID}/comments/

	var commentList schema.CommentList

	rows, err := db.c.Query(`
		SELECT 
			comments.id, 
			comments.authorId, users.username,
			comments.commentText, comments.timeOfCreation
		FROM 
			comments INNER JOIN users ON users.id = comments.authorId 
		WHERE 
			photoId = ? AND
			NOT EXISTS (
				SELECT 1 FROM bans WHERE bannerId = comments.authorId AND bannedId = ?
			)
		ORDER BY 
			comments.timeOfCreation DESC
		LIMIT 24 OFFSET ?`, 
		post_id, user_id, 24*page_numb)
	
	if err != nil {return commentList, err}
	defer rows.Close()

	commentList.PageNumber = page_numb
	for rows.Next() {
		var (
			commentId uint64
			authorId  uint64
			username  string
			text      string
			creation  int64
		)
		
		if err := rows.Scan(&commentId, &authorId, &username, &text, &creation); err != nil {return commentList, err}

		commentList.Comments = append(commentList.Comments, schema.Comment {
			Id: commentId,
			Author: schema.ReducedUser{Id:authorId, Username: username},
			CommentText: text,
			TimeOfCreation: creation,
		})
	}

	if err := rows.Err(); err != nil {return commentList, err}
	return commentList, err
}
