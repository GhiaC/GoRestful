package Struct

type User struct {
	Id          int    `xorm:"pk autoincr"`
	Username    string `xorm:"varchar(256) unique not null"`
	Password    string `xorm:"varchar(256) not null"`
	Name        string `xorm:"varchar(256)"`
	PhoneNumber string `xorm:"varchar(20)"`
	Type        int // 0 root admin // 1 admin // 2 user
	Imei        string `xorm:"varchar(40)"`
	Token       string `xorm:"varchar(256) default null"`
	Status      int
}

func NewUser(username, password, name, phonenumber, IMEI string, Type, status int) *User {
	newUser := new(User)
	newUser.Username = username
	newUser.Password = password
	newUser.Name = name
	newUser.PhoneNumber = phonenumber
	newUser.Imei = IMEI
	newUser.Type = Type
	newUser.Status = status
	return newUser
}
