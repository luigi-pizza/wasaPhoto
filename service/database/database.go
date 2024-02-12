/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/luigi-pizza/wasaPhoto/service/components/schema"
)

// §§ Definisci Interfaccia

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// Assert
	IsBanned(bannerId uint64, bannedId uint64) (bool, error)
	IsCommentId(comment_id uint64) (bool, uint64, uint64, error)
	IsFollowed(followerId uint64, followedId uint64) (bool, error)
	IsLiked(user_id uint64, photo_id uint64) (bool, error)
	IsPhotoId(photo_id uint64) (bool, uint64, error)
	IsUserId(user_id uint64) (bool, error)

	// Delete
	Delete_ban(pardonerId uint64, pardonedId uint64) error
	Delete_comment(commentId uint64, photoId uint64) error
	Delete_follow(followerId uint64, followedId uint64) error
	Delete_like(user_id uint64, photo_id uint64) error
	Delete_photo(id uint64) error
	// NO delete_user

	// Insert
	Insert_ban(bannerId uint64, bannedId uint64) error
	Insert_comment(photoId uint64, authorId uint64, text string, timeOfCreation int64) (uint64, error)
	Insert_follow(followerId uint64, followedId uint64) error
	Insert_like(user_id uint64, photo_id uint64) error
	Insert_photo(authorId uint64, caption string, creation int64) (uint64, error)
	Insert_user(username string) (uint64, bool, error)

	// select
	Select_commentList(user_id uint64, post_id uint64, page_numb uint64) (schema.CommentList, error)
	Select_completeUser(requestingUser uint64, requestedUser uint64) (schema.CompleteUser, error)
	Select_postList(requestingUser uint64, requestedUser uint64, page_numb uint64) (schema.PostList, error)
	Select_reducedUser(userId uint64) (schema.ReducedUser, error)
	Select_stream(requestingUser uint64, page_numb uint64) (schema.PostList, error)
	Select_userList(requestingUser uint64, prompt string) (schema.UserList, error)

	// update
	Update_username(userId uint64, newUsername string) error

	Ping() error
}

// §§ Implementa interfaccia in file singoli

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {

	// Check Existence of passed Database
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	TableMapping := map[string]string{
		"bans":     bansTableCreationStatement,
		"users":    usersTableCreationStatement,
		"likes":    likesTableCreationStatement,
		"photos":   photosTableCreationStatement,
		"follows":  followsTableCreationStatement,
		"comments": commentsTableCreationStatement,
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	for tableName, sqlStmt := range TableMapping {
		err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name= ? ;`, tableName).Scan(&tableName)

		if errors.Is(err, sql.ErrNoRows) {
			_, err = db.Exec(sqlStmt)

			if err != nil {
				return nil, fmt.Errorf("error creating database structure.\n%s -> %w", tableName, err)
			}
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
