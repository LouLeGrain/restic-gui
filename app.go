package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"simbookee/restic-gui/models"
)

const PORT = "8000"
const DB_TYPE = "sqlite"

type PageData struct {
	Title   string
	Err     string
	Message string
	Backups models.Backups
	Data    string
}

type JsonResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func main() {
	_, err := models.GetDb(DB_TYPE)
	checkErr(err)

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/", IndexHandler).Methods("GET")
	//r.HandleFunc("/api/check/path", ApiHandler).Methods("POST")
	r.HandleFunc("/test", TestHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(getPort(), r))
}

func render(w http.ResponseWriter, tmpl string, p PageData, l string) {
	if l == "" {
		l = "layout"
	}

	layout := "templates/" + l + ".html"

	tmpl = fmt.Sprintf("templates/%s", tmpl)    // prefix the name passed in with templates/
	t, err := template.ParseFiles(tmpl, layout) //parse the template file held in the templates folder

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if l != "" {
		//execute the template and pass in the variables to fill the gaps
		err = t.ExecuteTemplate(w, l, p)
	}

	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":" + PORT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
