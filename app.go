package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"simbookee/restic-gui/models"
	"simbookee/restic-gui/utils"
)

const PORT = "8000"
const DB_TYPE = "sqlite"

type PageData struct {
	Title   string
	Err     string
	Message string
	Repos   models.Repos
	Backups models.Backups
	Data    string
}

type JsonResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

var PassFile string
var Destination string

func main() {
	_, err := models.GetDb(DB_TYPE)
	utils.Check(err, "fatal")

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/runtime/").Handler(http.StripPrefix("/runtime/", http.FileServer(http.Dir("./runtime/"))))
	r.HandleFunc("/test", TestHandler).Methods("GET")

	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/api/backup/{backup_id}/snapshots", SnapShotsHandler).Methods("GET")
	r.HandleFunc("/api/snapshot/{snapshot_id}/files/", SnapShotHandler).Methods("GET")

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

//todo
//restic -r /Users/andi/backup/test init
//restic -r /tmp/backup backup ~/work
//restic -r /Users/andi/backup/test snapshots
//restic -r /tmp/backup snapshots --path="/srv" (--host luigi , )
//restic -r /tmp/backup diff 5845b002 2ab627a6
//restic -r /tmp/backup restore 79766175 --target /tmp/restore-work
//cat mypassword > repo_pwd.txt
//restic -r /Volumes/restic/jussi -p repo_pwd.txt backup --exclude-file exclude.txt ~/Music/GarageBand
//restic -r /tmp/backup backup --tag projectX --tag foo --tag bar ~/work
//restic -r /tmp/backup restore latest --target /tmp/restore-art --path "/home/art" --host luigi
//restic -r /tmp/backup restore 79766175 --target /tmp/restore-work --include /work/foo
//restic -r /tmp/backup mount /mnt/restic
//restic -r /tmp/backup dump latest production.sql | mysql
//restic -r /tmp/backup forget bdbd3439
//restic -r /tmp/backup prune
//restic forget --keep-last 1 --prune
//restic forget --tag foo --keep-last 1
//restic forget --tag foo --tag bar --keep-last 1
//restic forget --tag foo,tag bar --keep-last 1
//forget --keep-daily 7 --keep-weekly 5 --keep-monthly 12 --keep-yearly 75
//restic -r /tmp/backup key list
