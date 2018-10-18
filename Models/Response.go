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
	url      string
	Error    string
	FileName string
}

type UResponse struct {
	Url      string
}

func NewUploadFileResponse(url string) *UResponse {
	newFile := new(UResponse)
	newFile.Url = url
	return newFile
}
