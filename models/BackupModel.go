package models

import (
	"database/sql"
	"errors"
	"simbookee/restic-gui/utils"
	"strconv"
)

type BackupData struct {
	RepoId int    `json:"repo_id"`
	Name   string `json:"name"`
	Source string `json:"source"`
	Status int    `json:"status"`
	Data   string `json:"data"`
}

type Backup struct {
	Id       int    `json:"id"`
	Created  string `json:"created"`
	RepoId   int    `json:"repo_id"`
	RepoPath string `json:"repo_path"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Source   string `json:"source"`
	Status   int    `json:"status"`
}

type Backups map[int]Backup

func GetBackups() (Backups, error) {
	var items = map[int]Backup{}
	query := "SELECT b.backup_id, b.created, b.repository_id, r.path, r.type, b.name, b.source, b.status " +
		"FROM backups AS b JOIN repositories AS r USING(repository_id) ORDER BY b.repository_id, b.name"

	rows, err := Db.Query(query)
	utils.Check(err, "fatal")
	var backup_id int
	var created string
	var repository_id int
	var path string
	var repoType string
	var name string
	var source string
	var status int

	defer rows.Close()
	for rows.Next() {
		var item Backup
		err = rows.Scan(&backup_id, &created, &repository_id, &path, &repoType, &name, &source, &status)
		utils.Check(err, "fatal")
		item = Backup{
			backup_id,
			created,
			repository_id,
			path,
			utils.UcFirst(repoType),
			name,
			source,
			status}
		items[item.Id] = item
	}
	return items, nil
}

func GetBackupDetails(id int) (map[string]string, error) {

	var m map[string]string
	m = make(map[string]string)

	var source string
	var path string
	var kind string
	var passwd string
	var data string

	query := "SELECT b.source, r.path, r.type, r.password, r.data FROM backups AS b JOIN repositories AS r USING(repository_id) WHERE b.backup_id = ?"
	err := Db.QueryRow(query, id).Scan(&source, &path, &kind, &passwd, &data)
	switch {
	case err == sql.ErrNoRows:
		err = errors.New("Bad request for id: " + strconv.Itoa(id))
	case err != nil:
		utils.Check(err, "fatal")
	default:
		m["source"] = source
		m["destination"] = path
		m["kind"] = kind
		m["passwd"] = passwd
		m["data"] = data
	}

	return m, err
}

func AddBackup(buData BackupData) (int, error) {
	sql := "INSERT INTO backups (name, source, repository_id, status) VALUES (?,?,?,?);"
	stmt, err := Db.Prepare(sql)
	res, err := stmt.Exec(buData.Name, buData.Source, buData.RepoId, buData.Status)
	utils.Check(err, "")
	repoId, err := res.LastInsertId()
	utils.Check(err, "")

	return int(repoId), nil
}
