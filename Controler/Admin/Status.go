package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"GoRestful/Models/Struct"
)

func Status(w http.ResponseWriter, r *http.Request) {
	if ok, _ ,_:= Controler.Authenticated(r); ok {
		var users []Struct.Admin
		Controler.GetEngine().Table("admin").Cols("id", "username").Find(&users)
		//if err == nil {
		result := Models.StatusPageVariables{Users: users}
		Controler.OpenTemplate(w, r, result, "status.html", Models.HeaderVariables{Title: "Admins"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func Users(w http.ResponseWriter, r *http.Request) {
	if ok, _ ,_:= Controler.Authenticated(r); ok {
		var users []Struct.User
		Controler.GetEngine().Table("user").AllCols().Find(&users)
		//if err == nil {
		result := Models.UsersPageVariables{Users: users}
		Controler.OpenTemplate(w, r, result, "Users.html", Models.HeaderVariables{Title: "Users"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
