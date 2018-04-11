package main

import (
	"encoding/json"
	"net/http"
	"simbookee/restic-gui/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	pageData := PageData{Title: "Restic Backup"}
	repositories, err1 := models.GetRepositories()
	backups, err2 := models.GetBackups()
	jsonRepos, err3 := json.Marshal(repositories)
	jsonBus, err4 := json.Marshal(backups)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		pageData.Err = "System Error. Please contact support at support@simbookee.com"
	}

	pageData.Repos = mergeBackup(backups, backups)

	pageData.Backups = backups

	pageData.Data = "{\"repositories\":" + string(jsonRepos) + ",\"backups\":" + string(jsonBus) + "}"

	if pageData.Err != "" {
		render(w, "error.html", pageData, "layout")
	} else {
		render(w, "index.html", pageData, "layout")
	}
}

func mergeBackup(rep Repositories, bu Backups) (Repositories, error) {

	return rep, nil

}
