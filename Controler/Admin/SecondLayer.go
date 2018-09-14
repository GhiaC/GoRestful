package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"strconv"
	"github.com/gorilla/mux"
	"GoRestful/Models/Struct"
)

func SecondLayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _ ,_:= Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		submit := r.PostForm.Get("submit")

		result := Models.SecondLayerVariables{
			Answer:      "",
			//Url:         "/register",
			SubmitValue: "Add Subtitle",
		}

		if submit != "" && (username == "") {
			result.Answer = "text is empty"
		} else if username != "" {
			engine := Controler.GetEngine()
			id, _ := strconv.Atoi(vars["id"])
			newUser := Struct.NewSubtitle(int64(id), username)
			affected, err := engine.Table("subtitle").Insert(newUser)
			println(affected)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			}
		}

		var subtitles []Struct.Subtitle
		Controler.GetEngine().Table("subtitle").AllCols().Where("title_id = ?", vars["id"]).Find(&subtitles)

		result.Subtitles = subtitles
		Controler.OpenTemplate(w, r, result, "SecondLayer.html", Models.HeaderVariables{Title: "SecondLayer"})

	} else if ok, _ ,_:= Controler.Authenticated(r); ok {
		var subtitles []Struct.Subtitle
		Controler.GetEngine().Table("subtitle").AllCols().Where("title_id = ?", vars["id"]).Find(&subtitles)

		result := Models.SecondLayerVariables{
			TitleId:     vars["id"],
			Subtitles:   subtitles,
			Answer:      "",
			SubmitValue: "Add Subtitle",}

		Controler.OpenTemplate(w, r, result, "SecondLayer.html", Models.HeaderVariables{Title: "SecondLayer"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

