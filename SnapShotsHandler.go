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

type Row struct {
	Id   string `json:"id"`
	Date string `json:"date"`
	Time string `json:"time"`
	Host string `json:"host"`
	Path string `json:"path"`
}

type Rows []Row
type Opt map[string]string

func SnapShotsHandler(w http.ResponseWriter, r *http.Request) {

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

	fname := "./runtime/" + utils.GetMD5Hash(credentials["source"])
	PassFile, err = utils.SetPassFile(fname, credentials["passwd"])
	Destination = credentials["destination"]
	utils.Check(err, "")
	opt := Opt{"path": credentials["source"]}
	snapshots, err := GetSnapshots(opt)
	utils.Check(err, "")

	response.Data = snapshots
	json.NewEncoder(w).Encode(response)
}

func GetSnapshots(opt map[string]string) (Rows, error) {
	rows := Rows{}
	var lines []string
	var cmd = "restic -r " + Destination + " -p " + PassFile + " snapshots --path=" + opt["path"]

	out, err := exec.Command("bash", "-c", cmd).Output()
	utils.Check(err, "")

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i, v := range lines {
		if i > 1 && i < len(lines)-2 {
			row := Row{}
			fields := strings.Fields(v)
			row.Id = fields[0]
			row.Date = fields[1]
			row.Time = fields[2]
			row.Host = fields[3]
			row.Path = fields[4]
			rows = append(rows, row)
		}
	}

	return rows, err
}
