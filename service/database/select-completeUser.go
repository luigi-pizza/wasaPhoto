package database

import (
	"github.com/luigi-pizza/wasaPhoto/service/components/schema"
)

func (db *appdbimpl) Select_completeUser(requestingUser uint64, requestedUser uint64) (schema.CompleteUser, error) {
	/* get the requestedUser CompleteUser resource, as viewed by requestingUser
	   Optionally, it can check that the searched user had not banned the searching user
	   through the parameter checkBanned */
	// if used with checkBanned == true, Implements GET /users/{userID}

	var result schema.CompleteUser

	var (
		username          string
		numberOfFollowers uint64
		accountsFollowed  uint64
		numberOfPosts     uint64
		isBanned          bool
		isFollowed        bool
	)

	err := db.c.QueryRow(
		`SELECT 
			users.username,
			(SELECT count(*) FROM follows WHERE followedId = ?) AS numberOfFollowers,
			(SELECT count(*) FROM follows WHERE followerId = ?) AS accountsFollowed,
			(SELECT count(*) FROM photos WHERE authorId = ?) AS numberOfPhotos,
			(SELECT EXISTS(SELECT TRUE FROM follows WHERE followerId = ? AND followedId = ?)) AS isFollowed,
			(SELECT EXISTS(SELECT TRUE FROM Bans WHERE bannerId = ? AND bannedId = ?)) AS isBanned
		FROM users
		WHERE users.id = ?`,
		requestedUser, requestedUser, requestedUser, requestedUser,
		requestedUser, requestingUser, requestedUser, requestingUser).Scan(
		&username, &numberOfFollowers, &accountsFollowed, &numberOfPosts, &isFollowed, &isBanned,
	)

	result = schema.CompleteUser{
		Id:                requestedUser,
		Username:          username,
		NumberOfFollowers: numberOfFollowers,
		AccountsFollowed:  accountsFollowed,
		NumberOfPosts:     numberOfPosts,
		IsBanned:          isBanned,
		IsFollowed:        isFollowed,
	}

	return result, err
}
