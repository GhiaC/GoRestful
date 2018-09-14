package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"strconv"
	"github.com/gorilla/mux"
	"GoRestful/Models/Struct"
	"github.com/go-xorm/builder"
)

func SecondLayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		submit := r.PostForm.Get("submit")
		pic1 := r.PostForm.Get("pic1")
		pic2 := r.PostForm.Get("pic2")
		id, _ := strconv.Atoi(vars["id"])

		result := Models.SecondLayerVariables{
			Answer:      "",
			SubmitValue: "Add Subtitle",
		}

		if submit != "" && (username == "") {
			result.Answer = "text is empty"
		} else if username != "" {
			engine := Controler.GetEngine()
			newUser := Struct.NewSubtitle(int64(id), username, pic1, pic2)
			affected, err := engine.Table(Struct.Subtitle{}).Insert(newUser)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			}else {
				result.Answer = "Database Problem!"
			}
		}

		var subtitles []Struct.Subtitle
		Controler.GetEngine().Table(Struct.Subtitle{}).AllCols().Where(builder.Eq{"TitleId": int64(id)}).Find(&subtitles)

		result.Subtitles = subtitles
		Controler.OpenTemplate(w, r, result, "SecondLayer.html", Models.HeaderVariables{Title: "SecondLayer"})

	} else if ok, _, _ := Controler.Authenticated(r); ok {
		var subtitles []Struct.Subtitle
		Controler.GetEngine().Table(Struct.Subtitle{}).AllCols().Where(builder.Eq{"title_id": vars["id"]}).Find(&subtitles)

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
