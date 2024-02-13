package schema

type Comment struct {
	Id             uint64      `json:"commentId"`
	Author         ReducedUser `json:"author"`
	PhotoId        uint64      `json:"photoId"`
	CommentText    string      `json:"text"`
	TimeOfCreation int64       `json:"creation"`
}

type CommentList struct {
	PageNumber uint64    `json:"page"`
	Comments   []Comment `json:"comments"`
}
