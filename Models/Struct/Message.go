package Struct

import "time"

type Message struct {
	Id          int64     `xorm:"pk autoincr"`
	UserId      int       `xorm:"not null"`
	AnswerTo    int       `xorm:"not null"`
	Text        string    `xorm:"text not null"`
	FileAddress string    `xorm:"varchar(40)"`
	Seen        bool
	Created     time.Time `xorm:"created"`
}

func NewMessage(userId, AnswerTo int, text, fileAddress string) *Message {
	newMessage := new(Message)
	newMessage.UserId = userId
	newMessage.AnswerTo = AnswerTo
	newMessage.FileAddress = fileAddress
	newMessage.Seen = false
	newMessage.Text = text
	newMessage.Created = time.Now()
	return newMessage
}

//type UserMessage struct {
//	Id     int64 `xorm:"pk autoincr"`
//	UserId int64 `xorm:"not null"`
//	//UserId   int64     `xorm:"not null"`
//	Text    string    `xorm:"text not null"`
//	Created time.Time `xorm:"created"`
//}
//
//func NewUserMessage(userId int64, text string) *UserMessage {
//	newMessage := new(UserMessage)
//	newMessage.UserId = userId
//	newMessage.Text = text
//	newMessage.Created = time.Now()
//	return newMessage
//}
