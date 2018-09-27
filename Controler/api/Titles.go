package api

import (
	"net/http"
	"../../Controler"
	"encoding/json"
	"log"
	"../../Models/Struct"
	"../../Models"
)

func Titles(w http.ResponseWriter, r *http.Request) {
	//TODO check token
	var users [] Models.TitleJoinFile
	Controler.GetEngine().Table(Struct.Title{}).Select("title.*,file.type").
		Join("LEFT", Struct.File{}, "title.picture = file.key").
		Find(&users)
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
