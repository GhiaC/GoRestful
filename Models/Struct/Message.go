package Struct

import "time"

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

