package main

import (
	//"database/sql"
	//"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OwnData struct {
	ValueName string
	Value     string
}

/*
func (OwnData) TableName() string {
	return "own_data"
}
*/

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
	db       *gorm.DB
}

func NewDB(filename, password string) (*DB, error) {
	ret := new(DB)
	ret.filename = filename
	ret.password = password

	//fmt.Println("db filename", filename)

	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	ret.db = db

	/*
		_, err = db.Exec("PRAGMA key = ?;", password)
		if err != nil {
			db.Close()
			return nil, err
		}
	*/

	err = db.Exec("VACUUM").Error
	if err != nil {
		db.Close()
		return nil, err
	}

	if !db.HasTable(&OwnData{}) {
		db.CreateTable(&OwnData{})
	}

	return ret, nil

}

func (self *DB) SetOwnPrivKey(txt string) {
	var own_key OwnData
	if err := self.db.First(
		&own_key,
		&OwnData{ValueName: "privkey"},
	).Error; err != nil {
		self.db.Create(&OwnData{ValueName: "privkey", Value: txt})
	} else {
		own_key.Value = txt
		self.db.Save(&own_key)
	}
}

func (self *DB) GetOwnPrivKey() (string, error) {
	var own_key OwnData
	if err := self.db.First(
		&own_key,
		&OwnData{ValueName: "privkey"},
	).Error; err != nil {
		return "", err
	}
	return own_key.Value, nil
}

func (self *DB) SetOwnTLSCertificate(txt string) {
	var t OwnData
	if err := self.db.First(
		&t,
		&OwnData{ValueName: "tls_certificate"},
	).Error; err != nil {
		self.db.Create(&OwnData{ValueName: "tls_certificate", Value: txt})
	} else {
		t.Value = txt
		self.db.Save(&t)
	}
}

func (self *DB) GetOwnTLSCertificate() (string, error) {
	var t OwnData
	if err := self.db.First(
		&t,
		&OwnData{ValueName: "tls_certificate"},
	).Error; err != nil {
		return "", err
	}
	return t.Value, nil
}
