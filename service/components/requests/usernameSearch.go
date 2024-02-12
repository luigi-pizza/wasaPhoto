package requests

type UsernameSearch struct {
	Username string `json:"username"`
}

func (request *UsernameSearch) IsValid() bool {
	return check_usernameSearch(request.Username)
}
