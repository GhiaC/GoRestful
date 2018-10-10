package Controler

import (
	"../Models"
	"../Models/Struct"
	"github.com/go-xorm/builder"
	"net/http"
)

func Authenticated(r *http.Request) (bool, string, int, bool) {
	session, _ := Store.Get(r, "cookie-name")

	_, ok2 := session.Values["type"].(int)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth || !ok2 {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		return false, "", 0, false
	}
	return true, session.Values["username"].(string), session.Values["id"].(int), session.Values["type"].(int) == 0
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	submit := r.PostForm.Get("submit")

	vars := Models.LoginPageVariables{
		Answer:      "",
		SubmitValue: "Login",
	}

	if submit == "Login" && (username == "" || password == "") {
		vars.Answer = "username or password is empty"
	} else if username != "" && password != "" {
		var user Struct.User
		engine := GetEngine()
		has, err := engine.Table(Struct.User{}).Where(builder.Eq{"Username": username, "Password": password}).
			Cols("id", "type").Get(&user)
		if has && err == nil && user.Id > 0 && user.Type != 2 {
			session, _ := Store.Get(r, "cookie-name")
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Values["id"] = user.Id
			session.Values["type"] = user.Type
			session.Save(r, w)
		} else {
			vars.Answer = "username or password is wrong"
		}
	}

	if ok, _, _, _ := Authenticated(r); !ok {
		OpenTemplate(w, r, vars, "login.html", Models.HeaderVariables{Title: "Login"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Values["username"] = "empty"
	session.Values["id"] = 0
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

//var messages []Models.AnswerQuery
//Controler.GetEngine().Table(Struct.Message{}).
//Select("user.username,message.*,file.type").
//Join("INNER", Struct.User{}, "message.user_id = user.id ").
//Join("LEFT", Struct.File{}, "message.file_address = file.key").
//Where(builder.Eq{"user_id": userid}).Or(builder.Eq{"answer_to": userid}).OrderBy("message.created").
//Find(&messages)

func GetUserPhonenumber(message_id int) string {
	var phonenumber string
	engine := GetEngine()
	has, err := engine.Table(Struct.User{}).
		Select("user.phone_number").
		Join("INNER", Struct.Message{}, "message.user_id = user.id").
		Where(builder.Eq{"message.id": message_id}).
		Get(&phonenumber)
	if has && err == nil && phonenumber != "" {
		return phonenumber
	} else {
		println("error in find user phonenumber")
		return "Error"
	}
}
