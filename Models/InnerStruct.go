package Models

import "time"

type AnswerQuery struct  {
	Id     int64 `xorm:"pk autoincr"`
	UserId int64 `xorm:"not null"`
	Username string `xorm:"varchar(256) not null"`
	Text    string    `xorm:"text not null"`
	Created time.Time `xorm:"created"`
}

