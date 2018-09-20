package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"GoRestful/Models/Struct"
	"html"
)



func News(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	if ok, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
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

		var news []Struct.News
		Controler.GetEngine().Table(Struct.News{}).AllCols().OrderBy("created").Find(&news)

		result.News = news
		Controler.OpenTemplate(w, r, result, "AddNews.html", Models.HeaderVariables{Title: "News"})

	} else if ok, _, _ := Controler.Authenticated(r); ok {

		var news []Struct.News
		Controler.GetEngine().Table(Struct.News{}).AllCols().OrderBy("created").Find(&news)

		result := Models.NewsLayerVariables{
			News:        news,
			Answer:      "",
			SubmitValue: "Add News",}

		Controler.OpenTemplate(w, r, result, "AddNews.html", Models.HeaderVariables{Title: "News"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
