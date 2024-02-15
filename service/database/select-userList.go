package database

import "github.com/luigi-pizza/wasaPhoto/service/components/schema"

func (db *appdbimpl) Select_userList(requestingUser uint64, prompt string) (schema.UserList, error) {
	/* returns at most 24 ReducedUsers records with the requested prompt
	in their username among those who have not banned the requesteduser */
	// Implements GET /users/

	var result schema.UserList

	rows, err := db.c.Query(`
		SELECT * FROM users
		WHERE username LIKE '%' || ? || '%' AND
		NOT EXISTS (SELECT 1 FROM bans WHERE 
			bannerId = users.id AND bannedId = ?)
		ORDER BY LENGTH(username) ASC
		LIMIT 24;`, prompt, requestingUser)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var user schema.ReducedUser
		if err := rows.Scan(&user.Id, &user.Username); err != nil {
			return result, err
		}
		result.Users = append(result.Users, user)
	}

	if err := rows.Err(); err != nil {
		return result, err
	}
	return result, err
}
