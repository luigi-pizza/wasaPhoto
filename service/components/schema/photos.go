package schema

type Post struct {
	Id             uint64      `json:"photoId"`
	Author         ReducedUser `json:"author"`
	Caption        string      `json:"caption"`
	Likes          uint64      `json:"numberOfLikes"`
	Comments       uint64      `json:"numberOfComments"`
	TimeOfCreation int64       `json:"creation"`
	IsLiked        bool        `json:"isliked"`
}

type PostList struct {
	PageNumber uint64 `json:"page"`
	Posts      []Post `json:"posts"`
}
