package Models

type LoginRequest struct {
	Username string
	Password string
}

type SendMessageRequest struct {
	Token string
	Text  string
}
