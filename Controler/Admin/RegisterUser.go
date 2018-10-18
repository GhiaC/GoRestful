package Admin

import (
	"../../Controler"
	"../../Models"
	"../../Models/Struct"
	"github.com/go-xorm/builder"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
		edit, er1 := strconv.Atoi(r.PostForm.Get("edit"))

		result := Models.LoginPageVariables{
			Answer:      "",
			SubmitValue: "افزودن کاربر",
		}

		if username == "" || password == "" {
			result.Answer = "username or password is empty"
		} else if hasUser(username) && er1 != nil {
			result.Answer = "username has already been taken"
		} else if username != "" && password != "" {
			newUser := Struct.NewUser(username, password, name, PhoneNumber, IMEI, 2, 1) // type 2 = user
			result.Answer = Controler.InsertOrUpdate(&Struct.User{}, newUser, edit, er1)
		}

		var editObject Struct.User
		if edit > 0 && er1 == nil {
			Controler.GetEngine().Table(Struct.User{}).Where(builder.Eq{"id": edit}).Get(&editObject)
		}
		result.PreviousValue = editObject
		Controler.OpenTemplate(w, r, result, "Register.html", Models.HeaderVariables{Title: "ثبت"})
	}
}

func AddUserGet(w http.ResponseWriter, r *http.Request) {
	if ok, _, _, _ := Controler.Authenticated(r); !ok {
		http.Redirect(w, r, "/", http.StatusForbidden)
	} else {
		vars := mux.Vars(r)
		var editObject Struct.User
		if vars["id"] != "" {
			Controler.GetEngine().Table(Struct.User{}).Where(builder.Eq{"id": vars["id"]}).Get(&editObject)
		}

		result := Models.LoginPageVariables{
			Answer:        "",
			SubmitValue:   "ثبت",
			PreviousValue: editObject,
		}

		Controler.OpenTemplate(w, r, result, "Register.html", Models.HeaderVariables{Title: "افزودن کاربر"})
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
