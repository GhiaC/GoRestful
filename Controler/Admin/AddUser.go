package Admin

import (
	"net/http"
	"GoRestful/Models"
	"GoRestful/Controler"
	"GoRestful/Models/Struct"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	if ok, _ ,_:= Controler.Authenticated(r); !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		submit := r.PostForm.Get("submit")

		vars := Models.LoginPageVariables{
			Answer: "",
			//Url:         "/register",
			SubmitValue: "Add User",
		}

		if submit != "" && (username == "" || password == "") {
			vars.Answer = "username or password is empty"
		} else if hasUser(username) {
			vars.Answer = "username has already been taken"
		} else if username != "" && password != "" {
			engine := Controler.GetEngine()
			newUser := Struct.NewUser(username, password)
			affected, err := engine.Table("user").Insert(newUser)
			println(affected)
			if affected > 0 && err == nil {
				vars.Answer = "Successful. Go to Login Page"
			}
		}
		Controler.OpenTemplate(w, r, vars, "login.html", Models.HeaderVariables{Title: "Add User"})
	}
}

func hasUser(username string) bool {
	var id int
	engine := Controler.GetEngine()
	has, err := engine.Table("user").Where("username = ?", username).Cols("id").Get(&id)
	if has && err == nil && id > 0 {
		return true
	}
	return false
}
