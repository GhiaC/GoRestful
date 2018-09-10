package Models

import "time"

type Admin struct {
	Id       int64  `xorm:"pk autoincr"`
	Username string `xorm:"varchar(256) unique not null"`
	Password string `xorm:"varchar(256) not null"`
}

func NewAdmin(username, password string) *Admin {
	newUser := new(Admin)
	newUser.Username = username
	newUser.Password = password
	return newUser
}

type User struct {
	Id       int64  `xorm:"pk autoincr"`
	Username string `xorm:"varchar(256) not null"`
	Password string `xorm:"varchar(256) not null"`
	Token    string `xorm:"varchar(256) default null"`
}

func NewUser(username, password string) *User {
	newUser := new(User)
	newUser.Username = username
	newUser.Password = password
	return newUser
}

type AdminMessage struct {
	Id      int64     `xorm:"pk autoincr"`
	UserId  int64     `xorm:"not null"`
	AdminId int64     `xorm:"not null"`
	Text    string    `xorm:"text not null"`
	Created time.Time `xorm:"created"`
	//TODO date
}

func NewAdminMessage(userId, adminId int64, text string) *AdminMessage {
	newMessage := new(AdminMessage)
	newMessage.UserId = userId
	newMessage.AdminId = adminId
	newMessage.Text = text
	newMessage.Created = time.Now()
	return newMessage
}

type UserMessage struct {
	Id     int64 `xorm:"pk autoincr"`
	UserId int64 `xorm:"not null"`
	//UserId   int64     `xorm:"not null"`
	Text    string    `xorm:"text not null"`
	Created time.Time `xorm:"created"`
}

func NewUserMessage(userId int64, text string) *UserMessage {
	newMessage := new(UserMessage)
	newMessage.UserId = userId
	newMessage.Text = text
	newMessage.Created = time.Now()
	return newMessage
}

type Title struct {
	Id    int64  `xorm:"pk autoincr"`
	Title string `xorm:"varchar(256) unique not null"`
}

func NewTitle(title string) *Title {
	newUser := new(Title)
	newUser.Title = title
	return newUser
}

type Subtitle struct {
	Id      int64  `xorm:"pk autoincr"`
	TitleId int64  `xorm:"index"`
	Title   string `xorm:"varchar(256) unique not null"`
}

func NewSubtitle(titleId int64, title string) *Subtitle {
	newUser := new(Subtitle)
	newUser.Titleid = titleId
	newUser.Title = title
	return newUser
}

type Media struct {
	Id         int64  `xorm:"pk autoincr"`
	Subtitleid int64  `xorm:"index"`
	Text       string `xorm:"text not null"`
}

func NewMedia(subtitleid int64, Text string) *Media {
	newUser := new(Media)
	newUser.Subtitleid = subtitleid
	newUser.Text = Text
	return newUser
}

type Picture struct {
	Id       int64     `xorm:"pk autoincr"`
	UserId   int64     `xorm:"index"`
	FileName string    `xorm:"varchar(256) unique not null"`
	Created  time.Time `xorm:"created"`
}

func NewPicture(userId int64, fileName string) *Picture {
	newPicture := new(Picture)
	newPicture.UserId = userId
	newPicture.FileName = fileName
	newPicture.Created = time.Now()
	return newPicture
}

type AdminFile struct {
	Id       int64     `xorm:"pk autoincr"`
	AdminId  int64     `xorm:"index"`
	FileName string    `xorm:"varchar(256) unique not null"`
	Created  time.Time `xorm:"created"`
}

func NewAdminFile(userId int64, fileName string) *Picture {
	newFile := new(Picture)
	newFile.UserId = userId
	newFile.FileName = fileName
	newFile.Created = time.Now()
	return newFile
}

type AboutUs struct {
	Id      int64     `xorm:"pk autoincr"`
	Text    string    `xorm:"text unique not null"`
	Created time.Time `xorm:"created"`
}

func NewAboutUs(Text string) *AboutUs {
	newAboutUs := new(AboutUs)
	newAboutUs.Text = Text
	newAboutUs.Created = time.Now()
	return newAboutUs
}

type News struct {
	Id       int64     `xorm:"pk autoincr"`
	Text     string    `xorm:"text unique not null"`
	FileName string    `xorm:"varchar(256) unique not null"`
	Created  time.Time `xorm:"created"`
}

func NewNews(Text, FileName string) *News {
	news := new(News)
	news.Text = Text
	news.FileName = FileName
	news.Created = time.Now()
	return news
}
