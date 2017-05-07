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

type ApplicationStatus struct {
	Name        string `gorm:"primary_key"`
	Builtin     bool
	Checksum    string
	Enabled     bool
	LastDBReKey *time.Time
	DBKey       string
}

func (OwnData) TableName() string {
	return "settings"
}

func (ApplicationStatus) TableName() string {
	return "application_status"
}

type AppDB struct {
	Name string
	DB   *gorm.DB
}

type DB struct {
	username string
	key      string
	db       *gorm.DB
	app_db   []*AppDB
}

func NewDB(
	username string,
	key string,
) (*DB, error) {
	ret := new(DB)
	ret.username = username
	ret.key = key

	db, err := OpenMainStorage(username)
	if err != nil {
		return nil, err
	}

	ret.db = db

	/*
		err = db.Exec("VACUUM;").Error
		if err != nil {
			db.Close()
			fmt.Println("vacuum error:", err.Error())
			return nil, err
		}
	*/

	/*
		if err := db.Commit().Error; err != nil {
			fmt.Println("Commit error:", err.Error())
		}
	*/

	if !db.HasTable(&OwnData{}) {
		if err := db.CreateTable(&OwnData{}).Error; err != nil {
			fmt.Println("error creating OwnData table")
		}
	}

	if !db.HasTable(&ApplicationStatus{}) {
		if err := db.CreateTable(&ApplicationStatus{}).Error; err != nil {
			fmt.Println("error creating ApplicationStatus table")
		}
	}

	/*
		if err := db.Commit().Error; err != nil {
			fmt.Println("Commit error:", err.Error())
		}
	*/

	return ret, nil

}

func (self *DB) GetAppDB(name string) (*AppDB, error) {

	for _, i := range self.app_db {
		if i.Name == name {
			return i, nil
		}
	}

	db, err := OpenApplicationStorage(self.username, name)
	if err != nil {
		return nil, err
	}

	ret := &AppDB{
		Name: name,
		DB:   db,
	}

	self.app_db = append(
		self.app_db,
		ret,
	)

	return ret, nil
}

func (self *DB) ListApplicationStatusNames() []string {
	ret := make([]string, 0)

	var aps []ApplicationStatus

	if err := self.db.Find(&aps, &ApplicationStatus{}).Error; err == nil {
		for _, i := range aps {
			ret = append(ret, i.Name)
		}
	}

	return ret
}

/*
	Use this not only for getting info on name, but also for creating new
	Info for name
*/
func (self *DB) GetApplicationStatus(name string) (*ApplicationStatus, error) {

	var ap ApplicationStatus

	if err := self.db.First(
		&ap,
		&ApplicationStatus{Name: name},
	).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ap.Name = name
			ap.Checksum = ""
			ap.Enabled = false
			ap.LastDBReKey = new(time.Time)
			ap.DBKey = ""
			if self.db.NewRecord(ap) {
				self.db.Create(ap)
			}
			return &ap, nil
		} else {
			return nil, err
		}
	} else {
		return &ap, nil
	}
}

func (self *DB) DelApplicationStatus(name string) {
	var as []ApplicationStatus

	if self.db.Find(&as, &ApplicationStatus{Name: name}).Error == nil {

		for _, i := range as {
			self.db.Delete(i)
		}
	}
}

func (self *DB) SetApplicationStatus(value *ApplicationStatus) error {
	var err error

	if self.db.NewRecord(value) {

		err = self.db.Create(value).Error
	} else {

		err = self.db.Save(value).Error
	}
	return err
}
