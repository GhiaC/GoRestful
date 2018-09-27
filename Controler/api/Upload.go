package api

import (
	"net/http"
	"mime/multipart"
	"../../Models"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
	"../../Controler"
	"strconv"
	"../../Models/Struct"
)

var fileKey = Controler.TokenGenerator()

//UploadFile uploads a file to the server
func UploadFile(w http.ResponseWriter, r *http.Request) {
	mem, _ := strconv.Atoi(r.Header["Content-Length"][0])
	r.ParseMultipartForm(int64(mem))
	Token := r.FormValue("Token")
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	response := Models.UploadFileResponse{
		Error:    "",
		Result:   false,
		FileName: "",
	}
	logged, id := Authenticated(Token)
	if !logged && !(id > 0) {
		response.Error = "Access denied"
		jsonResponse(w, http.StatusOK, &response)
		return
	}
	file, handle, err := r.FormFile("file")
	description := "User uploaded"

	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	//switch mimeType {
	//case "image/jpeg":
	//	if saveFile(w, file, handle, &response) {
	//		FileName := insertFileInfo(id, handle.Filename, description, mimeType)
	//		response.Result = true
	//		response.FileName = FileName
	//	}
	//case "image/png":
	fileKey = Controler.TokenGenerator() //and filename

	if saveFile(w, file, handle, &response, fileKey) {
		insertFileInfo(id, fileKey, description, mimeType)
		response.Result = true
		response.FileName = fileKey
	}
	//default:
	//	response.Error = "The format file is not valid."
	//	jsonResponse(w, http.StatusBadRequest, &response)
	//	return
	//}
	jsonResponse(w, http.StatusOK, &response)

}

func insertFileInfo(userid int, filename, description, Type string) {
	engine := Controler.GetEngine()
	newFile := Struct.NewFile(userid, filename, filename, description, Type,false)
	engine.Table(Struct.File{}).Insert(newFile) //has result
}

func saveFile(
	w http.ResponseWriter,
	file multipart.File,
	handle *multipart.FileHeader,
	response *Models.UploadFileResponse,
	fileKey string) bool {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return false
	}

	err = ioutil.WriteFile("./files/"+fileKey, data, 0666)
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
