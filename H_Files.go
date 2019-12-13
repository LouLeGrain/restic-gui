package main

import (
	"bufio"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/exec"
	"rest/models"
	"rest/utils"
	"strconv"
	"strings"
)

type File struct {
	Id     string `json:"id"`
	Parent string `json:"parent"`
	Text   string `json:"text"`
	Link   string `json:"li_attr"`
}

type SnFiles []File

func FilesHandler(w http.ResponseWriter, r *http.Request) {
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
	opt := Opt{"id": snid, "path": credentials["source"]}
	files, err := GetRestoreFiles(opt)
	utils.Check(err, "fatal")

	snfiles, err := BuildFilesData(files)
	utils.Check(err, "fatal")

	response.Data = snfiles
	json.NewEncoder(w).Encode(response)
}

func GetRestoreFiles(opt map[string]string) (Files, error) {
	filerows := []string{}
	var cmd = "restic -r " + os.Getenv("RESTIC_REPOSITORY") + " ls " + opt["id"]
	out, err := exec.Command("bash", "-c", cmd).Output()
	utils.Check(err, "fatal")
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		filerows = append(filerows, scanner.Text())
	}

	return filerows[1:], nil
}

func BuildFilesData(f []string) (SnFiles, error) {
	snfiles := SnFiles{}
	for i, v := range f {
		var fobj = File{}
		fileSlice := strings.SplitAfter(v, "/")

		if len(fileSlice) <= 2 {
			fobj = File{"key_" + strconv.Itoa(i), "#", fileSlice[1], "{\"path\":\"" + v + "\"}"}
		} else {
			path := strings.Replace(strings.Join(fileSlice[:len(fileSlice)-1], "/"), "//", "/", len(fileSlice)-1)
			idx := utils.SliceIndex(f, path[:len(path)-1])
			fobj = File{"key_" + strconv.Itoa(i), "key_" + strconv.Itoa(idx), fileSlice[len(fileSlice)-1], "{\"path\":\"" + v + "\"}"}
		}
		snfiles = append(snfiles, fobj)
	}

	return snfiles, nil
}
