package Controler

import (
	"github.com/go-xorm/builder"
)

func InsertOrUpdate(table interface{}, object interface{}, edit int, editErr error) string {
	engine := GetEngine()
	var err error
	var affected int64
	if edit > 0 && editErr == nil {
		affected, err = engine.Where(builder.Eq{"id": edit}).Update(object)
	} else {
		affected, err = engine.Table(table).Insert(object)
	}
	if affected > 0 && err == nil {
		return "Successful."
	} else {
		return "Database Problem!"
	}
}
