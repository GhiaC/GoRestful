package Struct

import "time"

type Picture struct {
	Id       int64     `xorm:"pk autoincr"`
	UserId   int
	FileName string    `xorm:"varchar(256) unique not null"`
	Key      string    `xorm:"varchar(40) unique not null"`
	Created  time.Time `xorm:"created"`
}

func NewPicture(userId int, fileName, key string) *Picture {
	newPicture := new(Picture)
	newPicture.UserId = userId
	newPicture.FileName = fileName
	newPicture.Key = key
	newPicture.Created = time.Now()
	return newPicture
}


type File struct {
	Id       int64     `xorm:"pk autoincr"`
	UserId  int
	FileName string    `xorm:"varchar(256) not null"`
	Key      string    `xorm:"varchar(40) not null"`
	Created  time.Time `xorm:"created"`
}

func NewFile(userId int, fileName, key string) *File {
	newFile := new(File)
	newFile.UserId = userId
	newFile.FileName = fileName
	newFile.Key = key
	newFile.Created = time.Now()
	return newFile
}


//type AdminPicture struct {
//	Id       int64     `xorm:"pk autoincr"`
//	AdminId  int64
//	FileName string    `xorm:"varchar(256) unique not null"`
//	Key      string    `xorm:"varchar(40) not null"`
//	Created  time.Time `xorm:"created"`
//}
//
//func NewAdminPicture(adminId int64, fileName, key string) *AdminPicture {
//	newPicture := new(AdminPicture)
//	newPicture.AdminId = adminId
//	newPicture.FileName = fileName
//	newPicture.Key = key
//	newPicture.Created = time.Now()
//	return newPicture
//}
