package Admin

import (
	"../../Controler"
	"../../Models/Struct"
	"github.com/go-xorm/builder"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	removeGeneral(w, r, &Struct.User{}, "/admin/users")
}

func RemoveFirstLayer(w http.ResponseWriter, r *http.Request) {
	removeGeneral(w, r, &Struct.Title{}, "/admin/FirstLayer/")
}

func RemoveSecondLayer(w http.ResponseWriter, r *http.Request) {
	removeGeneral(w, r, &Struct.Subtitle{}, "/admin/SecondLayer/")
}

func RemoveMedia(w http.ResponseWriter, r *http.Request) {
	removeGeneral(w, r, &Struct.Media{}, "/admin/Media/")
}

func RemoveSubMedia(w http.ResponseWriter, r *http.Request) {
	removeGeneral(w, r, &Struct.SubMedia{}, "/admin/SubMedia/")
}

func RemoveAdmin(w http.ResponseWriter, r *http.Request) {
	removeGeneral(w, r, &Struct.User{}, "/admin/admins")
}

func RemoveNews(w http.ResponseWriter, r *http.Request) {
	removeGeneral(w, r, &Struct.News{}, "/admin/news")
}

func RemoveFile(w http.ResponseWriter, r *http.Request) {
	removeGeneral(w, r, &Struct.File{}, "/admin/upload")
}

func RemoveMessages(w http.ResponseWriter, r *http.Request) {
	removeGeneral(w, r, &Struct.Message{}, "/admin/messages")
}

func removeGeneral(w http.ResponseWriter, r *http.Request, model interface{}, previousPage string) {
	vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); !(ok) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		id, _ := strconv.Atoi(vars["id"])
		//pid, _ := strconv.Atoi(vars["pid"])
		engine := Controler.GetEngine()
		affected, err := engine.Where(builder.Eq{"id": id}).Limit(1).Delete(model)
		if affected > 0 && err == nil {
			http.Redirect(w, r, previousPage+vars["pid"], http.StatusSeeOther)
			//w.Write([]byte("Remove Successful."))
		} else {
			//w.Write([]byte("Remove Failed."))
		}
	}
}
