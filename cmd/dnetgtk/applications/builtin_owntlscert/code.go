package builtin_keysandcerts

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

