package database

import "github.com/luigi-pizza/wasaPhoto/service/components/schema"

func (db *appdbimpl) Select_reducedUser(userId uint64) (schema.ReducedUser, error) {
	// ReducedUser_info_by_id returns the ReducedUser with the given userId
	var user schema.ReducedUser
	err := db.c.QueryRow("SELECT id, username FROM users WHERE id = ?", userId).Scan(&user.Id, &user.Username)
	return user, err
}
