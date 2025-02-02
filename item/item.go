package item

import (
	"github.com/vovka1200/yagz/host"
)

type Item struct {
	Id                   string      `json:"itemid,omitempty"`
	Type                 string      `json:"type,omitempty"`
	SnmpCommunity        string      `json:"snmp_community,omitempty"`
	SnmpOid              string      `json:"snmp_oid,omitempty"`
	Hostid               string      `json:"hostid,omitempty"`
	Name                 string      `json:"name,omitempty"`
	Key                  string      `json:"key_,omitempty"`
	Delay                string      `json:"delay,omitempty"`
	History              string      `json:"history,omitempty"`
	Trends               string      `json:"trends,omitempty"`
	LastValue            string      `json:"lastvalue,omitempty"`
	LastClock            string      `json:"lastclock,omitempty"`
	PrevValue            string      `json:"prevvalue,omitempty"`
	Status               string      `json:"status,omitempty"`
	ValueType            string      `json:"value_type,omitempty"`
	TrapperHosts         string      `json:"trapper_hosts,omitempty"`
	Units                string      `json:"units,omitempty"`
	Multiplier           string      `json:"multiplier,omitempty"`
	Delta                string      `json:"delta,omitempty"`
	PrevOrgValue         string      `json:"prevorgvalue,omitempty"`
	Snmpv3SecurityName   string      `json:"snmpv3_securityname,omitempty"`
	Snmpv3SecurityLevel  string      `json:"snmpv3_securitylevel,omitempty"`
	Snmpv3AuthPassPhrase string      `json:"snmpv3_authpassphrase,omitempty"`
	Snmpv3PrivPassPhrase string      `json:"snmpv3_privpassphrase,omitempty"`
	Formula              string      `json:"formula,omitempty"`
	Error                string      `json:"error,omitempty"`
	LastLogSize          string      `json:"lastlogsize,omitempty"`
	LogTimeFmt           string      `json:"logtimefmt,omitempty"`
	TemplateId           string      `json:"templateid,omitempty"`
	ValueMapId           string      `json:"valuemapid,omitempty"`
	DelayFlex            string      `json:"delay_flex,omitempty"`
	Params               string      `json:"params,omitempty"`
	IpmiSensor           string      `json:"ipmi_sensor,omitempty"`
	DataType             string      `json:"data_type,omitempty"`
	AuthType             string      `json:"authtype,omitempty"`
	Username             string      `json:"username,omitempty"`
	Password             string      `json:"password,omitempty"`
	PublicKey            string      `json:"publickey,omitempty"`
	PrivateKey           string      `json:"privatekey,omitempty"`
	Mtime                string      `json:"mtime,omitempty"`
	LastNS               string      `json:"lastns,omitempty"`
	Flags                string      `json:"flags,omitempty"`
	Filter               string      `json:"filter,omitempty"`
	InterfaceId          string      `json:"interfaceid,omitempty"`
	Port                 string      `json:"port,omitempty"`
	Description          string      `json:"description,omitempty"`
	InventoryLink        string      `json:"inventory_link,omitempty"`
	Lifetime             string      `json:"lifetime,omitempty"`
	Hosts                []host.Host `json:"hosts,omitempty"`
}

type Items []Item
