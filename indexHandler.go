package main

import (
	"encoding/json"
	"net/http"
	"simbookee/restic-gui/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageData := PageData{Title: "Restic Backup"}
	backups, err := models.GetBackups()
	jsonBus, err1 := json.Marshal(backups)

	if err1 != nil || err != nil {
		pageData.Err = "System Error. Please contact support at support@simbookee.com"
	}
	pageData.Backups = backups
	pageData.Data = "{\"backups\":" + string(jsonBus) + "}"

	if pageData.Err != "" {
		render(w, "error.html", pageData, "layout")
	} else {
		render(w, "index.html", pageData, "layout")
	}
}
