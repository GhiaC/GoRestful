package api

import (
	"net/http"
	"restful/Models"
	"restful/Controler"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

func Media(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var users []Models.Media
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
