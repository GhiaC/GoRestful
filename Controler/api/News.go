package api

import (
	"net/http"
	"encoding/json"
	"log"
	"GoRestful/Controler"
	"GoRestful/Models/Struct"
)

func AllNews(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	var news []Struct.News
	Controler.GetEngine().Table("news").AllCols().Find(&news)
	var jsonData []byte

	jsonData, err := json.Marshal(news)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}

