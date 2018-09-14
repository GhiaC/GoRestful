package Models

type LoginRequest struct {
	Username string
	Password string
}

type SendMessageRequest struct {
	FileAddress string
	Token       string
	Text        string
}
