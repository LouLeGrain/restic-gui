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

type BuDataStruct struct {
	Name       string            `json:"name,omitempty"`
	Repository int               `json:"repoId,omitempty"`
	Source     string            `json:"source,omitempty"`
	Status     int               `json:"status,omitempty"`
	Now        bool              `json:"buNow,omitempty"`
	Data       map[string]string `json:"data,omitempty"`
}

var BuData BuDataStruct

func BackupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}
	v := mux.Vars(r)

	id, _ := strconv.Atoi(v["backup_id"])

	backup, err := RunBackup(id)
	if err != nil {
		response.Status = 403
		response.Data = "Bad request"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = backup
	json.NewEncoder(w).Encode(response)
}

func RunBackup(id int) (bool, error) {
	credentials, err := models.GetBackupDetails(id)
	utils.Check(err, "")
	if err != nil {
		return false, err
	}
	utils.SetEnvVars(credentials)
	opt := Opt{"path": credentials["source"]}

	return NewBackup(opt)
}

func DeleteBackupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}
	v := mux.Vars(r)

	id, _ := strconv.Atoi(v["backup_id"])

	backup, err := RunBackup(id)
	if err != nil {
		response.Status = 403
		response.Data = "Bad request"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = backup
	json.NewEncoder(w).Encode(response)
}

func NewBackup(opt map[string]string) (bool, error) {
	var ret = false
	var l string
	var cmd = "restic backup " + opt["path"]
	out, err := exec.Command("bash", "-c", cmd).Output()
	//utils.Check(err, "fatal")
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		l = scanner.Text()
	}

	fields := strings.Fields(l)
	if len(fields) == 3 && fields[2] == "saved" {
		ret = true
	}

	return ret, err
}

func InitBackupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&BuData)
	if err != nil {
		response.Status = 403
		response.Data = "Bad request"
		json.NewEncoder(w).Encode(response)
		return
	}

	buId, err := ResticBackupInit()
	if BuData.Now {
		_, err := RunBackup(buId)
		if err != nil {
			response.Status = 403
			response.Data = "Bad request"
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	response.Data = map[string]int{"buId": buId}
	json.NewEncoder(w).Encode(response)
}

func ResticBackupInit() (int, error) {
	jsonString, err := json.Marshal(BuData.Data)
	var buData = models.BackupData{BuData.Repository, BuData.Name, BuData.Source, 1, string(jsonString)}
	buId, err := models.AddBackup(buData)
	utils.Check(err, "fatal")

	return buId, err
}
