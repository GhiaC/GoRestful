package Controler

import (
	"github.com/go-xorm/xorm"
	"fmt"
	//"GoRestful/Models"
	//"log"
	//"log/syslog"
	"GoRestful/Models"
)

var engine *xorm.Engine
var flag bool

const userDB = "root"
const passDB  = ""
const nameDB  = "restful"

func GetEngine() *xorm.Engine{
	if !flag {
		flag = true
		var errDB error
		engine, errDB = xorm.NewEngine("mysql", userDB+":"+passDB+"@/"+nameDB+"?charset=utf8")
		if errDB != nil {
			fmt.Println(errDB)
		}

		//engine.CreateTables(&Models.User{})
		//engine.CreateTables(&Models.UserMessage{})
		//engine.CreateTables(&Models.Admin{})
		//engine.CreateTables(&Models.AdminMessage{})
		//engine.CreateTables(&Models.Media{})
		//engine.CreateTables(&Models.Subtitle{})
		//engine.CreateTables(&Models.Title{})
		engine.CreateTables(&Models.Picture{})

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
