package api

import (
	"net/http"
	"GoRestful/Controler"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"GoRestful/Models/Struct"
)

func Media(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var users []Struct.Media
	Controler.GetEngine().Table(Struct.Media{}).AllCols().
	//Join("INNER", Struct.Subtitle{}, "subtitle.id = media.Pid ").
		Where("pid = ?", vars["id"]).Find(&users)
	var jsonData []byte

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(jsonData)))
}

func AllMedia(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	var users []Struct.Media
	Controler.GetEngine().Table(Struct.Media{}).AllCols().
	//Join("INNER", Struct.Subtitle{}, "subtitle.id = media.subtitleid ").
		Find(&users)
	var jsonData []byte

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(jsonData)))
}
