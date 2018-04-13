package models

import (
	"simbookee/restic-gui/utils"
)

type Repository struct {
	Id      int    `json:"id"`
	Created string `json:"created"`
	Path    string `json:"path"`
	Passwd  string `json:"passwd"`
	Kind    string `json:"kind"`
	Data    string `json:"data"`
}

type Repos map[int]Repository

func GetRepositories() (Repos, error) {
	var items = map[int]Repository{}
	rows, err := Db.Query("SELECT r.repository_id, r.created, r.path, r.password, r.type, r.data FROM repositories AS r ;")
	utils.CheckErr(err, "fatal")

	var repository_id int
	var created string
	var path string
	var password string
	var kind string
	var data string

	defer rows.Close()
	for rows.Next() {
		var item Repository
		err = rows.Scan(&repository_id, &created, &path, &password, &kind, &data)
		utils.CheckErr(err, "fatal")
		item = Repository{
			repository_id,
			created,
			path,
			password,
			kind,
			data}
		items[item.Id] = item
	}
	return items, nil
}
