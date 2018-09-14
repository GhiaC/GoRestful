package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"github.com/gorilla/mux"
	"strconv"
	"GoRestful/Models/Struct"
)

func Messages(w http.ResponseWriter, r *http.Request) {
	if ok, _, _ := Controler.Authenticated(r); ok {
		var messages []Struct.UserMessage
		Controler.GetEngine().Table(Struct.UserMessage{}).AllCols().Find(&messages)
		result := Models.MessagesLayerVariables{
			Messages: messages,
		}
		Controler.OpenTemplate(w, r, result, "Messages.html", Models.HeaderVariables{Title: "Messages"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Answer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _, _ := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		text := r.PostForm.Get("text")
		//text := r.PostForm.Get("text")
		submit := r.PostForm.Get("submit")

		result := Models.MediaLayerVariables{
			Answer:      "",
			SubmitValue: "Send Answer",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			engine := Controler.GetEngine()
			id, _ := strconv.Atoi(vars["id"])
			newUser := Struct.NewMedia(int64(id), text)
			affected, err := engine.Table("media").Insert(newUser)
			println(affected)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			}
		}

		var users []Struct.Media
		Controler.GetEngine().Table("media").Select("media.*,subtitle.*").
			Join("INNER", "subtitle", "subtitle.id = media.subtitleid ").Where("subtitleid = ?", vars["id"]).Find(&users)

		result.Medias = users
		Controler.OpenTemplate(w, r, result, "AddMedia.html", Models.HeaderVariables{Title: "Medias"})

	} else if ok, _, _ := Controler.Authenticated(r); ok {
		var msg Models.AnswerQuery
		Controler.GetEngine().Table("user_message").Select("user.username,user_message.*").
			Join("INNER", "user", "user_message.user_id = user.id ").Where("user_message.id = ?", vars["id"]).Get(&msg)
		result := Models.AnswerLayerVariables{
			TitleId:     vars["id"],
			Msg:         msg,
			Answer:      "",
			SubmitValue: "Send Answer",}

		Controler.OpenTemplate(w, r, result, "Answer.html", Models.HeaderVariables{Title: "Answer"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
