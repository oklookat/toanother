package database

import (
	"database/sql"
	"os"

	"github.com/oklookat/toanother/core/datadir"

	_ "github.com/mattn/go-sqlite3"
)

// connect to DB.
//
// If DB not exists, creates it and executes afterCreated().
//
// path: absolute/relative path to DB. Like: "./spotify.db"
func Load(path string, afterCreated func(db *sql.DB) error) (db *sql.DB, err error) {
	// check.
	var isExists bool
	if isExists, err = datadir.IsFileExists(path); err != nil {
		return
	}

	if isExists {
		// check size (maybe .db empty)
		file, errd := datadir.OpenFile(path, os.O_RDONLY)
		if errd != nil {
			return
		}
		defer file.Close()
		stat, errd := file.Stat()
		if errd != nil {
			return
		}
		isExists = stat.Size() > 100
	}

	if !isExists {
		// create.
		if err = datadir.WriteFile(path, nil); err != nil {
			println(err.Error())
			return
		}
	}

	// open.
	var absPath string
	if absPath, err = datadir.GetFullPath(path); err != nil {
		return
	}
	if db, err = sql.Open("sqlite3", absPath); err != nil {
		return
	}
	if _, err = db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return
	}

	// hook.
	if !isExists && afterCreated != nil {
		err = afterCreated(db)
	}

	return
}
