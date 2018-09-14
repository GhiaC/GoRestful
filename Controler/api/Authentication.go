package api

import (
	"net/http"
	"encoding/json"
	"GoRestful/Models"
	"GoRestful/Controler"
	"log"
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
		has, err := engine.Table("user").Where("username = ? and password = ? ", username, password).Cols("id").Get(&id)
		if has && err == nil && id > 0 {
			//TODO update token
			Response.Result = true
			Response.Error = ""
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

func Authenticated(token string) (bool, int64) {
	var id int64
	engine := Controler.GetEngine()
	has, err := engine.Table("user").Where("token = ?", token).Cols("id").Get(&id)
	if has && err == nil && id > 0 {
		return true, id
	}
	return false, 0
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//TODO update token
}
