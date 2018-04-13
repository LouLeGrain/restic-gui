package models

import (
	"simbookee/restic-gui/utils"
)

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
	utils.CheckErr(err, "fatal")
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
		utils.CheckErr(err, "fatal")
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