package main

import (
	"os"
	"os/user"
	"path"

	"github.com/jinzhu/gorm"
	_ "github.com/mutecomm/go-sqlcipher"
)

var DNET_CONFIG_DIR string

func init() {

	if _t, err := user.Current(); err == nil {
		DNET_CONFIG_DIR = path.Join(_t.HomeDir, ".config", "DNetGtk")
	} else {
		panic(err.Error())
	}

}

func GetMainStorageFileDirPath(user_name string) string {
	ret := path.Join(DNET_CONFIG_DIR, user_name)
	return ret
}

func GetMainStorageFilePath(
	user_name string,
) string {
	ret := path.Join(
		GetMainStorageFileDirPath(user_name),
		"main.db",
	)
	return ret
}

func GetApplicationStorageFilePath(
	user_name string,
	application_name string,
) string {
	ret := path.Join(
		GetMainStorageFileDirPath(user_name),
		"modules",
		application_name+".db",
	)
	return ret
}

func OpenMainStorage(
	username string,
) (*gorm.DB, error) {
	filename := GetMainStorageFilePath(username)

	{
		d := path.Dir(filename)
		err := os.MkdirAll(d, 0700)
		if err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	return db, nil
}

func OpenApplicationStorage(
	user_name string,
	application_name string,
) (*gorm.DB, error) {
	filename := GetApplicationStorageFilePath(user_name, application_name)

	{
		d := path.Dir(filename)
		err := os.MkdirAll(d, 0700)
		if err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	return db, nil
}
