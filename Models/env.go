package Models

type Configuration struct {
	Port   string
	PassDB string
	UserDB string
	DB     string

	//SMS
	AdminPhonenumber string
	UserApiKey       string
	SecretKey        string
	LineNumber       string
}
