package models

import (
	"strings"
)

type Repository struct {
	Id        int      `json:"id"`
	Created   string   `json:"created"`
	Path      string   `json:"path"`
	Passwd    string   `json:"passwd"`
	Kind      string   `json:"kind"`
	Data      string   `json:"data"`
	BackupIds []string `json:"backup_ids"`
	Backups   Backups  `json:"backups"`
	Status    bool     `json:"status"`
}

type Backup struct {
	Id      int    `json:"id"`
	Created string `json:"created"`
	RepoId  int    `json:"repo_id"`
	Name    string `json:"name"`
	Source  string `json:"source"`
	Status  int    `json:"status"`
}

type Repos map[int]Repository

type Backups map[int]Backup

func GetRepositories() (Repos, error) {
	var items = map[int]Repository{}
	rows, err := Db.Query("SELECT r.repository_id, r.created, r.path, r.password, r.type, r.data, group_concat(b.backup_id) as backups " +
		"FROM repositories AS r JOIN backups AS b USING(repository_id) GROUP BY r.repository_id;")
	CheckErr(err)

	var repository_id int
	var created string
	var path string
	var password string
	var kind string
	var data string
	var backups string

	defer rows.Close()
	for rows.Next() {
		var item Repository
		err = rows.Scan(&repository_id, &created, &path, &password, &kind, &data, &backups)
		CheckErr(err)
		item = Repository{
			repository_id,
			created,
			path,
			password,
			kind,
			data,
			[]string (strings.Split(backups, ",")),
			nil,
			false}
		items[item.Id] = item
	}
	return items, nil
}

func GetBackups() (Backups, error) {
	var items = map[int]Backup{}
	rows, err := Db.Query("SELECT b.backup_id, b.created, b.repository_id, b.name, b.source, b.status FROM backups AS b")
	CheckErr(err)

	var backup_id int
	var created string
	var repository_id int
	var name string
	var source string
	var status int

	defer rows.Close()
	for rows.Next() {
		var item Backup
		err = rows.Scan(&backup_id, &created, &repository_id, &name, &source, &status)
		CheckErr(err)
		item = Backup{
			backup_id,
			created,
			repository_id,
			name,
			source,
			status}
		items[item.Id] = item
	}
	return items, nil
}
