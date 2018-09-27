package Controler

import (
	"../Models/Struct"
	"github.com/go-xorm/builder"
)

func Files() []Struct.File {
	var files []Struct.File
	GetEngine().Table(Struct.File{}).AllCols().Where(builder.Eq{"admin_file": true}).Find(&files)
	return files
}
