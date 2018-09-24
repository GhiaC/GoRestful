package Struct

import "time"

//type Picture struct {
//	Id          int64     `xorm:"pk autoincr"`
//	UserId      int
//	Description string    `xorm:"text not null"`
//	FileName    string    `xorm:"varchar(256) unique not null"`
//	Key         string    `xorm:"varchar(40) unique not null"`
//	Type        string    `xorm:"varchar(40) unique not null"`
//	Created     time.Time `xorm:"created"`
//}

//func NewPicture(userId int, fileName, key, description string) *Picture {
//	newPicture := new(Picture)
//	newPicture.UserId = userId
//	newPicture.FileName = fileName
//	newPicture.Description = description
//	newPicture.Key = key
//	newPicture.Created = time.Now()
//	return newPicture
//}

type File struct {
	Id          int64     `xorm:"pk autoincr"`
	UserId      int
	Description string    `xorm:"text not null"`
	FileName    string    `xorm:"varchar(256) not null"`
	Key         string    `xorm:"varchar(40) not null"`
	Type        string    `xorm:"varchar(40) unique not null"`
	AdminFile      bool
	Created     time.Time `xorm:"created"`
}

func NewFile(userId int, fileName, key, description, Type string,adminFile bool) *File {
	newFile := new(File)
	newFile.UserId = userId
	newFile.FileName = fileName
	newFile.Description = description
	newFile.AdminFile = adminFile
	newFile.Key = key
	newFile.Type = Type
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
