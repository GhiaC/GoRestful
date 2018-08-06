package Models


type User struct {
	Id       int64
	Username string `xorm:"varchar(256) not null"`
	Password string `xorm:"varchar(256) not null"`
}

func NewUser(username , password string) *User  {
	newUser := new(User)
	newUser.Username = username
	newUser.Password = password
	return newUser
}

type Title struct {
	Id       int64
	Title string `xorm:"varchar(256) not null"`
}

func NewTitle(title string) *Title  {
	newUser := new(Title)
	newUser.Title = title
	return newUser
}

type Subtitle struct {
	Id       int64
	TitleId       int64
	Title string `xorm:"varchar(256) not null"`
}

func NewSubtitle(titleId int64,title string) *Subtitle  {
	newUser := new(Subtitle)
	newUser.TitleId = titleId
	newUser.Title = title
	return newUser
}


type Media struct {
	Id       int64
	SubTitleId       int64
	Text string `xorm:"text not null"`
}

func NewMedia(subtitleid int64,Text string) *Media  {
	newUser := new(Media)
	newUser.SubTitleId = subtitleid
	newUser.Text = Text
	return newUser
}