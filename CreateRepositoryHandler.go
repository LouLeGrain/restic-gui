package main

import (
	"encoding/json"
	"net/http"
	"simbookee/restic-gui/models"
)

func CreateRepositoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}

	repositories, err := models.GetRepositories()
	if err != nil {
		response.Status = 403
		response.Data = "Bad request"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = repositories
	json.NewEncoder(w).Encode(response)
}
