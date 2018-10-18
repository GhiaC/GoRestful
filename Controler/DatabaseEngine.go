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

		engine.Update(&Struct.User{})
		//engine.CreateTables(&Struct.UserMessage{})
		//engine.CreateTables(&Struct.Admin{})
		//engine.CreateTables(&Struct.AdminMessage{})
		engine.CreateTables(&Struct.Media{})
		engine.CreateTables(&Struct.Subtitle{})
		engine.CreateTables(&Struct.Title{})
		//engine.CreateTables(&Struct.Picture{})
		engine.CreateTables(&Struct.SubMedia{})
		engine.CreateTables(&Struct.Message{})
		//engine.CreateTables(&Struct.AdminPicture{})
		engine.CreateTables(&Struct.File{})
		engine.Update(&Struct.News{})


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
