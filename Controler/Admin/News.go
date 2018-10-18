package Admin

import (
	"../../Controler"
	"../../Models"
	"../../Models/Struct"
	"github.com/go-xorm/builder"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

func News(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		text := r.PostForm.Get("text")
		fileName := r.PostForm.Get("filename")
		Title := r.PostForm.Get("Title")
		submit := r.PostForm.Get("submit")
		edit, er1 := strconv.Atoi(r.PostForm.Get("edit"))

		result := Models.NewsLayerVariables{
			Answer:      "",
			SubmitValue: "Add News",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			newObj := Struct.NewNews(template.HTML(text), fileName, Title)
			result.Answer = Controler.InsertOrUpdate(&Struct.News{}, newObj, edit, er1)
		}
		http.Redirect(w, r, r.URL.Path+"?result="+result.Answer, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func NewsGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); ok {
		r.ParseForm()
		resultInsert := r.Form.Get("result")
		var news []Struct.News
		Controler.GetEngine().Table(Struct.News{}).AllCols().OrderBy("created").Find(&news)

		var editObject Struct.News
		if vars["id"] != "" {
			Controler.GetEngine().Table(Struct.News{}).Where(builder.Eq{"id": vars["id"]}).Get(&editObject)
			//p.Body = template.HTML(s)
		}

		result := Models.NewsLayerVariables{
			News:          news,
			Answer:        resultInsert,
			PreviousValue: editObject,
			SubmitValue:   "Add News",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "AddNews.html", Models.HeaderVariables{Title: "News"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
