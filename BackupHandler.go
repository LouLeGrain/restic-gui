package main

import (
	"bufio"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os/exec"
	"simbookee/restic-gui/models"
	"simbookee/restic-gui/utils"
	"strconv"
	"strings"
)

func CreateBackupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}

	json.NewEncoder(w).Encode(response)
}


func BackupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}

	v := mux.Vars(r)
	id, _ := strconv.Atoi(v["backup_id"])

	credentials, err := models.GetBackupDetails(id)
	utils.Check(err, "")
	if err != nil {
		response.Status = 403
		response.Data = "Bad request"
		json.NewEncoder(w).Encode(response)
		return
	}

	utils.SetEnvVars(credentials)

	opt := Opt{"path": credentials["source"]}
	backup, err := NewBackup(opt)
	response.Data = backup
	json.NewEncoder(w).Encode(response)

}

func NewBackup(opt map[string]string) (bool, error) {
	var ret = false
	var l string
	var cmd = "restic backup " + opt["path"]
	out, err := exec.Command("bash", "-c", cmd).Output()
	utils.Check(err, "")
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		l = scanner.Text()
	}

	fields := strings.Fields(l)
	if fields[2] == "saved" {
		ret = true
	}

	return ret, err
}
