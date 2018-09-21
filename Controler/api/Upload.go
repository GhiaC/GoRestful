package api

import (
	"net/http"
	"mime/multipart"
	"GoRestful/Models"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
	"GoRestful/Controler"
	"GoRestful/Models/Struct"
	"strconv"
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
		jsonResponse(w, http.StatusCreated, &response)
		return
	}
	file, handle, err := r.FormFile("file")
	description := "User uploaded"

	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	//mimeType := handle.Header.Get("Content-Type")
	//switch mimeType {
	//case "image/jpeg":
	//case "image/png":
	//	saveFile(w, file, handle)
	//default:
	//	jsonResponse(w, http.StatusBadRequest, Models.UploadFileResponse{
	//		Error:    "The format file is not valid.",
	//		Result:   false,
	//		FileName: "",
	//	})
	//}

	fileKey = Controler.TokenGenerator()

	saveFile(w, file, handle, &response)

	engine := Controler.GetEngine()
	newFile := Struct.NewFile(id, handle.Filename, fileKey, description)
	engine.Table(Struct.Picture{}).Insert(newFile) //has result
	response.FileName = fileKey
	jsonResponse(w, http.StatusCreated, &response)
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
