package main

import (
	"encoding/json"
	"net/http"
	"restic-gui/models"
	"restic-gui/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	pageData := PageData{Title: "Restic Backup"}

	repositories, err := models.GetRepositories()
	utils.Check(err, "fatal")
	jsonRepos, err := json.Marshal(repositories)
	utils.Check(err, "fatal")

	backups, err := models.GetBackups()
	utils.Check(err, "fatal")
	jsonBus, err := json.Marshal(backups)
	utils.Check(err, "fatal")

	_, err = utils.CheckProgExists("restic")
	if err != nil {
		pageData.Err = "System Error."
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
