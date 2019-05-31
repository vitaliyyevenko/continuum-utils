package models

import "time"

//AssetNetwork is the struct definition of /resources/asset/assetNetwork
type AssetNetwork struct {
	Vendor              string    `json:"vendor,omitempty" cql:"vendor"`
	Product             string    `json:"product,omitempty" cql:"product"`
	DhcpEnabled         bool      `json:"dhcpEnabled" cql:"dhcp_enabled"`
	DhcpServer          string    `json:"dhcpServer,omitempty" cql:"dhcp_server" default:"0.0.0.0"`
	DhcpLeaseObtained   time.Time `json:"dhcpLeaseObtained,omitempty" cql:"dhcp_lease_obtained"`
	DhcpLeaseExpires    time.Time `json:"dhcpLeaseExpires,omitempty" cql:"dhcp_lease_expires"`
	DnsServers          []string  `json:"dnsServers,omitempty" cql:"dns_servers"`
	IPEnabled           bool      `json:"ipEnabled" cql:"ip_enabled"`
	IPv4                string    `json:"ipv4,omitempty" cql:"ipv4" default:"0.0.0.0"`
	IPv4List            []string  `json:"ipv4List,omitempty" cql:"ipv4_list"`
	IPv6                string    `json:"ipv6,omitempty" cql:"ipv6" default:"::"`
	IPv6List            []string  `json:"ipv6List,omitempty" cql:"ipv6_list"`
	SubnetMask          string    `json:"subnetMask,omitempty" cql:"subnet_mask" default:"0.0.0.0"`
	SubnetMasks         []string  `json:"subnetMasks,omitempty" cql:"subnet_masks"`
	DefaultIPGateway    string    `json:"defaultIPGateway,omitempty" cql:"default_ip_gateway" default:"0.0.0.0"`
	DefaultIPGateways   []string  `json:"defaultIPGateways,omitempty" cql:"default_ip_gateways"`
	MacAddress          string    `json:"macAddress,omitempty" cql:"mac_address"`
	WinsPrimaryServer   string    `json:"winsPrimaryServer,omitempty" cql:"wins_primary_server" default:"0.0.0.0"`
	WinsSecondaryServer string    `json:"winsSecondaryServer,omitempty" cql:"wins_secondary_server" default:"0.0.0.0"`
	LogicalName         string    `json:"logicalName,omitempty" cql:"logical_name"`
	InterfaceIndex      int       `json:"interfaceIndex" cql:"interface_index"`
	Speed               int64     `json:"speed" cql:"speed"`
	IPv4IntfCname       string    `json:"ipv4IntfCname,omitempty" cql:"ipv4_intf_cname"`
	IPv4NtSign          string    `json:"ipv4NtSign,omitempty" cql:"ipv4_nt_sign"`
	EthMedoption        []string  `json:"ethMedoption,omitempty" cql:"eth_medoption"`
	EthMedSubtype       string    `json:"ethMedSubtype,omitempty" cql:"eth_med_subtype"`
	DhcpDomainname      string    `json:"dhcpDomainname,omitempty" cql:"dhcp_domainname"`
	DhcpLeaseDuration   int       `json:"dhcpLeaseDuration,omitempty" cql:"dhcp_lease_duration"`
	DhcpMsgType         string    `json:"dhcpMsgType,omitempty" cql:"dhcp_msg_type"`
	DhcpRouters         string    `json:"dhcpRouters,omitempty" cql:"dhcp_routers"`
	DhcpSubnetmask      string    `json:"dhcpSubnetmask,omitempty" cql:"dhcp_subnetmask"`
}
