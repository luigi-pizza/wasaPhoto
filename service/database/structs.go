package database

type SingleUser struct {
	id uint64;
	username string
}

type QueriedUser struct {
	id uint64;
	username string;
	numberOfFollowers uint64;
	accountsFollowed uint64;
	numberOfPosts uint64;
	isBanned bool;
	isFollowed bool	
}

type UsernameFoundError struct{}
func (m *UsernameFoundError) Error() string {
	return "the selected Username was already present in the database"
}
