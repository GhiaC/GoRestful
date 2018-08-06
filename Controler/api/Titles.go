package api

import (
	"net/http"
	"restful/Models"
	"restful/Controler"
	"encoding/json"
	"log"
)

func Titles(w http.ResponseWriter, r *http.Request) {
	//if ok, _ := Controler.Authenticated(r); ok {

	var users []Models.Title
	Controler.GetEngine().Table("title").Cols("Id", "Title").Find(&users)
	var jsonData []byte
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Write([]byte(string(jsonData)))

	//} else {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//}
}
