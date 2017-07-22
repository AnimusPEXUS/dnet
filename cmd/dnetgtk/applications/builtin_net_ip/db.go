package builtin_net_ip

type Data struct {
	ValueName string `gorm:"primary_key"`
	Value     string
}

type InstanceConfig struct {
	TCPListenerEnabled bool
	UDPBeaconEnabled   bool
	UDPListenerEnabled bool
	TCPListenerPort    int
	UDPPort            int
	UDPBeaconInterval  int
}

type DiscoveryHistory struct {
	dnet_addr_str   string
	ip              string
	port            int
	discovered_time string
}
