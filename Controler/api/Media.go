package api

import (
	"net/http"
	"restful/Models"
	"restful/Controler"
	"encoding/json"
	"log"
)

func Media(w http.ResponseWriter, r *http.Request) {
	var users []Models.Media
	Controler.GetEngine().Table("media").Cols("Id", "SubTitleId","Text").Find(&users)
	var jsonData []byte
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Write([]byte(string(jsonData)))
}
