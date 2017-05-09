package builtin_keysandcerts

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
