package main

import (
	"bufio"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/exec"
	"simbookee/restic-gui/models"
	"simbookee/restic-gui/utils"
	"strconv"
	"strings"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}

	v := mux.Vars(r)
	snid := v["snapshot_id"]
	buid, _ := strconv.Atoi(v["backup_id"])

	credentials, err := models.GetBackupDetails(buid)
	utils.Check(err, "")
	if err != nil {
		response.Status = 403
		response.Data = "Bad request"
		json.NewEncoder(w).Encode(response)
		return
	}

	utils.SetEnvVars(credentials)
	opt := Opt{"id": snid, "path": credentials["source"]}
	files, err := GetFiles(opt)
	utils.Check(err, "")
	response.Data = files
	json.NewEncoder(w).Encode(response)
}

func GetFiles(opt map[string]string) (Files, error) {
	var lines []string
	var files = Files{}
	var cmd = "restic -r " + os.Getenv("RESTIC_REPOSITORY") + " ls " + opt["id"]
	out, err := exec.Command("bash", "-c", cmd).Output()
	utils.Check(err, "")

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	_, files = lines[0], lines[1:]

	return files, nil
}
