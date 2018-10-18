package Controler

import (
	"../Models"
	"../Models/Struct"
	"github.com/go-xorm/builder"
	"github.com/gorilla/mux"

	//"github.com/gorilla/mux"
	"net/http"
	//"restful/Controler/Admin"
	"strconv"
)

func RegisterRoot(w http.ResponseWriter, r *http.Request) {
	register(w, r, 0)
}

func RegisterNormal(w http.ResponseWriter, r *http.Request) {
	register(w, r, 1)
}

func register(w http.ResponseWriter, r *http.Request, mode int) {
	vars := mux.Vars(r)
	if ok, _, _, isRootAdmin := Authenticated(r); !(ok && isRootAdmin && mode == 1) && mode != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		submit := r.PostForm.Get("submit")
		edit, er1 := strconv.Atoi(r.PostForm.Get("edit"))

		result := Models.LoginPageVariables{
			Answer:      "",
			SubmitValue: "Register",
		}

		if submit != "" && (username == "" || password == "") {
			result.Answer = "username or password is empty"
		} else if hasUser(username) && er1 != nil {
			result.Answer = "username has already been taken"
		} else if username != "" && password != "" {
			newUser := Struct.NewUser(username, password, "", "", "", mode, 1) //Type = 1 is for admin
			result.Answer = InsertOrUpdate(&Struct.User{}, newUser, edit, er1)
		}

		var editObject Struct.User
		if vars["id"] != "" {
			GetEngine().Table(Struct.User{}).Where(builder.Eq{"id": vars["id"]}).Get(&editObject)
		}
		result.PreviousValue = editObject

		OpenTemplate(w, r, result, "login.html", Models.HeaderVariables{Title: "Register"})
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
