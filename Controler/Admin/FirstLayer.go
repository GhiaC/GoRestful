package Admin

import (
	"../../Controler"
	"../../Models"
	"../../Models/Struct"
	"net/http"
)

func FirstLayer(w http.ResponseWriter, r *http.Request) {
	if ok, _, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		title := r.PostForm.Get("title")
		picture := r.PostForm.Get("picture")
		submit := r.PostForm.Get("submit")

		vars := Models.FirstLayerVariables{
			Answer:      "",
			SubmitValue: "Add Title",
		}

		if submit != "" && (title == "") {
			vars.Answer = "text is empty"
		} else if title != "" {
			engine := Controler.GetEngine()
			new := Struct.NewTitle(title, picture)
			affected, err := engine.Table(Struct.Title{}).Insert(new)
			if affected > 0 && err == nil {
				vars.Answer = "Successful."
			}
		}
		http.Redirect(w, r, r.RequestURI+"?result="+vars.Answer, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}

func FirstLayerGet(w http.ResponseWriter, r *http.Request) {
	if ok, _, _, _ := Controler.Authenticated(r); ok {
		r.ParseForm()
		resultInsert := r.Form.Get("result")
		var titles []Struct.Title
		Controler.GetEngine().Table(Struct.Title{}).Cols("Id", "Title", "Picture").Find(&titles)

		result := Models.FirstLayerVariables{
			Titles:      titles,
			Answer:      resultInsert,
			SubmitValue: "Add Title",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "FirstLayer.html", Models.HeaderVariables{Title: "FirstLayer"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}
