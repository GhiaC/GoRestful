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
	if ok, _, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		text := r.PostForm.Get("text")
		picture := r.PostForm.Get("picture")
		submit := r.PostForm.Get("submit")

		result := Models.MediaLayerVariables{
			Answer:      "",
			SubmitValue: "افزودن مدیا",
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

		http.Redirect(w, r, r.RequestURI+"?result="+result.Answer, http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}

func MediaGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _, _ := Controler.Authenticated(r); ok {
		r.ParseForm()
		resultInsert := r.Form.Get("result")
		var medias []Struct.Media
		Controler.GetEngine().Table(Struct.Media{}).AllCols().
			Where(builder.Eq{"pid": vars["id"]}).
			Find(&medias)

		result := Models.MediaLayerVariables{
			TitleId:     vars["id"],
			Medias:      medias,
			Answer:      resultInsert,
			SubmitValue: "افزودن مدیا",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "AddMedia.html", Models.HeaderVariables{Title: "Media"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
