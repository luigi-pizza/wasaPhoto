package database

const (
	usersTableCreationStatement = ` 
		CREATE TABLE IF NOT EXISTS users (
			id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username        TEXT NOT NULL UNIQUE
		);
	`

	photosTableCreationStatement = `
		CREATE TABLE IF NOT EXISTS photos (
			id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			authorId        INTEGER NOT NULL,
			caption			TEXT,
			timeOfCreation  INTEGER NOT NULL,
			likes			INTEGER,
			comments		INTEGER,
		
			FOREIGN KEY (authorId) REFERENCES users(id) ON DELETE CASCADE
		);
	`

	commentsTableCreationStatement = `
		CREATE TABLE IF NOT EXISTS comments (
			id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			authorId        INTEGER NOT NULL,
			photoId         INTEGER NOT NULL,
			commentText     TEXT NOT NULL,
			timeOfCreation  INTEGER NOT NULL,
			
			FOREIGN KEY (authorId) REFERENCES users (id) ON DELETE CASCADE,
			FOREIGN KEY (photoId)  REFERENCES photos(id) ON DELETE CASCADE
		);
	`

	likesTableCreationStatement = `
		CREATE TABLE IF NOT EXISTS likes (
			userId  INTEGER NOT NULL, 
			photoId INTEGER NOT NULL, 
		
			PRIMARY KEY (userId, photoId),
			FOREIGN KEY (userId)  REFERENCES users (id) ON DELETE CASCADE,
			FOREIGN KEY (photoId) REFERENCES photos(id) ON DELETE CASCADE 
		);
	
	`

	followsTableCreationStatement = `
		CREATE TABLE IF NOT EXISTS follows (
			followerId INTEGER NOT NULL,
			followedId INTEGER NOT NULL,
		
			PRIMARY KEY (followerId, followedId)
			FOREIGN KEY (followerId) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (followedId) REFERENCES users(id) ON DELETE CASCADE
		);
	`

	bansTableCreationStatement = `
		CREATE TABLE IF NOT EXISTS bans (
			bannerId INTEGER NOT NULL,
			bannedId INTEGER NOT NULL,
			
			PRIMARY KEY (bannerId, bannedId),
			FOREIGN KEY (bannerId) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (bannedId) REFERENCES users(id) ON DELETE CASCADE
		);
	`
)
