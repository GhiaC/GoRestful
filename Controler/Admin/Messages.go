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

func Messages(w http.ResponseWriter, r *http.Request) {
	if ok, _, _, _ := Controler.Authenticated(r); ok {
		var messages []Models.AnswerQuery
		//.Distinct("user_id")
		Controler.GetEngine().Table(Struct.Message{}).
			Select("user.username,message.*,file.Type").
			Join("INNER", Struct.User{}, "message.user_id = user.id ").
			Join("LEFT", Struct.File{}, "message.file_address = file.key").
			Where(builder.Eq{"answer_to": 0}).OrderBy("message.created").AllCols().Find(&messages)
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
	message_id, _ := strconv.Atoi(vars["id"])
	if ok, _, UserId, _ := Controler.Authenticated(r); ok && r.Method == "POST" && message_id > 0 {
		r.ParseForm()
		text := r.PostForm.Get("text")
		fileAddress := r.PostForm.Get("fileaddress")
		submit := r.PostForm.Get("submit")

		result := Models.AnswerLayerVariables{
			Answer:      "",
			SubmitValue: "ارسال پاسخ",
		}

		if submit != "" && (text == "") {
			result.Answer = "text is empty"
		} else if text != "" {
			engine := Controler.GetEngine()
			newMessage := Struct.NewMessage(UserId, message_id, text, fileAddress)
			affected, err := engine.Table(Struct.Message{}).Insert(newMessage)
			if affected > 0 && err == nil {
				phonenumber := Controler.GetUserPhonenumber(message_id)
				if phonenumber != "Error" {
					go Controler.SendSms(phonenumber, message_id, text, true)
				}
				result.Answer = "Successful."
			}
		}
		http.Redirect(w, r, r.RequestURI+"?result="+result.Answer, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}

func AnswerGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message_id, _ := strconv.Atoi(vars["id"])
	if ok, _, _, _ := Controler.Authenticated(r); ok {
		r.ParseForm()
		resultInsert := r.Form.Get("result")
		var msg Models.AnswerQuery
		var answers [] Models.AnswerQuery

		Controler.GetEngine().Table(Struct.Message{}).
			Select("user.username,message.*").
			Join("INNER", Struct.User{}, "message.user_id = user.id ").
			Where(builder.Eq{"message.id": message_id}).Get(&msg)

		Controler.GetEngine().Table(Struct.Message{}).
			Select("user.username,message.*").
			Join("INNER", Struct.User{}, "message.user_id = user.id ").
			Where(builder.Eq{"message.answer_to": message_id}).Find(&answers)

		result := Models.AnswerLayerVariables{
			TitleId:     vars["id"],
			Msg:         msg,
			Answer:      resultInsert,
			Answers:     answers,
			SubmitValue: "ارسال پاسخ",}

		Controler.OpenTemplate(w, r, result, "Answer.html", Models.HeaderVariables{Title: "Answer"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
