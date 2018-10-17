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

func SecondLayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		submit := r.PostForm.Get("submit")
		pic1 := r.PostForm.Get("pic1")
		id, _ := strconv.Atoi(vars["id"])

		result := Models.SecondLayerVariables{
			Answer:      "",
			SubmitValue: "افزودن زیرعنوان",
		}

		if submit != "" && (username == "") {
			result.Answer = "text is empty"
		} else if username != "" {
			engine := Controler.GetEngine()
			newUser := Struct.NewSubtitle(int64(id), username, pic1)
			affected, err := engine.Table(Struct.Subtitle{}).Insert(newUser)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			} else {
				result.Answer = "Database Problem!"
			}
		}
		http.Redirect(w, r, r.RequestURI+"?result="+result.Answer, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}

func SecondLayerGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); ok {
		r.ParseForm()
		resultInsert := r.Form.Get("result")
		var subtitles []Struct.Subtitle
		Controler.GetEngine().Table(Struct.Subtitle{}).AllCols().Where(builder.Eq{"Pid": vars["id"]}).Find(&subtitles)

		result := Models.SecondLayerVariables{
			TitleId:     vars["id"],
			Subtitles:   subtitles,
			Answer:      resultInsert,
			SubmitValue: "افزودن زیرعنوان",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "SecondLayer.html", Models.HeaderVariables{Title: "SecondLayer"})
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}
