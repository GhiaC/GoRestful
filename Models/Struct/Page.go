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
	Pic1    string `xorm:"varchar(256) not null"`
	Pic2    string `xorm:"varchar(256) not null"`
}

func NewSubtitle(titleId int64, title, pic1, pic2 string) *Subtitle {
	newUser := new(Subtitle)
	newUser.TitleId = titleId
	newUser.Pic1 = pic1
	newUser.Pic2 = pic2
	newUser.TitleId = titleId
	newUser.Title = title
	return newUser
}

type Media struct {
	Id         int64  `xorm:"pk autoincr"`
	Subtitleid int64  `xorm:"index"`
	Picture    string `xorm:"varchar(256) not null"`
	Text       string `xorm:"text not null"`
}

func NewMedia(subtitleid int64, Text,Picture string) *Media {
	newUser := new(Media)
	newUser.Subtitleid = subtitleid
	newUser.Picture = Picture
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
