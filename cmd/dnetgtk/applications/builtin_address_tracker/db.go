package builtin_address_tracker

import (
	"time"

	"github.com/jinzhu/gorm"
)

type DB struct {
	db *gorm.DB
}

type AddressHistory struct {
	Key            uint64 `gorm:"primary_key"`
	DnetAddrStr    string
	DiscoveredDate *time.Time
	NetworkName    string
	NetworkAddress string
}
