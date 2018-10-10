package Struct

type SendSMSMessageRequest struct {
	Messages                 []string
	MobileNumbers            []string
	LineNumber               string
	SendDateTime             string
	CanContinueInCaseOfError string
}

type GetSMSTokenRequest struct {
	UserApiKey string
	SecretKey  string
}

type GetTokenResponse struct {
	TokenKey     string
	IsSuccessful string
	Message      string
}