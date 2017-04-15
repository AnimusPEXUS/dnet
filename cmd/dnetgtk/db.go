package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OwnData struct {
	ValueName string `gorm:"primary_key"`
	Value     string
}

/*
func (OwnData) TableName() string {
	return "own_data"
}
*/

type ApplicationPreset struct {
	Module    string `gorm:"primary_key"`
	Enabled   bool
	LastReKey time.Time
	Key       [200]byte
}

type NetworkPresetRemoved struct {
	Name    string `gorm:"primary_key"`
	Module  string
	Enabled bool
	Config  string
}

type AppDB struct {
	
}

type DB struct {
	filename string
	password string
	db       *gorm.DB
	app_db map[string]*gorm.DB
	
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
		if err := db.CreateTable(&OwnData{}).Error; err != nil {
			fmt.Println("error creating OwnData table")
		}
	}

	if !db.HasTable(&NetworkPreset{}) {
		if err := db.CreateTable(&NetworkPreset{}).Error; err != nil {
			fmt.Println("error creating NetworkPreset table")
		}
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

func (self *DB) SetNetPreset(
	name string,
	module string,
	enabled bool,
	config string,
) error {

	var pst []NetworkPreset

	if err := self.db.Find(
		&pst,
		&NetworkPreset{Name: name},
	).Error; err == nil && len(pst) != 0 {

		if len(pst) > 1 {
			for _, i := range pst {
				self.db.Delete(i)
			}
		}

		t := pst[0]

		t.Module = module
		t.Enabled = enabled
		t.Config = config
		self.db.Save(&t)

	} else {
		self.db.Create(
			&NetworkPreset{
				Name:    name,
				Module:  module,
				Enabled: enabled,
				Config:  config,
			},
		)
	}

	return nil
}

func (self *DB) GetNetPreset(name string) (
	found bool,
	module string,
	enabled bool,
	config string,
) {

	found = false
	module = ""
	enabled = false
	config = ""

	var pst NetworkPreset

	if err := self.db.First(
		&pst,
		&NetworkPreset{Name: name},
	).Error; err == nil {
		found = true
		module = pst.Module
		enabled = pst.Enabled
		config = pst.Config
	}

	return
}

func (self *DB) DelNetPreset(name string) {
	self.db.Delete(&NetworkPreset{Name: name})
}

func (self *DB) LstNetPresets() []string {
	ret := make([]string, 0)

	var pst []NetworkPreset

	if err := self.db.Find(&pst).Error; err == nil {
		for _, i := range pst {
			ret = append(ret, i.Name)
		}
	}

	return ret
}
