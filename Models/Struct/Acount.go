package Struct

type User struct {
	Id       int    `xorm:"pk autoincr"`
	Username string `xorm:"varchar(256) unique not null"`
	Password string `xorm:"varchar(256) not null"`
	Type     int // 0 root admin // 1 admin // 2 user
	Token    string `xorm:"varchar(256) default null"`
}

func NewUser(username, password string, Type int) *User {
	newUser := new(User)
	newUser.Username = username
	newUser.Password = password
	newUser.Type = Type
	return newUser
}

//type User struct {
//	Id       int64  `xorm:"pk autoincr"`
//	Username string `xorm:"varchar(256) not null"`
//	Password string `xorm:"varchar(256) not null"`
//	Token    string `xorm:"varchar(256) default null"`
//}
//
//func NewUser(username, password string) *User {
//	newUser := new(User)
//	newUser.Username = username
//	newUser.Password = password
//	return newUser
//}
//
//
