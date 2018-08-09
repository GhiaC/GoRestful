package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"restful/Controler"
	"restful/Models"
	"log"
	"restful/Controler/api"
	"github.com/gorilla/mux"
)

func main() {
	fs := http.FileServer(http.Dir("Resource"))
	http.Handle("/Resource/", http.StripPrefix("/Resource/", fs))

	r := mux.NewRouter()

	r.HandleFunc("/", HomePage)
	r.HandleFunc("/login", Controler.Login)
	r.HandleFunc("/register", Controler.Register)
	r.HandleFunc("/users", Controler.Status)
	r.HandleFunc("/logout", Controler.Logout)
	r.HandleFunc("/admin", Controler.Admin)
	r.HandleFunc("/admin/FirstLayer", Controler.FirstLayer)

	r.HandleFunc("/api/titles", api.Titles)
	r.HandleFunc("/api/subtitle/{id:[0-9]+}", api.SubTitles)
	r.HandleFunc("/api/media/{id:[0-9]+}", api.Media)
	r.HandleFunc("/SecondLayer/{id:[0-9]+}", Controler.SecondLayer)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}


func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now() // find the time right now
	HomePageVars := Models.HomePageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
		LoginStatus: "you aren't logged in",
	}
	if ok, username := Controler.Authenticated(r); ok {
		HomePageVars.LoginStatus = "dear " + username + ", you are logged in"
	}

	Controler.OpenTemplate(w,r,HomePageVars,"homepage.html",Models.HeaderVariables{Title:"Authentication GO"})
}