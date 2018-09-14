package api

import (
	"net/http"
	"GoRestful/Controler"
	"encoding/json"
	"log"
	"GoRestful/Models/Struct"
)

func Titles(w http.ResponseWriter, r *http.Request) {
	//if ok, _ := Controler.Authenticated(r); ok {
	//TODO check token
	var users [] Struct.Title
	Controler.GetEngine().Table("title").Cols("Id", "Title").Find(&users)
	var jsonData []byte
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))

	//} else {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//}
}
