package Models

type PageVariables struct {
	Answer string
}

type LoginPageVariables struct {
	Answer      string
	Url         string
	SubmitValue string
}

type HomePageVariables struct {
	Date        string
	Time        string
	LoginStatus string
}

type StatusPageVariables struct {
	Users [] User
}

type HeaderVariables struct {
	Title string
}

type NavigationVariables struct {
	LoggedIn bool
}

type FirstLayerVariables struct {
	Titles [] Title

	Answer      string
	Url         string
	SubmitValue string
}

type SecondLayerVariables struct {

	Subtitles [] Subtitle
	TitleId string

	Answer      string
	Url         string
	SubmitValue string
}
