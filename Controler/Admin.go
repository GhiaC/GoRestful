package Controler

import (
	"net/http"

	"restful/Models"
	"github.com/gorilla/mux"
	"strconv"
)

func Admin(w http.ResponseWriter, r *http.Request) {
}

func Status(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Authenticated(r); ok {
		var users []Models.User
		GetEngine().Table("user").Cols("id", "username").Find(&users)
		//if err == nil {
		result := Models.StatusPageVariables{Users: users}
		OpenTemplate(w, r, result, "status.html", Models.HeaderVariables{Title: "Users"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func FirstLayer(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		submit := r.PostForm.Get("submit")

		vars := Models.FirstLayerVariables{
			Answer:      "",
			Url:         "/register",
			SubmitValue: "Add Title",
		}

		if submit != "" && (username == "") {
			vars.Answer = "text is empty"
		} else if username != "" {
			engine := GetEngine()
			newUser := Models.NewTitle(username)
			affected, err := engine.Table("title").Insert(newUser)
			println(affected)
			if affected > 0 && err == nil {
				vars.Answer = "Successful."
			}
		}

		var titles []Models.Title
		GetEngine().Table("title").Cols("Id", "Title").Find(&titles)
		vars.Titles = titles
		OpenTemplate(w, r, vars, "FirstLayer.html", Models.HeaderVariables{Title: "FirstLayer"})

	} else if ok, _ := Authenticated(r); ok {

		var titles []Models.Title
		GetEngine().Table("title").Cols("Id", "Title").Find(&titles)

		result := Models.FirstLayerVariables{
			Titles:      titles,
			Answer:      "",
			Url:         "/register",
			SubmitValue: "Add Title",}

		OpenTemplate(w, r, result, "FirstLayer.html", Models.HeaderVariables{Title: "FirstLayer"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func SecondLayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if ok, _ := Authenticated(r); ok && r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		submit := r.PostForm.Get("submit")

		result := Models.SecondLayerVariables{
			Answer:      "",
			Url:         "/register",
			SubmitValue: "Add Subtitle",
		}

		if submit != "" && (username == "") {
			result.Answer = "text is empty"
		} else if username != "" {
			engine := GetEngine()
			id, _ := strconv.Atoi(vars["id"])
			newUser := Models.NewSubtitle(int64(id), username)
			affected, err := engine.Table("subtitle").Insert(newUser)
			println(affected)
			if affected > 0 && err == nil {
				result.Answer = "Successful."
			}
		}

		var subtitles []Models.Subtitle
		GetEngine().Table("subtitle").Cols("Id", "Titleid", "Title").Where("Titleid = ?", vars["id"]).Find(&subtitles)

		result.Subtitles = subtitles
		OpenTemplate(w, r, result, "SecondLayer.html", Models.HeaderVariables{Title: "SecondLayer"})

	} else if ok, _ := Authenticated(r); ok {
		var subtitles []Models.Subtitle
		GetEngine().Table("subtitle").Cols("Id", "Titleid", "Title").Where("Titleid = ?", vars["id"]).Find(&subtitles)

		result := Models.SecondLayerVariables{
			TitleId:     vars["id"],
			Subtitles:   subtitles,
			Answer:      "",
			Url:         "/register",
			SubmitValue: "Add Subtitle",}

		OpenTemplate(w, r, result, "SecondLayer.html", Models.HeaderVariables{Title: "SecondLayer"})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
