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

const PORT = "8008"

const DB_TYPE = "sqlite"

func main() {
	//init bd
	_, err := models.GetDb(DB_TYPE)
	utils.Check(err, "fatal")

	//setup router
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/runtime/").Handler(http.StripPrefix("/runtime/", http.FileServer(http.Dir("./runtime/"))))

	//routes
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/api/repositories", RepositoryHandler).Methods("GET")
	r.HandleFunc("/api/repositories/new", CreateRepositoryHandler).Methods("POST")
	r.HandleFunc("/api/newbackup", CreateBackupHandler).Methods("POST")
	r.HandleFunc("/api/snapshots/{backup_id}", SnapShotsHandler).Methods("GET")
	r.HandleFunc("/api/snapshots/new/{backup_id}", BackupHandler).Methods("GET")
	r.HandleFunc("/api/snapshots/forget/{backup_id}/{snapshot_id}", ForgetHandler).Methods("GET")
	r.HandleFunc("/api/snapshots/prune/{backup_id}", PruneHandler).Methods("GET")
	r.HandleFunc("/api/files/{backup_id}/{snapshot_id}", FilesHandler).Methods("GET")

	//test route
	r.HandleFunc("/test", TestHandler).Methods("GET")
	http.Handle("/", r)

	utils.OpenBrowser("http://localhost" + getPort() + "/")
	//run app
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
	p := os.Getenv("RESTIC_PORT")
	if p != "" {
		return ":" + p
	}
	return ":" + PORT
}

/*todo
restic init
restic backup --exclude-file exclude.txt ~/Music/GarageBand

restic snapshots
restic snapshots --path --host --tag )

restic diff 5845b002 2ab627a6

restic restore 79766175 --target /tmp/restore-work
restic restore latest --target /tmp/restore-art --path "/home/art" --host luigi
restic restore 79766175 --target /tmp/restore-work --include /work/foo

restic mount /mnt/restic
restic dump latest production.sql | mysql

restic forget bdbd3439
restic forget --keep-last 1 --prune
restic forget --tag foo --keep-last 1
restic forget --tag foo --tag bar --keep-last 1
restic forget --tag foo,tag bar --keep-last 1
restic rorget --keep-daily 7 --keep-weekly 5 --keep-monthly 12 --keep-yearly 75
restic prune

restic key list
*/
