package api

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"encoding/json"
	"log"
	"GoRestful/Models/Struct"
	"github.com/go-xorm/builder"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	Response := Models.SendMessageResponse{
		Error:     "Access Denied.",
		Result:    false,
		MessageId: 0,
	}
	if r.Header["Content-Type"][0] != "application/json" {
		Response.Error = "Content-type is not valid"
		Controler.JsonResponse(w, Response)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest Models.SendMessageRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		log.Fatal(err)
	}
	token := decodedRequest.Token
	text := decodedRequest.Text
	FileAddress := decodedRequest.FileAddress

	if len(token) > 15 {
		ok, id := Authenticated(token)
		if ok {
			if text == "" {
				Response.Error = "Message is empty"
			} else {
				engine := Controler.GetEngine()
				newMessage := Struct.NewMessage(id, 0, text, FileAddress)
				affected, err := engine.Table(Struct.Message{}).Insert(newMessage)
				if affected > 0 && err == nil {
					Response.Result = true
					Response.Error = ""
					Response.MessageId = newMessage.Id
				} else {
					Response.Error = "Database Problem!"
				}
			}
		}
	}
	Controler.JsonResponse(w, Response)
}
func GetMessage(w http.ResponseWriter, r *http.Request) {
	Response := Models.GetMessageResponse{
		Error:  "",
		Result: false,
	}
	if r.Header["Content-Type"][0] != "application/json" {
		Response.Error = "Content-type is not valid"
		Controler.JsonResponse(w, Response)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest Models.GetMessageRequest
	err := decoder.Decode(&decodedRequest)
	token := decodedRequest.Token
	if err != nil {
		log.Fatal(err)
	}
	if len(token) > 15 {
		logged, userid := Authenticated(token)
		if logged {
			var messages []Models.AnswerQuery
			Controler.GetEngine().Table(Struct.Message{}).
				Select("user.username,message.*").
				Join("INNER", Struct.User{}, "message.user_id = user.id ").
				Where(builder.Eq{"user_id": userid}).Or(builder.Eq{"answer_to": userid}).
				Find(&messages)
			Response.Result = true
			Response.Messages = messages
		}
	} else {
		Response.Error = "Access Denied"
	}
	Controler.JsonResponse(w, Response)
}
