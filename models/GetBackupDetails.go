package models

import (
	"database/sql"
	"errors"
	"simbookee/restic-gui/utils"
	"strconv"
)

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
