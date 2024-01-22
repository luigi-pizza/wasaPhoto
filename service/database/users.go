package database

import (
	"database/sql"
)

// UserInfo returns all general info about a user
func (db *appdbimpl) singleUserInfo(userId uint64) (SingleUser, error) {
	var user SingleUser
	err := db.c.QueryRow("SELECT * FROM users WHERE id=?", userId).Scan(&user)
	return user, err
}

func (db *appdbimpl) queriedUserInfo(requestingUser uint64, queriedUser uint64) (QueriedUser, error) {
	var result QueriedUser
	err := db.c.QueryRow(`
		Select 
			*
			U.id, U.username,
			nof.numberOfFollowers, af.accountsFollowed,
			nop.numberOfPosts,
			CASE WHEN EXISTS B THEN 1 ELSE 0 END AS isBanned,
			CASE WHEN EXISTS F THEN 1 ELSE 0 END AS isFollowed
		FROM
			(select * from users where id= ? ) as u,
			(select count(*) as numberOfFollowers from follows where followedId= ? ) as nof,
			(select count(*) as accountsFollowed  from follows where followerId= ? ) as af,
			(select count(*) as numberOfPosts from photos where authorId= ? ) as nop,
			(select * from bans where bannedId= ? and bannerId= ? ) as B,
			(select * from follows where followedId= ? and followerId= ? ) as F;
		`, queriedUser, queriedUser, queriedUser, queriedUser, 
		   queriedUser, requestingUser, queriedUser, requestingUser).Scan(&result)
	return result, err
}

func (db *appdbimpl) usernameLike(prompt string) ([]SingleUser, error) {
	var result []SingleUser
	rows, err := db.c.Query(`
		SELECT * FROM users
		WHERE username LIKE '?%'
		LIMIT 24;`, prompt)
	if err != nil {return nil, err}
	
	defer rows.Close()
	for rows.Next() {
		var user SingleUser

		if err := rows.Scan(&user.id, &user.username); err != nil {return nil, err}
		result = append(result, user)
	}
	return result, err
}

func (db *appdbimpl) modifyUsername(userId uint64, newUsername string) (error) {

	var user SingleUser
	err := db.c.QueryRow("SELECT * FROM users WHERE username=?", newUsername).Scan(&user)

	if err != sql.ErrNoRows { return &UsernameFoundError{}} else if err != nil { return err } // check errors

	_, err = db.c.Exec(`UPDATE users SET username=? WHERE id = ?;`, newUsername, userId)

	return err
}




