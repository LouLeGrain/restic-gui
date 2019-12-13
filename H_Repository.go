package main

import (
	"encoding/json"
	"net/http"
	"os"
	"os/exec"
	"rest/models"
	"rest/utils"
)

type RepoDataStruct struct {
	Password string            `json:"password,omitempty"`
	Name     string            `json:"name,omitempty"`
	Type     string            `json:"type,omitempty"`
	Data     map[string]string `json:"data,omitempty"`
}

var RepoData RepoDataStruct

func RepositoryHandler(w http.ResponseWriter, r *http.Request) {
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

func InitRepositoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := JsonResponse{200, nil}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&RepoData)
	if err != nil {
		response.Status = 403
		response.Data = "Bad request"
		json.NewEncoder(w).Encode(response)
		return
	}

	repoId, err := ResticRepoInit()

	response.Data = map[string]int{"repoId": repoId}
	json.NewEncoder(w).Encode(response)
}

func ResticRepoInit() (int, error) {
	var repoId int

	os.Setenv("RESTIC_PASSWORD", RepoData.Password)
	funcs := map[string]func() error{"local": ResticLocalInit, "sftp": ResticSftpInit, "bb": ResticBBInit, "s3": ResticS3Init}
	err := funcs[RepoData.Type]()
	utils.Check(err, "")

	modelfuncs := map[string]func() (int, error){"local": ResticLocalSave, "sftp": ResticSftpSave, "bb": ResticBBSave, "s3": ResticS3Save}
	repoId, err = modelfuncs[RepoData.Type]()
	utils.Check(err, "")

	return repoId, err
}

func ResticLocalInit() error {
	var cmd = "restic init --repo " + RepoData.Data["path"]
	_, err := exec.Command("bash", "-c", cmd).Output()
	utils.Check(err, "")

	return err
}

func ResticSftpInit() error {
	os.Setenv("RESTIC_PASSWORD", RepoData.Password)
	var cmd = "restic init --repo " + RepoData.Data["path"]
	_, err := exec.Command("bash", "-c", cmd).Output()
	utils.Check(err, "fatal")

	return nil
}

func ResticBBInit() error {
	os.Setenv("RESTIC_PASSWORD", RepoData.Password)
	var cmd = "restic init --repo " + RepoData.Data["path"]
	_, err := exec.Command("bash", "-c", cmd).Output()
	utils.Check(err, "fatal")

	return nil
}

func ResticS3Init() error {
	os.Setenv("RESTIC_PASSWORD", RepoData.Password)
	var cmd = "restic init --repo " + RepoData.Data["path"]
	_, err := exec.Command("bash", "-c", cmd).Output()
	utils.Check(err, "fatal")

	return nil
}

func ResticLocalSave() (int, error) {
	var repoData models.Repository

	repoData.Name = RepoData.Name
	repoData.Passwd = RepoData.Password
	repoData.Path = RepoData.Data["path"]
	repoData.Kind = RepoData.Type
	jsonString, err := json.Marshal(RepoData.Data)
	repoData.Data = string(jsonString)

	repoId, err := models.AddRepository(repoData)
	utils.Check(err, "fatal")

	return repoId, err
}

func ResticSftpSave() (int, error) {
	var repoData models.Repository

	repoData.Name = RepoData.Name
	repoData.Passwd = RepoData.Password
	repoData.Path = "sftp://" + RepoData.Data["user"] + "@" + RepoData.Data["server"] + ":" + RepoData.Data["path"]
	repoData.Kind = RepoData.Type
	jsonString, err := json.Marshal(RepoData.Data)
	repoData.Data = string(jsonString)

	repoId, err := models.AddRepository(repoData)
	utils.Check(err, "fatal")

	return repoId, err
}

func ResticBBSave() (int, error) {
	var repoData models.Repository

	repoData.Name = RepoData.Name
	repoData.Passwd = RepoData.Password
	repoData.Path = "b2:" + RepoData.Data["bucket_name"] + ":" + RepoData.Data["path"]
	repoData.Kind = RepoData.Type
	jsonString, err := json.Marshal(RepoData.Data)
	repoData.Data = string(jsonString)

	repoId, err := models.AddRepository(repoData)
	utils.Check(err, "fatal")

	return repoId, err
}

func ResticS3Save() (int, error) {
	var repoData models.Repository

	repoData.Name = RepoData.Name
	repoData.Passwd = RepoData.Password
	repoData.Path = "s3:s3.amazonaws.com/" + RepoData.Data["bucket_name"]
	repoData.Kind = RepoData.Type
	jsonString, err := json.Marshal(RepoData.Data)
	repoData.Data = string(jsonString)

	repoId, err := models.AddRepository(repoData)
	utils.Check(err, "fatal")

	return repoId, err
}
