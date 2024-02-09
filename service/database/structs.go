package database

type ReducedUser struct {
	id uint64;
	username string
}

type CompleteUser struct {
	id uint64;
	username string;
	numberOfFollowers uint64;
	accountsFollowed uint64;
	numberOfPosts uint64;
	isBanned bool;
	isFollowed bool	
}

type ReducedComment struct {
	id uint64;
	authorId uint64;
	photoId uint64;
	commentText string ;
	timeOfCreation uint64
}

type CompleteComment struct {
	ReducedComment;
	user CompleteUser
}

type UsernameFoundError struct{}
func (m *UsernameFoundError) Error() string {
	return "the requested Username was already present in the database"
}

type BannedRequestingUserError struct{}
func (m *BannedRequestingUserError) Error() string {
	return "the requested user has banned the requesting user"
}
