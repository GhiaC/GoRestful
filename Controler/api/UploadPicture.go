package api

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"GoRestful/Models"
	"encoding/json"
	"log"
)

// UploadFile uploads a file to the server
func UploadFile(w http.ResponseWriter, r *http.Request) {
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

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		saveFile(w, file, handle)
	case "image/png":
		saveFile(w, file, handle)
	default:
		jsonResponse(w, http.StatusBadRequest, Models.UploadFileResponse{
			Error:    "The format file is not valid.",
			Result:   false,
			FileName: "",
		})
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
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
	jsonResponse(w, http.StatusCreated, Models.UploadFileResponse{
		Error:    "File uploaded successfully!.",
		Result:   true,
		FileName: "/files/" + handle.Filename,
	})
}

func jsonResponse(w http.ResponseWriter, code int, message Models.UploadFileResponse) {
	var jsonData []byte
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(jsonData)))
	fmt.Fprint(w, message)
}
