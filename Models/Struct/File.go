package Struct

import "time"

type UserPicture struct {
	Id       int64     `xorm:"pk autoincr"`
	UserId   int64
	FileName string    `xorm:"varchar(256) unique not null"`
	Key      string    `xorm:"varchar(40) not null"`
	Created  time.Time `xorm:"created"`
}

func NewUserPicture(userId int64, fileName string) *UserPicture {
	newPicture := new(UserPicture)
	newPicture.UserId = userId
	newPicture.FileName = fileName
	newPicture.Created = time.Now()
	return newPicture
}

type AdminPicture struct {
	Id       int64     `xorm:"pk autoincr"`
	AdminId  int64
	FileName string    `xorm:"varchar(256) unique not null"`
	Key      string    `xorm:"varchar(40) not null"`
	Created  time.Time `xorm:"created"`
}

func NewAdminPicture(adminId int64, fileName string) *AdminPicture {
	newPicture := new(AdminPicture)
	newPicture.AdminId = adminId
	newPicture.FileName = fileName
	newPicture.Created = time.Now()
	return newPicture
}

type AdminFile struct {
	Id       int64     `xorm:"pk autoincr"`
	AdminId  int
	FileName string    `xorm:"varchar(256) not null"`
	Key      string    `xorm:"varchar(40) not null"`
	Created  time.Time `xorm:"created"`
}

func NewAdminFile(userId int, fileName string) *AdminFile {
	newFile := new(AdminFile)
	newFile.AdminId = userId
	newFile.FileName = fileName
	newFile.Created = time.Now()
	return newFile
}

