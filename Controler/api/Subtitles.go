package api

import (
	"net/http"
	"restful/Models"
	"restful/Controler"
	"encoding/json"
	"log"
)

func SubTitles(w http.ResponseWriter, r *http.Request) {
	var users []Models.Subtitle
	Controler.GetEngine().Table("subtitle").Cols("Id", "TitleId","Title").Find(&users)
	var jsonData []byte
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Write([]byte(string(jsonData)))
}
