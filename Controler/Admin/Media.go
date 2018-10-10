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

func Media(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _, _:= Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		text := r.PostForm.Get("text")
		picture := r.PostForm.Get("picture")
		submit := r.PostForm.Get("submit")

		result := Models.MediaLayerVariables{
			Answer:      "",
			SubmitValue: "Add Media",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			engine := Controler.GetEngine()
			id, _ := strconv.Atoi(vars["id"])
			newUser := Struct.NewMedia(int64(id), text, picture)
			affected, err := engine.Table(Struct.Media{}).Insert(newUser)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			}
		}

		var medias []Struct.Media
		Controler.GetEngine().Table(Struct.Media{}).AllCols().
			//Join("INNER", Struct.Subtitle{}, "subtitle.id = media.pid ").
			Where(builder.Eq{"pid": vars["id"]}).
			Find(&medias)
		result.OptionFiles = Controler.Files()
		result.Medias = medias
		Controler.OpenTemplate(w, r, result, "AddMedia.html", Models.HeaderVariables{Title: "Medias"})

	} else if ok, _, _ ,_:= Controler.Authenticated(r); ok {

		var medias []Struct.Media
		Controler.GetEngine().Table(Struct.Media{}).AllCols().
			Where(builder.Eq{"pid": vars["id"]}).
			Find(&medias)

		result := Models.MediaLayerVariables{
			TitleId:     vars["id"],
			Medias:      medias,
			Answer:      "",
			SubmitValue: "Add Media",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "AddMedia.html", Models.HeaderVariables{Title: "Media"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
