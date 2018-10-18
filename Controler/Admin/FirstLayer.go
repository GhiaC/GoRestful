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

func FirstLayer(w http.ResponseWriter, r *http.Request) {
	if ok, _, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		title := r.PostForm.Get("title")
		picture := r.PostForm.Get("picture")
		submit := r.PostForm.Get("submit")
		edit, er1 := strconv.Atoi(r.PostForm.Get("edit"))

		vars := Models.FirstLayerVariables{
			Answer:      "",
			SubmitValue: "Add Title",
		}

		if submit != "" && (title == "") {
			vars.Answer = "text is empty"
		} else if title != "" {
			new := Struct.NewTitle(title, picture)
			vars.Answer = Controler.InsertOrUpdate(&Struct.Title{}, new, edit, er1)
		}
		http.Redirect(w, r, r.URL.Path+"?result="+vars.Answer, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}

func FirstLayerGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); ok {
		r.ParseForm()
		resultInsert := r.Form.Get("result")
		var titles []Struct.Title
		Controler.GetEngine().Table(Struct.Title{}).Cols("Id", "Title", "Picture").Find(&titles)

		var editObject Struct.Title
		if vars["id"] != "" {
			Controler.GetEngine().Table(Struct.Title{}).Where(builder.Eq{"id": vars["id"]}).Get(&editObject)
		}

		result := Models.FirstLayerVariables{
			Titles:        titles,
			Answer:        resultInsert,
			PreviousValue: editObject,
			SubmitValue:   "Add Title",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "FirstLayer.html", Models.HeaderVariables{Title: "FirstLayer"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}
