package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"github.com/gorilla/mux"
	"strconv"
	"GoRestful/Models/Struct"
	"github.com/go-xorm/builder"
)

func Messages(w http.ResponseWriter, r *http.Request) {
	if ok, _, _ := Controler.Authenticated(r); ok {
		var messages []Struct.Message

		Controler.GetEngine().Table(Struct.Message{}).AllCols().Distinct("user_id").Find(&messages)
		result := Models.MessagesLayerVariables{
			Messages: messages,
		}
		Controler.OpenTemplate(w, r, result, "Messages.html", Models.HeaderVariables{Title: "Messages"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Answer(w http.ResponseWriter, r *http.Request) { // ajib daghon
	vars := mux.Vars(r)
	if ok, _, UserId := Controler.Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		text := r.PostForm.Get("text")
		fileAddress := r.PostForm.Get("fileaddress")
		submit := r.PostForm.Get("submit")

		result := Models.AnswerLayerVariables{
			Answer:      "",
			SubmitValue: "Send Answer",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			engine := Controler.GetEngine()
			id, _ := strconv.Atoi(vars["id"])
			newMessage := Struct.NewMessage(UserId, id, text, fileAddress)
			affected, err := engine.Table(Struct.Media{}).Insert(newMessage)
			println(affected)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			}
		}

		var msg Models.AnswerQuery

		Controler.GetEngine().Table(Struct.Message{}).
			Select("user.username,message.*").
			Join("INNER", Struct.User{}, "message.user_id = user.id ").
			Where(builder.Eq{"message.id": vars["id"]}).Get(&msg)
		result.Msg = msg

		Controler.OpenTemplate(w, r, result, "AddMedia.html", Models.HeaderVariables{Title: "Medias"})

	} else if ok, _, _ := Controler.Authenticated(r); ok {

		var msg Models.AnswerQuery

		Controler.GetEngine().Table(Struct.Message{}).
			Select("user.username,message.*").
			Join("INNER", Struct.User{}, "message.user_id = user.id ").
			Where(builder.Eq{"message.id": vars["id"]}).Get(&msg)
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
