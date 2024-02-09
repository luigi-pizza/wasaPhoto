package database

func (db *appdbimpl) ReducedUser_info_by_id(userId uint64) (ReducedUser, error) {
	// ReducedUser_info_by_id returns the ReducedUser with the given userId
	var user ReducedUser
	err := db.c.QueryRow("SELECT * FROM users WHERE id=?", userId).Scan(&user)
	return user, err
}


func (db *appdbimpl) ReducedUser_info_by_username(username string) (ReducedUser, error) {
	// ReducedUser_info_by_id returns the ReducedUser with the given username
	var user ReducedUser
	err := db.c.QueryRow("SELECT * FROM users WHERE username=?", username).Scan(&user)
	return user, err
}

func (db *appdbimpl) ReducedUser_usernameLike(requestingUserId uint64, prompt string) ([]ReducedUser, error) {
	/* returns at most 24 ReducedUsers records with the requested prompt 
	in their username among those who have not banned the requesteduser */

	var result []ReducedUser

	rows, err := db.c.Query(`
		SELECT * FROM users
		WHERE username LIKE '%' || ? || '%' AND
		NOT EXISTS (SELECT 1 FROM bans WHERE 
			bannerId = users.userId AND bannedId = ?)
		LIMIT 24;`, prompt, requestingUserId)
	if (err != nil) {return nil, err}
	
	defer rows.Close()
	for rows.Next() {
		var user ReducedUser

		if err := rows.Scan(&user.id, &user.username); err != nil {return nil, err}
		result = append(result, user)
	}
	return result, err
}