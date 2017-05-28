package builtin_net_ip

import (
	"net"

	"github.com/jinzhu/gorm"
)

type DB struct {
	db *gorm.DB
}

type DiscoveryHistory struct {
	dnet_addr_str   string
	ip              string
	port            int
	discovered_time string
}
