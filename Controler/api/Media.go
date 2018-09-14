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
	Controler.GetEngine().Table("media").Select("media.*,subtitle.*").
		Join("INNER", "subtitle", "subtitle.id = media.subtitleid ").Where("subtitleid = ?", vars["id"]).Find(&users)
	var jsonData []byte

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}


func AllMedia(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	var users []Struct.Media
	Controler.GetEngine().Table("media").Select("media.*,subtitle.*").
		Join("INNER", "subtitle", "subtitle.id = media.subtitleid ").Find(&users)
	var jsonData []byte

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}
