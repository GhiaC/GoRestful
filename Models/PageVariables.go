package Models

import "GoRestful/Models/Struct"

type PageVariables struct {
	Answer string
}

type LoginPageVariables struct {
	Answer string
	//Url         string
	SubmitValue string
}

type HomePageVariables struct {
	Date        string
	Time        string
	LoginStatus string
}

type StatusPageVariables struct {
	Users [] Struct.User
}
type UsersPageVariables struct {
	Users [] Struct.User
}

type HeaderVariables struct {
	Title string
}

type NavigationVariables struct {
	LoggedIn bool
}

type FirstLayerVariables struct {
	Titles [] Struct.Title

	Answer string
	//Url         string
	SubmitValue string
}

type SecondLayerVariables struct {
	Subtitles [] Struct.Subtitle
	TitleId   string

	Answer string
	//Url         string
	SubmitValue string
}

type MediaLayerVariables struct {
	Medias  [] Struct.Media
	TitleId string

	Answer string
	//Url         string
	SubmitValue string
}
type AnswerLayerVariables struct {
	Msg     AnswerQuery
	Answers []AnswerQuery
	TitleId string

	Answer      string
	SubmitValue string
}

type NewsLayerVariables struct {
	News [] Struct.News
	//TitleId string

	Answer      string
	Url         string
	SubmitValue string
}

type AdminFileLayerVariables struct {
	Files [] Struct.File
}

type MessagesLayerVariables struct {
	Messages [] AnswerQuery
}
