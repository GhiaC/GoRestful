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
		edit, er1 := strconv.Atoi(r.PostForm.Get("edit"))
		pid, _ := strconv.Atoi(vars["pid"])

		result := Models.SecondLayerVariables{
			Answer:      "",
			SubmitValue: "افزودن زیرعنوان",
		}

		if submit != "" && (username == "") {
			result.Answer = "text is empty"
		} else if username != "" {
			newUser := Struct.NewSubtitle(int64(pid), username, pic1)
			result.Answer = Controler.InsertOrUpdate(&Struct.Subtitle{}, newUser, edit, er1)
		}
		http.Redirect(w, r, r.URL.Path+"?result="+result.Answer, http.StatusSeeOther)
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
		Controler.GetEngine().Table(Struct.Subtitle{}).AllCols().Where(builder.Eq{"Pid": vars["pid"]}).Find(&subtitles)

		var editSubtitles Struct.Subtitle
		if vars["id"] != "" {
			Controler.GetEngine().Table(Struct.Subtitle{}).Where(builder.Eq{"id": vars["id"]}).Get(&editSubtitles)
		}

		result := Models.SecondLayerVariables{
			TitleId:       vars["pid"],
			Subtitles:     subtitles,
			Answer:        resultInsert,
			PreviousValue: editSubtitles,
			SubmitValue:   "افزودن زیرعنوان",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "SecondLayer.html", Models.HeaderVariables{Title: "SecondLayer"})
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}
