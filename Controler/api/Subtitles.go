package api

import (
	"net/http"
	"GoRestful/Controler"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"GoRestful/Models/Struct"
	"strconv"
	"github.com/go-xorm/builder"
	"GoRestful/Models"
)

func SubTitles(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var users []Models.SubtitleJoinFile
	Controler.GetEngine().Table(Struct.Subtitle{}).Select("subtitle.*,file.type").
		Join("LEFT", Struct.File{}, "subtitle.picture = file.key").
		Where(builder.Eq{"Pid": int64(id)}).Find(&users)
	var jsonData []byte
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}

func AllSubTitles(w http.ResponseWriter, r *http.Request) {
	var users []Struct.Subtitle
	Controler.GetEngine().Table(Struct.Subtitle{}).AllCols().Find(&users)
	var jsonData []byte
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}
