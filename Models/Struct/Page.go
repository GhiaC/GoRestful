package Struct

import "time"

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
	TitleId int64
	Title   string `xorm:"varchar(256) not null"`
}

func NewSubtitle(titleId int64, title string) *Subtitle {
	newUser := new(Subtitle)
	newUser.TitleId = titleId
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

