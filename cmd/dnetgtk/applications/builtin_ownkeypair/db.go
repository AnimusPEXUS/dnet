package builtin_ownkeypair

import (
	"github.com/jinzhu/gorm"
)

type OwnData struct {
	ValueName string `gorm:"primary_key"`
	Value     string
}

type DB struct {
	db *gorm.DB
}

func (self *DB) SetOwnPrivKey(txt string) error {
	var own_key OwnData
	if err := self.db.First(
		&own_key,
		&OwnData{ValueName: "privkey"},
	).Error; err != nil {
		return self.db.Create(&OwnData{ValueName: "privkey", Value: txt}).Error
	} else {
		own_key.Value = txt
		return self.db.Save(&own_key).Error
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
