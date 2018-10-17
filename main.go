package main

import (
	"./Controler"
	"./Controler/Admin"
	"./Controler/api"
	"./Models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"

	//"github.com/nytimes/gziphandler"
	"github.com/tkanos/gonfig"
	"time"
)

func main() {
	err := gonfig.GetConf("./conf.json", &Controler.Configuration)
	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("Resource"))
	http.Handle("/Resource/", http.StripPrefix("/Resource/", fs))

	// Only matches if domain is "www.example.com".
	//r.Host("www.example.com")
	// Matches a dynamic subdomain.
	//r.Host("{subdomain:[a-z]+}.domain.com")

	r.HandleFunc("/", HomePage)                            //DONE
	r.HandleFunc("/login", Controler.Login)                //DONE
	r.HandleFunc("/register/root", Controler.RegisterRoot) //DONE

	s1 := r.PathPrefix("/api").Subrouter()
	{
		s1.HandleFunc("/titles", api.Titles)                  //DONE
		s1.HandleFunc("/subtitles", api.AllSubTitles)         //DONE
		s1.HandleFunc("/subtitle/{id:[0-9]+}", api.SubTitles) //DONE
		s1.HandleFunc("/media/{id:[0-9]+}", api.Media)        //DONE
		s1.HandleFunc("/medias", api.AllMedia)                //DONE
		s1.HandleFunc("/submedia/{id:[0-9]+}", api.SubMedia)  //DONE
		s1.HandleFunc("/submedias", api.AllSubMedia)          //DONE
		s1.HandleFunc("/login", api.Login)                    //DONE
		s1.HandleFunc("/aboutus", api.SubTitles)              //DONE
		s1.HandleFunc("/news", api.AllNews)                   //DONE
		s12 := s1.PathPrefix("/ticket").Subrouter()
		//if true
		{
			//s12.HandleFunc("/getMessage/{id:[0-9]+}", api.GetMessage) //TODO
			s12.HandleFunc("/getMessage", api.GetMessage)     //DONE
			s12.HandleFunc("/sendMessage", api.SendMessage)   //DONE
			s12.HandleFunc("/upload/picture", api.UploadFile) //DONE //name = file
		}
	}

	s2 := r.PathPrefix("/admin").Subrouter()
	{
		s2.HandleFunc("/register", Controler.RegisterNormal) //DONE

		s2.HandleFunc("/FirstLayer", Admin.FirstLayer).Methods("POST")   //DONE
		s2.HandleFunc("/FirstLayer", Admin.FirstLayerGet).Methods("GET") //DONE

		s2.HandleFunc("/SecondLayer/{id:[0-9]+}", Admin.SecondLayer).Methods("POST")   //DONE
		s2.HandleFunc("/SecondLayer/{id:[0-9]+}", Admin.SecondLayerGet).Methods("GET") //DONE

		s2.HandleFunc("/Media/{id:[0-9]+}", Admin.Media).Methods("POST")   //DONE
		s2.HandleFunc("/Media/{id:[0-9]+}", Admin.MediaGet).Methods("GET") //DONE

		s2.HandleFunc("/SubMedia/{id:[0-9]+}", Admin.SubMedia).Methods("POST")   //DONE
		s2.HandleFunc("/SubMedia/{id:[0-9]+}", Admin.SubMediaGet).Methods("GET") //DONE

		s2.HandleFunc("/adduser", Admin.AddUser).Methods("POST")   //DONE
		s2.HandleFunc("/adduser", Admin.AddUserGet).Methods("GET") //DONE

		s2.HandleFunc("/users", Admin.StatusOfUsers).Methods("GET") //DONE

		s2.HandleFunc("/admins", Admin.StatusOfAdmins).Methods("GET") //DONE

		s2.HandleFunc("/news", Admin.News).Methods("POST")   //DONE
		s2.HandleFunc("/news", Admin.NewsGet).Methods("GET") //DONE

		s2.HandleFunc("/upload", Admin.UploadPage).Methods("GET") //DONE

		s2.HandleFunc("/upload/picture", Admin.UploadPicture).Methods("POST") //DONE

		s2.HandleFunc("/upload/file", Admin.UploadFile).Methods("POST") //DONE

		s2.HandleFunc("/messages", Admin.Messages).Methods("GET") //DONE

		s2.HandleFunc("/message/answer/{id:[0-9]+}", Admin.Answer).Methods("POST")   //DONE
		s2.HandleFunc("/message/answer/{id:[0-9]+}", Admin.AnswerGet).Methods("GET") //DONE

		s2.HandleFunc("/logout", Controler.Logout) //DONE

		s4 := s2.PathPrefix("/delete").Subrouter()
		{
			s4.HandleFunc("/user/{id:[0-9]+}", Admin.RemoveUser)
			s4.HandleFunc("/FirstLayer/{id:[0-9]+}", Admin.RemoveFirstLayer)
			s4.HandleFunc("/SecondLayer/{pid:[0-9]+}/{id:[0-9]+}", Admin.RemoveSecondLayer)
			s4.HandleFunc("/Media/{pid:[0-9]+}/{id:[0-9]+}", Admin.RemoveMedia)
			s4.HandleFunc("/SubMedia/{pid:[0-9]+}/{id:[0-9]+}", Admin.RemoveSubMedia)
			s4.HandleFunc("/admin/{id:[0-9]+}", Admin.RemoveAdmin)
			s4.HandleFunc("/news/{id:[0-9]+}", Admin.RemoveNews)
			s4.HandleFunc("/file/{id:[0-9]+}", Admin.RemoveFile)
			s4.HandleFunc("/message/{id:[0-9]+}", Admin.RemoveMessages)
		}
	}

	r.HandleFunc("/file/", Controler.HandleClient)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	http.Handle("/", r)

	go func() {
		if err := http.ListenAndServe(":"+Controler.Configuration.Port, nil); err != nil {
			log.Println(err)
		}
	}()

	//Graceful Shutdown
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	//var wait time.Duration
	// Create a deadline to wait for.
	//ctx, cancel := context.WithTimeout(context.Background(), wait)
	//defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	//r.(func() {})

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 - Not Found!\n")
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now() // find the time right now
	HomePageVars := Models.HomePageVariables{ //store the date and time in a struct
		Date:        now.Format("02-01-2006"),
		Time:        now.Format("15:04:05"),
		LoginStatus: "you aren't logged in",
	}
	if ok, username, _, _ := Controler.Authenticated(r); ok {
		HomePageVars.LoginStatus = "" + username + " عزیز" + ", خوش آمدید."
	}

	Controler.OpenTemplate(w, r, HomePageVars, "homepage.html", Models.HeaderVariables{Title: "پنل مدیریت بینا"})
}
