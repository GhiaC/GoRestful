package Admin

import (
	"../../Controler"
	"../../Models"
	"../../Models/Struct"
	"github.com/go-xorm/builder"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func SubMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		text := r.PostForm.Get("text")
		picture := r.PostForm.Get("picture")
		submit := r.PostForm.Get("submit")

		result := Models.SubMediaLayerVariables{
			Answer:      "",
			SubmitValue: "Add SubMedia",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			engine := Controler.GetEngine()
			id, _ := strconv.Atoi(vars["id"])
			newUser := Struct.NewSubMedia(int64(id), picture, text)
			affected, err := engine.Table(Struct.SubMedia{}).Insert(newUser)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			}
		}

		http.Redirect(w, r, r.RequestURI+"?result="+result.Answer, http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}

func SubMediaGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); ok {
		r.ParseForm()
		resultInsert := r.Form.Get("result")
		var medias []Struct.SubMedia
		Controler.GetEngine().Table(Struct.SubMedia{}).AllCols().
			Where(builder.Eq{"pid": vars["id"]}).
			Find(&medias)

		result := Models.SubMediaLayerVariables{
			TitleId:     vars["id"],
			SubMedias:   medias,
			Answer:      resultInsert,
			SubmitValue: "Add SubMedia",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "SubMedia.html", Models.HeaderVariables{Title: "SubMedia"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
