package Admin


import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
)

func FirstLayer(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		submit := r.PostForm.Get("submit")

		vars := Models.FirstLayerVariables{
			Answer:      "",
			Url:         "/register",
			SubmitValue: "Add Title",
		}

		if submit != "" && (username == "") {
			vars.Answer = "text is empty"
		} else if username != "" {
			engine := Controler.GetEngine()
			newUser := Models.NewTitle(username)
			affected, err := engine.Table("title").Insert(newUser)
			println(affected)
			if affected > 0 && err == nil {
				vars.Answer = "Successful."
			}
		}

		var titles []Models.Title
		Controler.GetEngine().Table("title").Cols("Id", "Title").Find(&titles)
		vars.Titles = titles
		Controler.OpenTemplate(w, r, vars, "FirstLayer.html", Models.HeaderVariables{Title: "FirstLayer"})

	} else if ok, _ := Controler.Authenticated(r); ok {

		var titles []Models.Title
		Controler.GetEngine().Table("title").Cols("Id", "Title").Find(&titles)

		result := Models.FirstLayerVariables{
			Titles:      titles,
			Answer:      "",
			Url:         "/register",
			SubmitValue: "Add Title",}

		Controler.OpenTemplate(w, r, result, "FirstLayer.html", Models.HeaderVariables{Title: "FirstLayer"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

