package main

import (
	//"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OwnData struct {
	gorm.Model
	ValueName string
	Value     string
}

/*
type Networks struct {
	gorm.Model
	Name                      string
	NetworkModuleName         string
	NetworkModuleSettingsData string
	AutoRun                   bool
}
*/

type DB struct {
	filename string
	password string
	DB       *gorm.DB
}

func NewDB(filename, password string) (*DB, error) {
	ret := new(DB)
	ret.filename = filename
	ret.password = password

	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	ret.DB = db

	/*
		_, err = db.Exec("PRAGMA key = ?;", password)
		if err != nil {
			db.Close()
			return nil, err
		}
	*/

	return ret, nil

}

