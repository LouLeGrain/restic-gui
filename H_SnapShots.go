package main

import (
	"bufio"
	"encoding/json"
	"github.com/gorilla/mux"
	"os/exec"
	"simbookee/restic-gui/models"
	"simbookee/restic-gui/utils"
	"strconv"
	"strings"

	"net/http"
)

func SnapShotsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}

	v := mux.Vars(r)
	id, _ := strconv.Atoi(v["backup_id"])

	credentials, err := models.GetBackupDetails(id)
	utils.Check(err, "fatal")
	if err != nil {
		response.Status = 403
		response.Data = "Bad request"
		json.NewEncoder(w).Encode(response)
		return
	}

	utils.SetEnvVars(credentials)

	opt := Opt{"path": credentials["source"]}
	snapshots, err := GetSnapshots(opt)
	utils.Check(err, "fatal")
	response.Data = snapshots
	json.NewEncoder(w).Encode(response)
}

func GetSnapshots(opt map[string]string) (Rows, error) {
	rows := Rows{}
	var lines []string
	var cmd = "restic snapshots --path=" + opt["path"]

	out, err := exec.Command("bash", "-c", cmd).Output()

	utils.Check(err, "fatal")

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i, v := range lines {
		if i > 1 && i < len(lines)-2 {
			row := Row{}
			fields := strings.Fields(v)
			row.Id = fields[0]
			row.DateTime = fields[1] + " at " + fields[2]
			row.Host = fields[3]
			row.Path = fields[4]
			rows = append(rows, row)
		}
	}
	for left, right := 0, len(rows)-1; left < right; left, right = left+1, right-1 {
		rows[left], rows[right] = rows[right], rows[left]
	}

	return rows, err
}