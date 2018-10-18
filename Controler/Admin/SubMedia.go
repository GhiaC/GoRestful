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
		edit, er1 := strconv.Atoi(r.PostForm.Get("edit"))

		result := Models.SubMediaLayerVariables{
			Answer:      "",
			SubmitValue: "Add SubMedia",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			id, _ := strconv.Atoi(vars["pid"])
			newObj := Struct.NewSubMedia(int64(id), picture, text)
			result.Answer = Controler.InsertOrUpdate(&Struct.SubMedia{}, newObj, edit, er1)
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
			Where(builder.Eq{"pid": vars["pid"]}).
			Find(&medias)

		var editObject Struct.SubMedia
		if vars["id"] != "" {
			Controler.GetEngine().Table(Struct.SubMedia{}).Where(builder.Eq{"id": vars["id"]}).Get(&editObject)
		}

		result := Models.SubMediaLayerVariables{
			TitleId:     vars["pid"],
			SubMedias:   medias,
			Answer:      resultInsert,
			PreviousValue : editObject,
			SubmitValue: "Add SubMedia",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "SubMedia.html", Models.HeaderVariables{Title: "SubMedia"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
