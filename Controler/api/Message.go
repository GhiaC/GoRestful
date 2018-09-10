package api

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"encoding/json"
	"log"
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

	Response := Models.SendMessageResponse{
		Error:     "",
		Result:    false,
		MessageId: 0,
	}
	if ok, id := Authenticated(token); ok {
		if text == "" {
			Response.Error = "Message is empty"
		} else {
			engine := Controler.GetEngine()
			newMessage := Models.NewUserMessage(id, text)
			affected, err := engine.Table("UserMessage").Insert(newMessage)
			//println(affected)
			if affected > 0 && err == nil {
				Response.Result = true
				Response.Error = ""
				Response.MessageId = newMessage.Id
			}
		}

		var jsonData []byte
		jsonData, err := json.Marshal(Response)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(string(jsonData)))
	}
}
func GetMessage(w http.ResponseWriter, r *http.Request) {
	//TODO complete
	//vars := mux.Vars(r)

	decoder := json.NewDecoder(r.Body)
	var decodedRequest Models.SendMessageRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		panic(err)
	}
	token := decodedRequest.Token
	logged, userid := Authenticated(token)

	Response := Models.GetMessageResponse{
		Error:  "",
		Result: false,
	}
	if logged {
		var messages1 []Models.UserMessage
		Controler.GetEngine().Table("User_Message").Cols("id", "user_id", "text", "created").
			Where("user_id = ?", userid).
			Find(&messages1)
		var messages2 []Models.AdminMessage
		Controler.GetEngine().Table("admin_Message").Cols("id", "user_id", "username", "text", "created").
			Join("INNER", "admin", "admin.Id = admin_message.admin_id ").
			Where("user_id = ?", userid).
			Find(&messages2)

		Response.Result = true
		Response.UserMessages = messages1
		Response.AdminMessages = messages2
	} else {
		Response.Error = "Access Denied"
	}

	var jsonData []byte
	jsonData, err = json.Marshal(Response)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonData)))
}
