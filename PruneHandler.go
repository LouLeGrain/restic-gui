package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/exec"
	"simbookee/restic-gui/models"
	"simbookee/restic-gui/utils"
	"strconv"
)

func PruneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}
	v := mux.Vars(r)
	snid := v["snapshot_id"]
	buid, _ := strconv.Atoi(v["backup_id"])

	credentials, err := models.GetBackupDetails(buid)
	utils.Check(err, "fatal")

	if err != nil {
		response.Status = 403
		response.Data = "Bad request"
		json.NewEncoder(w).Encode(response)
		return
	}

	utils.SetEnvVars(credentials)
	opt := Opt{"id": snid}

	result, err := PruneBackup(opt)
	response.Data = result
	json.NewEncoder(w).Encode(response)

}

func PruneBackup(opt map[string]string) (string, error) {
	var cmd = "restic -r " + os.Getenv("RESTIC_REPOSITORY") + " prune "
	resp, err := exec.Command("bash", "-c", cmd).Output()
	utils.Check(err, "fatal")
	return string(resp), err
}
