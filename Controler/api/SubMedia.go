package api

import (
	"net/http"
	"../../Controler"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"../../Models/Struct"
	"../../Models"
)

func SubMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var SubMedias []Models.SubMediaJoinFile
	Controler.GetEngine().Table(Struct.SubMedia{}).Select("sub_media.*,file.type").
		Join("LEFT", Struct.File{}, "sub_media.url = file.key").
		Where("pid = ?", vars["id"]).Find(&SubMedias)
	var jsonData []byte

	jsonData, err := json.Marshal(SubMedias)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(jsonData)))
}

func AllSubMedia(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	var SubMedias []Models.SubMediaJoinFile
	Controler.GetEngine().Table(Struct.SubMedia{}).Select("sub_media.*,file.type").
		Join("LEFT", Struct.File{}, "sub_media.url = file.key").
		Find(&SubMedias)
	var jsonData []byte

	jsonData, err := json.Marshal(SubMedias)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(jsonData)))
}
