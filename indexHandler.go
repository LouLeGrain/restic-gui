package main

import (
	"encoding/json"
	"net/http"
	"simbookee/restic-gui/models"
	"simbookee/restic-gui/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	pageData := PageData{Title: "Restic Backup"}
	repositories, err := models.GetRepositories()
	utils.CheckErr(err, "")
	backups, err := models.GetBackups()
	utils.CheckErr(err, "")
	jsonRepos, err := json.Marshal(repositories)
	utils.CheckErr(err, "")
	jsonBus, err := json.Marshal(backups)
	utils.CheckErr(err, "")

	if err != nil {
		pageData.Err = "System Error. Please contact support at support@simbookee.com"
	}
	pageData.Repos = repositories
	pageData.Backups = backups
	pageData.Data = "{\"repositories\":" + string(jsonRepos) + ",\"backups\":" + string(jsonBus) + "}"
	if pageData.Err != "" {
		render(w, "error.html", pageData, "layout")
	} else {
		render(w, "index.html", pageData, "layout")
	}
}
