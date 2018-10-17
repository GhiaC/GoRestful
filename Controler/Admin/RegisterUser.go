package Admin

import (
	"../../Controler"
	"../../Models"
	"../../Models/Struct"
	"github.com/go-xorm/builder"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	if ok, _, _, _ := Controler.Authenticated(r); !ok {
		http.Redirect(w, r, "/", http.StatusForbidden)
	} else {
		r.ParseForm()
		username := r.PostForm.Get("username")
		PhoneNumber := r.PostForm.Get("PhoneNumber")
		name := r.PostForm.Get("name")
		IMEI := r.PostForm.Get("IMEI")
		password := r.PostForm.Get("password")

		vars := Models.LoginPageVariables{
			Answer:      "",
			SubmitValue: "افزودن کاربر",
		}

		if username == "" || password == "" {
			vars.Answer = "username or password is empty"
		} else if hasUser(username) {
			vars.Answer = "username has already been taken"
		} else if username != "" && password != "" {
			engine := Controler.GetEngine()
			newUser := Struct.NewUser(username, password, name, PhoneNumber, IMEI, 2, 1) // type 2 = user
			affected, err := engine.Table(Struct.User{}).Insert(newUser)
			if affected > 0 && err == nil {
				vars.Answer = "Successful. Go to Login Page"
			}
		}
		Controler.OpenTemplate(w, r, vars, "Register.html", Models.HeaderVariables{Title: "افزودن کاربر"})
	}
}

func AddUserGet(w http.ResponseWriter, r *http.Request) {
	if ok, _, _, _ := Controler.Authenticated(r); !ok {
		http.Redirect(w, r, "/", http.StatusForbidden)
	} else {
		vars := Models.LoginPageVariables{
			Answer:      "",
			SubmitValue: "افزودن کاربر",
		}
		Controler.OpenTemplate(w, r, vars, "Register.html", Models.HeaderVariables{Title: "افزودن کاربر"})
	}
}

func hasUser(username string) bool {
	var id int
	engine := Controler.GetEngine()
	has, err := engine.Table(Struct.User{}).Where(builder.Eq{"Username": username}).Cols("id").Get(&id)
	if has && err == nil && id > 0 {
		return true
	}
	return false
}
