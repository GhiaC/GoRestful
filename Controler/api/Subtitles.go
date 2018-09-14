package api

import (
	"net/http"
	"GoRestful/Controler"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"GoRestful/Models/Struct"
	"strconv"
)

func SubTitles(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var users []Struct.Subtitle
	Controler.GetEngine().Table(Struct.Subtitle{}).AllCols().Where(Struct.Subtitle{TitleId: int64(id)}).Find(&users)
	var jsonData []byte
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}
