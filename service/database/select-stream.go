package database

import (
	"github.com/luigi-pizza/wasaPhoto/service/components/schema"
)

func (db *appdbimpl) Select_stream(requestingUser uint64, page_numb uint64) (schema.PostList, error) {

	var photoList schema.PostList

	rows, err := db.c.Query(`
		SELECT 
			photos.id, 
			photos.authorId, users.username,
			photos.caption, photos.timeOfCreation, photos.likes, photos.comments,
			(SELECT EXISTS(SELECT TRUE FROM likes WHERE photoId = photos.id AND userId = ?)) AS isLiked
		FROM 
			photos INNER JOIN users ON users.id = photos.authorId
		WHERE 
			photos.authorId IN (SELECT followedId FROM follows WHERE followerId = ?)
		ORDER BY photos.timeOfCreation DESC
		LIMIT 24 OFFSET ?`,
		requestingUser, requestingUser, 24*page_numb)

	if err != nil {
		return photoList, err
	}
	defer rows.Close()

	photoList.PageNumber = page_numb
	for rows.Next() {
		var (
			photoId  uint64
			authorId uint64
			username string
			caption  string
			creation int64
			likes    uint64
			comments uint64
			isLiked  bool
		)

		if err := rows.Scan(&photoId, &authorId, &username, &caption, &creation, &likes, &comments, &isLiked); err != nil {
			return photoList, err
		}

		photoList.Posts = append(photoList.Posts, schema.Post{
			Id:             photoId,
			Author:         schema.ReducedUser{Id: authorId, Username: username},
			Caption:        caption,
			TimeOfCreation: creation,
			Likes:          likes,
			Comments:       comments,
			IsLiked:        isLiked,
		})
	}

	if err := rows.Err(); err != nil {
		return photoList, err
	}
	return photoList, err
}
