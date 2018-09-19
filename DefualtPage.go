package main

import (
	"net/http"
	"fmt"
	"time"
	"GoRestful/Models"
	"GoRestful/Controler"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 - Not Found!\n")
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now() // find the time right now
	HomePageVars := Models.HomePageVariables{//store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
		LoginStatus: "you aren't logged in",
	}
	if ok, username, _ := Controler.Authenticated(r); ok {
		HomePageVars.LoginStatus = "dear " + username + ", you are logged in"
	}

	Controler.OpenTemplate(w, r, HomePageVars, "homepage.html", Models.HeaderVariables{Title: "Authentication GO"})
}

