package main

import (
	"os/user"
	"path"

	"github.com/jinzhu/gorm"
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
	user_name string,
) (*gorm.DB, error) {
	filename := GetMainStorageFilePath(user_name)

	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func OpenApplicationStorage(
	user_name string,
	application_name string,
) (*gorm.DB, error) {
	filename := GetApplicationStorageFilePath(user_name)

	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
