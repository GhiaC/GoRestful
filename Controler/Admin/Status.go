package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
)

func Status(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Controler.Authenticated(r); ok {
		var users []Models.Admin
		Controler.GetEngine().Table("admin").Cols("id", "username").Find(&users)
		//if err == nil {
		result := Models.StatusPageVariables{Users: users}
		Controler.OpenTemplate(w, r, result, "status.html", Models.HeaderVariables{Title: "Users"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
