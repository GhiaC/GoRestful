package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"GoRestful/Models/Struct"
	"github.com/go-xorm/builder"
)

func StatusOfAdmins(w http.ResponseWriter, r *http.Request) {
	status(w, r, 1, "status.html", "Admins") //type 1 = admin
}
func StatusOfUsers(w http.ResponseWriter, r *http.Request) {
	status(w, r, 2, "Users.html", "Users") //type 1 = admin
}

func status(w http.ResponseWriter, r *http.Request, Type int, filename, title string) {
	if ok, _, _ := Controler.Authenticated(r); ok {
		var users []Struct.User
		Controler.GetEngine().Table(Struct.User{}).
			Where(builder.Eq{"Type": Type}).
			Cols("id", "username","phone_number","name").
			Find(&users)
		//if err == nil {
		result := Models.StatusPageVariables{Users: users}
		Controler.OpenTemplate(w, r, result, filename, Models.HeaderVariables{Title: title})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
