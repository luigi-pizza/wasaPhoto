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
)

const (
	usersTableCreationStatement = ` 
	CREATE TABLE IF NOT EXISTS users (
		id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username        TEXT NOT NULL UNIQUE
	);`
	
	photosTableCreationStatement = `
	CREATE TABLE IF NOT EXISTS photos (
		id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		authorId        INTEGER NOT NULL,
		caption			TEXT,
		photoUrl        TEXT NOT NULL,
		timeOfCreation  INTEGER NOT NULL,
	
		FOREIGN KEY authorId REFERENCESid User (id) ON DELETE CASCADE
	);`
	
	commentsTableCreationStatement = `
	CREATE TABLE IF NOT EXISTS comments (
		id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		authorId        INTEGER NOT NULL,
		photoId         INTEGER NOT NULL,
		commentText     TEXT NOT NULL,
		timeOfCreation  INTEGER NOT NULL,
		
		FOREIGN KEY authorId REFERENCES User  (id) ON DELETE CASCADE,
		FOREIGN KEY photoId  REFERENCES Photo (id) ON DELETE CASCADE
	);`
	
	likesTableCreationStatement = `
	CREATE TABLE IF NOT EXISTS likes (
		userId  INTEGER NOT NULL, 
		photoId INTEGER NOT NULL, 
	
		PRIMARY KEY (userId, photoId),
		FOREIGN KEY userId  REFERENCES User (id)  ON DELETE CASCADE,
		FOREIGN KEY photoId REFERENCES Photo (id) ON DELETE CASCADE 
	);`
	
	followsTableCreationStatement = `
	CREATE TABLE IF NOT EXISTS follows (
		followerId INTEGER NOT NULL,
		followedId INTEGER NOT NULL,
	
		PRIMARY KEY (followerId, followedId)
		FOREIGN KEY followerId REFERENCES User (id) ON DELETE CASCADE,
		FOREIGN KEY followedId REFERENCES User (id) ON DELETE CASCADE
	);`
	
	bansTableCreationStatement = `
	CREATE TABLE IF NOT EXISTS bans (
		bannerId INTEGER NOT NULL,
		bannedId INTEGER NOT NULL,
		
		PRIMARY KEY (bannerId, bannedId),
		FOREIGN KEY bannerId REFERENCES User (id) ON DELETE CASCADE,
		FOREIGN KEY bannedId REFERENCES User (id) ON DELETE CASCADE
	);`
)

// §§ Definisci Interfaccia

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// GetName() (string, error)
	// SetName(name string) error
	// User()
	// InsertLike(postid string) error

	// Ping() error
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
		"bans": 	bansTableCreationStatement,
		"users": 	usersTableCreationStatement,
		"likes": 	likesTableCreationStatement,
		"photos": 	photosTableCreationStatement,
		"follows": 	followsTableCreationStatement,
		"comments": commentsTableCreationStatement,
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	for tableName, sqlStmt := range TableMapping {
		err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name= ? ;`, tableName).Scan(&tableName)
		
		if errors.Is(err, sql.ErrNoRows) {
			_, err = db.Exec(sqlStmt)

			if (err != nil) {return nil, fmt.Errorf("error creating database structure.\n%s -> %w", tableName, err)}
		}
	}
	

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
