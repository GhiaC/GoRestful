package Controler

import (
	"net/http"
	"GoRestful/Models"
	"GoRestful/Models/Struct"
)

func Authenticated(r *http.Request) (bool, string, int) {
	session, _ := Store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		return false, "", 0
	}
	return true, session.Values["username"].(string), session.Values["id"].(int)
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
		var id int
		engine := GetEngine()
		has, err := engine.Table(Struct.User{}).Where(Struct.User{Username:username,Password:password}).Cols("id").Get(&id)
		if has && err == nil && id > 0 {
			session, _ := Store.Get(r, "cookie-name")
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Values["id"] = id
			session.Save(r, w)
		} else {
			vars.Answer = "username or password is wrong"
		}
	}

	if ok, _, _ := Authenticated(r); !ok {
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
