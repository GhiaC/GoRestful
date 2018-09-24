package Models

type LoginRequest struct {
	Username string
	Password string
	IMEI string
}

type SendMessageRequest struct {
	FileAddress string
	Token       string
	Text        string
}

type GetMessageRequest struct {
	Token       string
}
