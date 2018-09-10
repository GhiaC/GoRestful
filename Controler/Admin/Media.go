package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"github.com/gorilla/mux"
	"strconv"
)

func Media(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		text := r.PostForm.Get("text")
		//text := r.PostForm.Get("text")
		submit := r.PostForm.Get("submit")

		result := Models.MediaLayerVariables{
			Answer:      "",
			Url:         "/register",
			SubmitValue: "Add Media",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			engine := Controler.GetEngine()
			id, _ := strconv.Atoi(vars["id"])
			newUser := Models.NewMedia(int64(id), text)
			affected, err := engine.Table("media").Insert(newUser)
			println(affected)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			}
		}

		var users []Models.Media
		Controler.GetEngine().Table("media").Select("media.*,subtitle.*").
			Join("INNER", "subtitle", "subtitle.id = media.subtitleid ").Where("subtitleid = ?", vars["id"]).Find(&users)

		result.Medias = users
		Controler.OpenTemplate(w, r, result, "AddMedia.html", Models.HeaderVariables{Title: "Medias"})

	} else if ok, _ := Controler.Authenticated(r); ok {

		var medias []Models.Media
		Controler.GetEngine().Table("media").Select("media.*,subtitle.*").
			Join("INNER", "subtitle", "subtitle.id = media.subtitleid ").Where("subtitleid = ?", vars["id"]).Find(&medias)

		result := Models.MediaLayerVariables{
			TitleId:     vars["id"],
			Medias:   medias,
			Answer:      "",
			Url:         "/register",
			SubmitValue: "Add Media",}

		Controler.OpenTemplate(w, r, result, "AddMedia.html", Models.HeaderVariables{Title: "Media"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}