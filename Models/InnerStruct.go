package Models

import "time"

type AnswerQuery struct {
	Id          int64     `xorm:"pk autoincr"`
	UserId      int       `xorm:"not null"`
	AnswerTo    int       `xorm:"not null"`
	Text        string    `xorm:"text not null"`
	FileAddress string    `xorm:"varchar(40)"`
	Seen        bool
	Created     time.Time `xorm:"created"`

	Username string `xorm:"varchar(256) not null"`
	Type     string `xorm:"varchar(40)"`
}

type MediaJoinFile struct {
	Id      int64  `xorm:"pk autoincr"`
	Pid     int64
	Picture string `xorm:"varchar(256) not null"`
	Title   string `xorm:"text not null"`

	Type string `xorm:"varchar(40)"`
}

type SubMediaJoinFile struct {
	//Changed
	Id   int64  `xorm:"pk autoincr"`
	Pid  int64
	Url  string `xorm:"varchar(256) not null"`
	Text string `xorm:"text not null"`

	Type string `xorm:"varchar(40)"`
}

type NewsJoinFile struct {
	Id       int64     `xorm:"pk autoincr"`
	Title    string    `xorm:"text unique not null"`
	Text     string    `xorm:"text unique not null"`
	FileName string    `xorm:"varchar(256) unique not null"`
	Created  time.Time `xorm:"created"`

	Type string `xorm:"varchar(40)"`
}

type SubtitleJoinFile struct {
	//changed
	Id      int64  `xorm:"pk autoincr"`
	Pid     int64
	Title   string `xorm:"varchar(256) not null"`
	Picture string `xorm:"varchar(256) not null"`

	Type string `xorm:"varchar(40)"`
}

type TitleJoinFile struct {
	Id      int64  `xorm:"pk autoincr"`
	Title   string `xorm:"varchar(256) unique not null"`
	Picture string `xorm:"varchar(256) not null"`

	Type string `xorm:"varchar(40)"`
}
