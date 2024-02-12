package requests

type Username struct {
	Username string `json:"username"`
}

func (request *Username) IsValid() bool {
	return check_username(request.Username)
}
