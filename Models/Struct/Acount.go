package Struct


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


