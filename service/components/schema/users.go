package schema

type ReducedUser struct {
	Id       uint64 `json:"userId"`
	Username string `json:"username"`
}

type UserList struct {
	Users []ReducedUser `json:"users"`
}

type CompleteUser struct {
	Id                uint64 `json:"userId"`
	Username          string `json:"username"`
	NumberOfFollowers uint64 `json:"numberOfFollowers"`
	AccountsFollowed  uint64 `json:"accountsFollowed"`
	NumberOfPosts     uint64 `json:"numberOfPosts"`
	IsBanned          bool   `json:"isBanned"`
	IsFollowed        bool   `json:"isFollowed"`
}
