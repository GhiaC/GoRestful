package Controler

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"GoRestful/Models"
	"encoding/json"
	"log"
	"GoRestful/Models/Struct"
)

func UploadToUserPicture(w http.ResponseWriter, r *http.Request) {
	uploadFile(w, r, Struct.UserPicture{})
}

func UploadToAdminPicture(w http.ResponseWriter, r *http.Request) {
	uploadFile(w, r, Struct.AdminPicture{})
}

func UploadToAdminFile(w http.ResponseWriter, r *http.Request) {
	uploadFile(w, r, Struct.AdminFile{})
}

//UploadFile uploads a file to the server
func uploadFile(w http.ResponseWriter, r *http.Request, table interface{}) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	//mimeType := handle.Header.Get("Content-Type")
	//switch mimeType {
	//case "image/jpeg":

	//if submit != "" && (text == "") {
	//	result.Answer = "text is empty"
	//} else if text != "" {

	response := Models.UploadFileResponse{
		Error:    "",
		Result:   false,
		FileName: "",
	}
	saveFile(w, file, handle, &response)

	if auth, _, id := Authenticated(r); auth && id > 0 {
		engine := GetEngine()
		//id, _ := strconv.Atoi(vars["id"])
		newFile := Struct.NewAdminFile(id, handle.Filename)
		affected, _ := engine.Table(table).Insert(newFile)
		println(affected)
		response.FileName = "/file/admin/picture/" + string(newFile.Id)
		jsonResponse(w, http.StatusCreated, &response)
	}
	//if affected > 0 && err == nil {
	//	result.Answer = "Successful."
	//}
	//}

	//case "image/png":
	//	saveFile(w, file, handle)
	//default:
	//	jsonResponse(w, http.StatusBadRequest, Models.UploadFileResponse{
	//		Error:    "The format file is not valid.",
	//		Result:   false,
	//		FileName: "",
	//	})
	//}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader, response *Models.UploadFileResponse) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	response.Error = "File uploaded successfully!."
	response.Result = true
	//response.FileName = "/files/" + handle.Filename
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
	//fmt.Fprint(w, message)
}
