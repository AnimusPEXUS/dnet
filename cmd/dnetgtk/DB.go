package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mutecomm/go-sqlcipher"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Data struct {
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

func (Data) TableName() string {
	return "settings"
}

func (ApplicationStatus) TableName() string {
	return "application_status"
}

type AppDB struct {
	name string
	db   *gorm.DB
}

func (self *AppDB) Key(value string) error {
	// TODO https://github.com/jinzhu/gorm/issues/1498
	return self.db.Exec(
		fmt.Sprintf("PRAGMA key = %v;", self.db.NewScope(nil).Quote(value)),
	).Error
}

func (self *AppDB) ReKey(value string) error {
	// TODO https://github.com/jinzhu/gorm/issues/1498
	return self.db.Exec(
		fmt.Sprintf("PRAGMA rekey = %v;", self.db.NewScope(nil).Quote(value)),
	).Error
}

type DB struct {
	username string
	key      string
	db       *gorm.DB
	app_db   []*AppDB

	get_app_db_mutex *sync.Mutex
}

func NewDB(
	username string,
	key string,
) (*DB, error) {
	ret := new(DB)
	ret.username = username
	ret.key = key

	ret.get_app_db_mutex = &sync.Mutex{}

	db, err := OpenMainStorage(username)
	if err != nil {
		return nil, err
	}

	ret.db = db

	if err := ret.Key(key); err != nil {
		return nil, err
	}

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

	if !db.HasTable(&Data{}) {
		if err := db.CreateTable(&Data{}).Error; err != nil {
			return nil, err
		}
	}

	if !db.HasTable(&ApplicationStatus{}) {
		if err := db.CreateTable(&ApplicationStatus{}).Error; err != nil {
			return nil, err
		}
	}

	return ret, nil

}

func (self *DB) GetAppDB(name string) (*AppDB, error) {

	self.get_app_db_mutex.Lock()
	defer self.get_app_db_mutex.Unlock()

	for _, i := range self.app_db {
		if i.name == name {
			return i, nil
		}
	}

	db, err := OpenApplicationStorage(self.username, name)
	if err != nil {
		return nil, err
	}

	ret := &AppDB{
		name: name,
		db:   db,
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

func (self *DB) Key(value string) error {
	// TODO https://github.com/jinzhu/gorm/issues/1498
	return self.db.Exec(
		fmt.Sprintf("PRAGMA key = %v;", self.db.NewScope(nil).Quote(value)),
	).Error
}

func (self *DB) ReKey(value string) error {
	// TODO https://github.com/jinzhu/gorm/issues/1498
	return self.db.Exec(
		fmt.Sprintf("PRAGMA rekey = %v;", self.db.NewScope(nil).Quote(value)),
	).Error
}
