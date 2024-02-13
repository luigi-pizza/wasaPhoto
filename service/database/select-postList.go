package database

import (
	"github.com/luigi-pizza/wasaPhoto/service/components/schema"
)

func (db *appdbimpl) Select_postList(requestingUser uint64, requestedUser uint64, page_numb uint64) (schema.PostList, error) {
	var postList schema.PostList

	rows, err := db.c.Query(`
		SELECT 
			photos.id, 
			users.username,
			photos.caption, photos.timeOfCreation, photos.likes, photos.comments
			(SELECT EXISTS(SELECT TRUE FROM likes WHERE photoId = photos.id AND userId = ?)) AS isLiked
		FROM 
			photos INNER JOIN users ON users.id = photos.authorId 
		WHERE 
			authorId = ? AND
		ORDER BY 
			photos.timeOfCreation DESC
		LIMIT 24 OFFSET ?`,
		requestingUser, requestedUser, 24*page_numb)

	if err != nil {
		return postList, err
	}
	defer rows.Close()

	postList.PageNumber = page_numb
	for rows.Next() {
		var (
			photoid        uint64
			username       string
			caption        string
			timeOfCreation int64
			likes          uint64
			comments       uint64
			isLiked        bool
		)

		if err := rows.Scan(&photoid, &username, &caption, &timeOfCreation, &likes, &comments, &isLiked); err != nil {
			return postList, err
		}

		postList.Posts = append(postList.Posts, schema.Post{
			Id:             photoid,
			Author:         schema.ReducedUser{Id: requestedUser, Username: username},
			Caption:        caption,
			Likes:          likes,
			Comments:       comments,
			TimeOfCreation: timeOfCreation,
			IsLiked:        isLiked,
		})
	}

	if err := rows.Err(); err != nil {
		return postList, err
	}
	return postList, err
}
