package Models

type LoginResponse struct {
	Result      bool
	Token       string
	Name        string
	PhoneNumber string
	Error       string
}
type SendMessageResponse struct {
	Result    bool
	Error     string
	MessageId int64
}
type GetMessageResponse struct {
	Result   bool
	Error    string
	Messages []AnswerQuery
}

type UploadFileResponse struct {
	Result   bool
	Error    string
	FileName string
}
