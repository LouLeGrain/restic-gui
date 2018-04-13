package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"simbookee/restic-gui/utils"
)

const DB_PATH = "./backups.db"

var DbPath string
var Db *sql.DB

func GetDb(t string) (bool, error) {
	switch t {
	case "sqlite":
		_, err := sqliteConnect()
		utils.CheckErr(err, "")
	}

	return true, nil
}

func sqliteConnect() (bool, error) {
	isNew := false
	if _, err := os.Stat(DB_PATH); os.IsNotExist(err) {
		isNew = true
	}
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil || db == nil {
		log.Printf("%+v\n", err)
		panic(err)
	}

	Db = db
	if isNew == true {
		sqliteMigrate()
	}

	return true, nil
}

func sqliteMigrate() {
	sql := `PRAGMA foreign_keys = false;
		
		CREATE TABLE  IF NOT EXISTS repositories (
			 repository_id integer,
			 created integer NOT NULL DEFAULT CURRENT_TIMESTAMP,
			 path text,
			 password text,
			 type text DEFAULT 'local',
			 data text DEFAULT '{}',
			PRIMARY KEY("repository_id")
		);
		
		CREATE TABLE  IF NOT EXISTS backups (
			 backup_id integer,
			 created integer NOT NULL DEFAULT CURRENT_TIMESTAMP,
			 repository_id integer NOT NULL , 
			 name text,
			 source text,
			 status integer,
			PRIMARY KEY("backup_id")
		);
		
		PRAGMA foreign_keys = true;
		
		INSERT INTO repositories (path, password) VALUES ('/backups','secretpasswd');
		INSERT INTO repositories (path, password) VALUES ('/backups/new','secretpasswd');
		INSERT INTO backups (repository_id, name, source, status) VALUES (1, 'Desktop', '~/Desktop', 1);
		INSERT INTO backups (repository_id, name, source, status) VALUES (2, 'Movies', '~/Movies', 1);
		INSERT INTO backups (repository_id, name, source, status) VALUES (1, 'Users', '/Users', 0);
		`

	_, err := Db.Exec(sql)
	utils.CheckErr(err, "")
	Db.Close()

	_, err = sqliteConnect()
	utils.CheckErr(err, "")
}

func CheckErr(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
		panic(err)
	}
}
