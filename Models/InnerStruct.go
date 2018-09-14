package Models

import "time"

type AnswerQuery struct  {
	Id          int64     `xorm:"pk autoincr"`
	UserId      int       `xorm:"not null"`
	AnswerTo    int       `xorm:"not null"`
	Text        string    `xorm:"text not null"`
	FileAddress string    `xorm:"varchar(40)"`
	Seen        bool
	Created     time.Time `xorm:"created"`

	Username string `xorm:"varchar(256) not null"`
}

