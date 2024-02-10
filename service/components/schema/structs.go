package schema

type ReducedUser struct {
	id       uint64 `json:"userId"`
	username string `json:"username"`
}

type CompleteUser struct {
	id                uint64 `json:"userId"`
	username          string `json:"username"`
	numberOfFollowers uint64 `json:"numebrOfFollowers"`
	accountsFollowed  uint64 `json:"accountsFollowed"`
	numberOfPosts     uint64 `json:"numberOfPosts"`
	isBanned          bool   `json:"isBanned"`
	isFollowed        bool   `json:"isFollowed"`
}

type Comment struct {
	id             uint64      `json:"commentId"`
	authorId       ReducedUser `json:"author"`
	photoId        uint64      `json:"photoId"`
	commentText    string      `json:"text"`
	timeOfCreation uint64      `json:"creation"`
}

type Post struct {
	id             uint64      `json:"photoId"`
	author         ReducedUser `json:"author"`
	caption        string      `json:"caption"`
	likes          uint64      `json:numberOfLikes""`
	comments       uint64      `json:"numberOfComments"`
	timeOfCreation uint64      `json:"creation"`
	isLiked        bool        `json:"isliked"`
}

type UsernameFoundError struct{}

func (m *UsernameFoundError) Error() string {
	return "the requested Username was already present in the database"
}

type BannedRequestingUserError struct{}

func (m *BannedRequestingUserError) Error() string {
	return "the requested user has banned the requesting user"
}
