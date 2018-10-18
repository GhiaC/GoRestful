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
		edit, er1 := strconv.Atoi(r.PostForm.Get("edit"))

		result := Models.MediaLayerVariables{
			Answer:      "",
			SubmitValue: "افزودن مدیا",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			id, _ := strconv.Atoi(vars["pid"])
			new := Struct.NewMedia(int64(id), text, picture)
			result.Answer = Controler.InsertOrUpdate(&Struct.Media{}, new, edit, er1)
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
			Where(builder.Eq{"pid": vars["pid"]}).
			Find(&medias)

		var editMedia Struct.Media
		if vars["id"] != "" {
			Controler.GetEngine().Table(Struct.Media{}).Where(builder.Eq{"id": vars["id"]}).Get(&editMedia)
		}

		result := Models.MediaLayerVariables{
			TitleId:       vars["pid"],
			Medias:        medias,
			Answer:        resultInsert,
			PreviousValue: editMedia,
			SubmitValue:   "افزودن مدیا",}
		result.OptionFiles = Controler.Files()
		Controler.OpenTemplate(w, r, result, "AddMedia.html", Models.HeaderVariables{Title: "Media"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
