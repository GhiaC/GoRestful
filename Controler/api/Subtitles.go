package api

import (
	"net/http"
	"restful/Models"
	"restful/Controler"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

func SubTitles(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var users []Models.Subtitle
	Controler.GetEngine().Table("subtitle").Cols("Id", "Titleid", "Title").Where("Titleid = ?", vars["id"]).Find(&users)
	var jsonData []byte
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}
