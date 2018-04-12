package main

import (
	"encoding/json"
	"net/http"
	"simbookee/restic-gui/models"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageData := PageData{Title: "Restic Backup"}
	repositories, err1 := models.GetRepositories()
	backups, err2 := models.GetBackups()
	jsonRepos, err3 := json.Marshal(repositories)
	jsonBus, err4 := json.Marshal(backups)
	repos, err := MergeBackup(repositories, backups)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err != nil {
		pageData.Err = "System Error. Please contact support at support@simbookee.com"
	}
	pageData.Repos = repos
	pageData.Backups = backups
	pageData.Data = "{\"repositories\":" + string(jsonRepos) + ",\"backups\":" + string(jsonBus) + "}"
	if pageData.Err != "" {
		render(w, "error.html", pageData, "layout")
	} else {
		render(w, "index.html", pageData, "layout")
	}
}

func MergeBackup(repos models.Repos, bu models.Backups) (models.Repos, error) {
	for k, rep := range repos {
		var backups = models.Backups{}
		for _, v := range rep.BackupIds {
			i, _ := strconv.Atoi(v)
			backups[i] = bu[i]
		}
		rep.Backups = backups
		repos[k] = rep
	}

	return repos, nil
}
