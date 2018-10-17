package Admin

import (
	"../../Controler"
	"../../Models"
	"../../Models/Struct"
	"html"
	"net/http"
)

func News(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		text := r.PostForm.Get("text")
		fileName := r.PostForm.Get("filename")
		Title := r.PostForm.Get("Title")
		submit := r.PostForm.Get("submit")

		result := Models.NewsLayerVariables{
			Answer:      "",
			SubmitValue: "Add News",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			engine := Controler.GetEngine()
			//id, _ := strconv.Atoi(vars["id"])
			newUser := Struct.NewNews(html.UnescapeString(text), fileName, Title)
			affected, err := engine.Table(Struct.News{}).Insert(newUser)
			println(affected)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			}
		}
		http.Redirect(w, r, r.RequestURI+"?result="+result.Answer, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func NewsGet(w http.ResponseWriter, r *http.Request) {
	if ok, _, _, _ := Controler.Authenticated(r); ok {
		r.ParseForm()
		resultInsert := r.Form.Get("result")
		var news []Struct.News
		Controler.GetEngine().Table(Struct.News{}).AllCols().OrderBy("created").Find(&news)

		result := Models.NewsLayerVariables{
			News:        news,
			Answer:      resultInsert,
			SubmitValue: "Add News",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "AddNews.html", Models.HeaderVariables{Title: "News"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
