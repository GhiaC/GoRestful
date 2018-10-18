package Struct

import (
	"html/template"
	"time"
)

type Title struct {
	Id      int64  `xorm:"pk autoincr"`
	Title   string `xorm:"varchar(256) unique not null"`
	Picture string `xorm:"varchar(256) not null"`
}

func NewTitle(title, picture string) *Title {
	newUser := new(Title)
	newUser.Title = title
	newUser.Picture = picture
	return newUser
}

type Subtitle struct {
	//changed
	Id      int64 `xorm:"pk autoincr"`
	Pid     int64
	Title   string `xorm:"varchar(256) not null"`
	Picture string `xorm:"varchar(256) not null"`
}

func NewSubtitle(pid int64, title, picture string) *Subtitle {
	newUser := new(Subtitle)
	newUser.Pid = pid
	newUser.Picture = picture
	newUser.Title = title
	return newUser
}

type Media struct {
	//Changed
	Id      int64 `xorm:"pk autoincr"`
	Pid     int64
	Picture string `xorm:"varchar(256) not null"`
	Title   string `xorm:"text not null"`
}

func NewMedia(pid int64, Title, picture string) *Media { //changed
	newUser := new(Media)
	newUser.Pid = pid
	newUser.Picture = picture
	newUser.Title = Title
	return newUser
}

type SubMedia struct {
	//Changed
	Id   int64 `xorm:"pk autoincr"`
	Pid  int64
	Url  string `xorm:"varchar(256) not null"`
	Text string `xorm:"text not null"`
}

func NewSubMedia(pid int64, picture, Text string) *SubMedia { //changed
	newUser := new(SubMedia)
	newUser.Pid = pid
	newUser.Url = picture
	newUser.Text = Text
	return newUser
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
	Id       int64         `xorm:"pk autoincr"`
	Title    string        `xorm:"text unique not null"`
	Text     template.HTML `xorm:"text unique not null"`
	FileName string        `xorm:"varchar(256) unique not null"`
	Created  time.Time     `xorm:"created"`
}

func NewNews(Text template.HTML, FileName, Title string) *News {
	news := new(News)
	news.Text = Text
	news.FileName = FileName
	news.Title = Title
	news.Created = time.Now()
	return news
}
