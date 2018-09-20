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

func SendMessage(w http.ResponseWriter, r *http.Request) { //TODO

	decoder := json.NewDecoder(r.Body)
	var decodedRequest Models.SendMessageRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		panic(err)
	}
	token := decodedRequest.Token
	text := decodedRequest.Text
	FileAddress := decodedRequest.FileAddress
	Response := Models.SendMessageResponse{
		Error:     "Access Denied.",
		Result:    false,
		MessageId: 0,
	}
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
	var jsonData []byte
	jsonData, err = json.Marshal(Response)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))

}
func GetMessage(w http.ResponseWriter, r *http.Request) {
	Response := Models.GetMessageResponse{
		Error:  "",
		Result: false,
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest Models.GetMessageRequest
	err := decoder.Decode(&decodedRequest)
	token := decodedRequest.Token
	if err != nil {
		panic(err)
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

	var jsonData []byte
	jsonData, err = json.Marshal(Response)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}
