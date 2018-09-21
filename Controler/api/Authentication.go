package api

import (
	"net/http"
	"encoding/json"
	"GoRestful/Models"
	"GoRestful/Controler"
	"log"
	"GoRestful/Models/Struct"
	"github.com/go-xorm/builder"
)

func Login(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var decodedRequest Models.LoginRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		panic(err)
	}
	username := decodedRequest.Username
	password := decodedRequest.Password

	Response := Models.LoginResponse{
		Result: false,
		Error:  "",
		Token:  "",
	}

	if username == "" || password == "" {
		Response.Error = "username or password is empty"
	} else if username != "" && password != "" {
		var id int
		engine := Controler.GetEngine()
		has, err := engine.Table(Struct.User{}).
			Where(builder.Eq{"Username": username, "Password": password, "Type": 2}).
			Cols("id").Get(&id)
		if has && err == nil && id > 0 {
			Token :=
				Controler.TokenGenerator() + Controler.TokenGenerator() +
					Controler.TokenGenerator() + Controler.TokenGenerator() +
					Controler.TokenGenerator() + Controler.TokenGenerator()
			engine.Table(Struct.User{}).Omit("id", "username", "password", "created").
				Update(Struct.User{Token: Token})
			Response.Result = true
			Response.Error = ""
			Response.Token = Token
		} else {
			Response.Error = "username or password is wrong"
		}
	}

	jsonData, err := json.Marshal(Response)
	if err != nil {
		log.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(string(jsonData)))
}

func Authenticated(token string) (bool, int) {
	var id int
	engine := Controler.GetEngine()
	has, err := engine.Table(Struct.User{}).Where(Struct.User{Token: token}).Cols("id").Get(&id)
	if has && err == nil && id > 0 && token != ""{
		return true, id
	}
	return false, 0
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//TODO update token
}
