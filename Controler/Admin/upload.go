package Admin

import (
	"net/http"
	"../../Controler"
	"../../Models"
	"../../Models/Struct"
	"mime/multipart"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
	"strconv"
	"github.com/go-xorm/builder"
)

func UploadPage(w http.ResponseWriter, r *http.Request) {
	if ok, _, _ ,_:= Controler.Authenticated(r); ok {

		var files []Models.FileInner
		Controler.GetEngine().Table(Struct.File{}).Select("file.*,user.username").
			Join("INNER", Struct.User{}, "file.user_id = user.id ").
			Where(builder.Eq{"admin_file": true}).Find(&files)

		result := Models.AdminFileLayerVariables{
			Files: files,
		}

		Controler.OpenTemplate(w, r, result, "UploadFile.html", Models.HeaderVariables{Title: "Upload"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func UploadPicture(w http.ResponseWriter, r *http.Request) {
	uploadFile(w, r, Struct.File{}, false)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	uploadFile(w, r, Struct.File{}, true)
}

var fileKey = ""

//UploadFile uploads a file to the server
func uploadFile(w http.ResponseWriter, r *http.Request, table interface{}, AdminFile bool) {
	mem, _ := strconv.Atoi(r.Header["Content-Length"][0])
	r.ParseMultipartForm(int64(mem))
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	response := Models.UploadFileResponse{
		Error:    "",
		Result:   false,
		FileName: "",
	}
	logged, _, id ,_:= Controler.Authenticated(r)
	if !logged && !(id > 0) {
		response.Error = "Access denied"
		jsonResponse(w, http.StatusOK, &response)
		return
	}
	file, handle, err := r.FormFile("file")
	description := r.FormValue("description")

	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	//switch mimeType {
	//case "image/jpeg":
	fileKey = Controler.TokenGenerator() //and filename
	if saveFile(w, file, handle, &response, fileKey) {
		insertFileInfo(id, fileKey, description, mimeType, table, AdminFile)
		response.Result = true
		response.FileName = fileKey
	}
	//case "image/png":
	//	if saveFile(w, file, handle, &response) {
	//		FileName := insertFileInfo(id, handle.Filename, description, mimeType, table)
	//		response.Result = true
	//		response.FileName = FileName
	//	}
	//default:
	//	response.Error = "The format file is not valid."
	//	jsonResponse(w, http.StatusBadRequest, &response)
	//	return
	//}
	jsonResponse(w, http.StatusOK, &response)

}

func insertFileInfo(userId int, filename, description, Type string, table interface{}, AdminFile bool) {
	engine := Controler.GetEngine()
	newFile := Struct.NewFile(userId, filename, filename, description, Type, AdminFile)
	engine.Table(table).Insert(newFile) //has result
}

func saveFile(
	w http.ResponseWriter,
	file multipart.File,
	handle *multipart.FileHeader,
	response *Models.UploadFileResponse,
	filename string) bool {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return false
	}

	err = ioutil.WriteFile("./files/"+filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return false
	}
	response.Error = "File uploaded successfully!."
	response.Result = true
	return true
}

func jsonResponse(w http.ResponseWriter, code int, message *Models.UploadFileResponse) {
	var jsonData []byte
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(jsonData)))
}
