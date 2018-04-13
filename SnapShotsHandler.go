package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"os/exec"
	"simbookee/restic-gui/models"
	"simbookee/restic-gui/utils"
	"strconv"

	"net/http"
)

func SnapShotsHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, _ := strconv.Atoi(v["backup_id"])

	status := 200

	data, err := models.GetBackupDetails(id)
	utils.CheckErr(err, "")
	if err != nil {
		status = 403
	}
	snapshots, err := GetSnapshots(data)
	response := JsonResponse{status, snapshots}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetSnapshots(data map[string]string) (interface{}, error) {

	fmt.Println(data)

	out, err := exec.Command("date").Output()

	utils.CheckErr(err, "")

	fmt.Printf("The date is %s\n", out)

	return data, nil
}
