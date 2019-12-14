package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"restic-gui/utils"
)

const DB_PATH = "backups.db"

var DbPath string
var Db *sql.DB

func GetDb(t string) (bool, error) {
	switch t {
	case "sqlite":
		_, err := sqliteConnect()
		utils.Check(err, "fatal")
	}

	return true, nil
}

func getDbPath() string {
	return utils.GetDataPath() + DB_PATH
}

func sqliteConnect() (bool, error) {
	isNew := false
	if _, err := os.Stat(getDbPath()); os.IsNotExist(err) {
		isNew = true
	}
	db, err := sql.Open("sqlite3", getDbPath())
	if err != nil || db == nil {
		log.Fatal("%+v\n", err)
		panic(err)
	}

	Db = db
	if isNew == true {
		err := sqliteMigrate()
		utils.Check(err, "fatal")

		//initialize demo repository
		/*os.Setenv("RESTIC_PASSWORD", "secretpasswd")
		var cmd = "restic init --repo /backups"
		_, err = exec.Command("bash", "-c", cmd).Output()
		utils.Check(err, "fatal")*/
	}

	return true, nil
}

func sqliteMigrate() error {
	sql := `PRAGMA foreign_keys = false;
		
		CREATE TABLE IF NOT EXISTS repositories (
			 repository_id integer,
			 created integer NOT NULL DEFAULT CURRENT_TIMESTAMP,
			 name text,
			 path text,
			 password text,
			 type text DEFAULT 'local',
			 data text DEFAULT '{}',
			PRIMARY KEY("repository_id")
		);
		
		CREATE TABLE IF NOT EXISTS backups (
			 backup_id integer,
			 created integer NOT NULL DEFAULT CURRENT_TIMESTAMP,
			 repository_id integer NOT NULL , 
			 name text,
			 source text,
			 status integer,
			 data text DEFAULT '{}',
			PRIMARY KEY("backup_id")
		);

		CREATE TABLE IF NOT EXISTS data (
   			id integer PRIMARY KEY AUTOINCREMENT,
			source text NOT NULL,
    		source_id int NOT NULL,
    		key text,
    		value text
		);
		CREATE UNIQUE INDEX data_source_source_id_key_uindex ON data (source, source_id, key);
		PRAGMA foreign_keys = true;
		
		--Initialze demo repo data
		--INSERT INTO repositories (name, path, password) VALUES ('Local Destination', '/backups','secretpasswd');
		--INSERT INTO backups (repository_id, name, source, status) VALUES (1, 'My Home Dir', '~/', 1);
		`

	_, err := Db.Exec(sql)
	utils.Check(err, "")
	Db.Close()

	_, err = sqliteConnect()
	utils.Check(err, "")

	return err
}

func Check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
		panic(err)
	}
}
