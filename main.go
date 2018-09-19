package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"GoRestful/Controler"
	"log"
	"GoRestful/Controler/api"
	"github.com/gorilla/mux"
	"GoRestful/Controler/Admin"
	"fmt"
	"github.com/nytimes/gziphandler"
	"github.com/tkanos/gonfig"
	"strings"
	"runtime/debug"
	"os"
)

func main() {
	err := gonfig.GetConf("./conf.json", &Controler.Configuration)
	if err != nil {
		fmt.Println(err)
	}

	fs := http.FileServer(http.Dir("Resource"))
	http.Handle("/Resource/", http.StripPrefix("/Resource/", fs))

	r := mux.NewRouter()

	r.HandleFunc("/", HomePage)                   //DONE
	r.HandleFunc("/login", Controler.Login)       //DONE
	r.HandleFunc("/register", Controler.Register) //DONE

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
		s1.HandleFunc("/aboutus", api.SubTitles)              //TODO
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
		s2.HandleFunc("/FirstLayer", Admin.FirstLayer)               //DONE
		s2.HandleFunc("/SecondLayer/{id:[0-9]+}", Admin.SecondLayer) //DONE
		s2.HandleFunc("/Media/{id:[0-9]+}", Admin.Media)             //DONE
		s2.HandleFunc("/SubMedia/{id:[0-9]+}", Admin.SubMedia)       //DONE
		s2.HandleFunc("/adduser", Admin.AddUser)                     //DONE
		s2.HandleFunc("/users", Admin.StatusOfUsers)                 //DONE
		s2.HandleFunc("/admins", Admin.StatusOfAdmins)               //DONE
		s2.HandleFunc("/news", Admin.News)                           //DONE
		s2.HandleFunc("/upload", Admin.UploadPage)                   //DONE
		s2.HandleFunc("/upload/picture", Admin.UploadPicture)        //DONE
		s2.HandleFunc("/upload/file", Admin.UploadFile)              //DONE
		s2.HandleFunc("/messages", Admin.Messages)                   //DONE //TODO
		s2.HandleFunc("/message/answer/{id:[0-9]+}", Admin.Answer)   //TODO
		s2.HandleFunc("/logout", Controler.Logout)                   //DONE
	}

	r.HandleFunc("/file/", Controler.HandleClient)

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	//http.Handle("/",Controler.Configuration.Port)

	// Now use the logger with your http.Server:
	logger := log.New(debugLogger{}, "", 0)

	server := &http.Server{
		Addr:     ":"+Controler.Configuration.Port,
		Handler:  gziphandler.GzipHandler(r),
		ErrorLog: logger,
	}
	log.Fatal(server.ListenAndServe())

}

type debugLogger struct{}

func (d debugLogger) Write(p []byte) (n int, err error) {
	s := string(p)
	if strings.Contains(s, "multiple response.WriteHeader") {
		debug.PrintStack()
	}
	return os.Stderr.Write(p)
}
