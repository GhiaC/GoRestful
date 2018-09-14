package Models

import "GoRestful/Models/Struct"

type LoginResponse struct {
	Result bool
	Token  string
	Error  string
}
type SendMessageResponse struct {
	Result    bool
	Error     string
	MessageId int64
}
type GetMessageResponse struct {
	Result        bool
	Error         string
	Messages  []Struct.Message
}

type UploadFileResponse struct {
	Result   bool
	Error    string
	FileName string
}
