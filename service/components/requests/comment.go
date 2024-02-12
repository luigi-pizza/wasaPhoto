package requests

type Text struct {
	Text string `json:"text"`
}

func (request *Text) IsValid() bool {
	return check_commentText(request.Text)
}
