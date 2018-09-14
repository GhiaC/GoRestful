package Admin

import (
	"net/http"
	"GoRestful/Controler"
	"GoRestful/Models"
	"GoRestful/Models/Struct"
)

func UploadPage(w http.ResponseWriter, r *http.Request) {
	if ok, _, _ := Controler.Authenticated(r); ok {

		var files []Struct.AdminFile
		Controler.GetEngine().Table(Struct.AdminFile{}).AllCols().Find(&files)

		result := Models.AdminFileLayerVariables{
			Files: files,
		}

		Controler.OpenTemplate(w, r, result, "UploadFile.html", Models.HeaderVariables{Title: "Upload"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
