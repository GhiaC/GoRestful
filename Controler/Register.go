package Controler

import (
	"net/http"
	"GoRestful/Models"
	"GoRestful/Models/Struct"
	"github.com/go-xorm/builder"
)

func Register(w http.ResponseWriter, r *http.Request) {

	if ok, _, _ := Authenticated(r); ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		submit := r.PostForm.Get("submit")

		vars := Models.LoginPageVariables{
			Answer:      "",
			SubmitValue: "Register",
		}

		if submit != "" && (username == "" || password == "") {
			vars.Answer = "username or password is empty"
		} else if hasUser(username) {
			vars.Answer = "username has already been taken"
		} else if username != "" && password != "" {
			engine := GetEngine()
			newUser := Struct.NewUser(username, password, 1) //Type = 1 is for admin
			affected, err := engine.Table(Struct.User{}).Insert(newUser)
			if affected > 0 && err == nil {
				vars.Answer = "Successful. Go to Login Page"
			}
		}
		OpenTemplate(w, r, vars, "login.html", Models.HeaderVariables{Title: "Register"})
	}
}

func hasUser(username string) bool {
	var id int
	engine := GetEngine()
	has, err := engine.Table(Struct.User{}).Where(builder.Eq{"Username": username}).Cols("id").Get(&id)
	if has && err == nil && id > 0 {
		return true
	}
	return false
}
