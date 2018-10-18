package Controler

import (
	"github.com/go-xorm/xorm"
	"fmt"
	//"GoRestful/*"
	//"log"
	//"log/syslog"
	"../Models/Struct"
)

var engine *xorm.Engine
var flag bool

func GetEngine() *xorm.Engine{
	if !flag {
		flag = true
		var errDB error
		engine, errDB = xorm.NewEngine("mysql", Configuration.UserDB+":"+Configuration.PassDB+"@/"+
			Configuration.DB+"?charset=utf8")
		if errDB != nil {
			fmt.Println(errDB)
		}

		engine.Sync(&Struct.User{})
		//engine.CreateTables(&Struct.UserMessage{})
		//engine.CreateTables(&Struct.Admin{})
		//engine.CreateTables(&Struct.AdminMessage{})
		engine.Sync(&Struct.Media{})
		engine.Sync(&Struct.Subtitle{})
		engine.Sync(&Struct.Title{})
		//engine.CreateTables(&Struct.Picture{})
		engine.Sync(&Struct.SubMedia{})
		engine.Sync(&Struct.Message{})
		//engine.CreateTables(&Struct.AdminPicture{})
		engine.Sync(&Struct.File{})
		engine.Sync(&Struct.News{})


		//engine.Sync2(new(Models.User))
		//logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
		//if err != nil {
		//	log.Fatalf("Fail to create xorm system logger: %v\n", err)
		//}
		//
		//logger := xorm.NewSimpleLogger(logWriter)
		//logger.ShowSQL(true)
		//engine.SetLogger(logger)
	}
	return engine
}
