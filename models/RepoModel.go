package models

import (
	"restic-gui/utils"
)

type Repository struct {
	Id      int    `json:"id"`
	Created string `json:"created"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	Passwd  string `json:"passwd"`
	Kind    string `json:"kind"`
	Data    string `json:"data"`
}

type Repos map[int]Repository

func GetRepositories() (Repos, error) {
	var items = map[int]Repository{}
	rows, err := Db.Query("SELECT r.repository_id, r.created, r.name, r.path, r.password, r.type, r.data FROM repositories AS r ;")
	utils.Check(err, "fatal")

	var repository_id int
	var created string
	var name string
	var path string
	var password string
	var kind string
	var data string

	defer rows.Close()
	for rows.Next() {
		var item Repository
		err = rows.Scan(&repository_id, &created, &name, &path, &password, &kind, &data)
		utils.Check(err, "fatal")

		item = Repository{
			repository_id,
			created,
			name,
			path,
			password,
			kind,
			data}
		items[item.Id] = item
	}
	return items, nil
}

func AddRepository(repoData Repository) (int, error) {
	sql := "INSERT INTO repositories (name, path, password, type, data) VALUES (?,?,?,?,?);"
	stmt, err := Db.Prepare(sql)
	res, err := stmt.Exec(repoData.Name, repoData.Path, repoData.Passwd, repoData.Kind, repoData.Data)
	utils.Check(err, "")
	repoId, err := res.LastInsertId()
	utils.Check(err, "")

	return int(repoId), nil
}
