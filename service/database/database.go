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
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var users string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&users)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS users (
			id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username        TEXT NOT NULL UNIQUE,
			name            TEXT NOT NULL,
			surname         TEXT NOT NULL
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var photos string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&photos)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS photos (
			id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			authorId        INTEGER NOT NULL,
			photoUrl        TEXT NOT NULL,
			timeOfCreation  INTEGER NOT NULL,
		
			FOREIGN KEY (authorId) REFERENCES User (id) ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var comments string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&comments)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS comments (
			id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			string          TEXT NOT NULL,
			timeOfCreation  INTEGER NOT NULL,
			authorId        INTEGER NOT NULL,
			photoId         INTEGER NOT NULL,
			
			FOREIGN KEY (authorId) REFERENCES User  (id) ON DELETE CASCADE,
			FOREIGN KEY (photoId)  REFERENCES Photo (id) ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var likes string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&likes)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS likes (
			userId  INTEGER NOT NULL, 
			photoId INTEGER NOT NULL, 
		
			PRIMARY KEY (userId, photoId),
			FOREIGN KEY (userId)  REFERENCES User (id)  ON DELETE CASCADE,
			FOREIGN KEY (photoId) REFERENCES Photo (id) ON DELETE CASCADE 
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var follows string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='follows';`).Scan(&follows)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS follows (
			followerId INTEGER NOT NULL,
			followedId INTEGER NOT NULL,
		
			PRIMARY KEY (followerId, followedId)
			FOREIGN KEY (followerId) REFERENCES User (id) ON DELETE CASCADE,
			FOREIGN KEY (followedId) REFERENCES User (id) ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var bans string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='bans';`).Scan(&bans)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS bans (
			bannerId INTEGER NOT NULL,
			bannedId INTEGER NOT NULL,
			
			PRIMARY KEY (bannerId, bannedId),
			FOREIGN KEY (bannerId) REFERENCES User (id) ON DELETE CASCADE,
			FOREIGN KEY (bannedId) REFERENCES User (id) ON DELETE CASCADE
		
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
