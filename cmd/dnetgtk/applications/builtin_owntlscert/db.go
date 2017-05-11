package builtin_owntlscert

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

func (self *DB) SetOwnTLSCertificate(txt string) error {
	var t OwnData
	if err := self.db.First(
		&t,
		&OwnData{ValueName: "tls_certificate"},
	).Error; err != nil {
		return self.db.Create(
			&OwnData{ValueName: "tls_certificate", Value: txt},
		).Error
	} else {
		t.Value = txt
		return self.db.Save(&t).Error
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
