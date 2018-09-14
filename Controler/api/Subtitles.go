package api

import (
	"net/http"
	"GoRestful/Controler"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"GoRestful/Models/Struct"
)

func SubTitles(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var users []Struct.Subtitle
	Controler.GetEngine().Table("subtitle").Cols("Id", "Titleid", "Title").Where("Titleid = ?", vars["id"]).Find(&users)
	var jsonData []byte
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}
