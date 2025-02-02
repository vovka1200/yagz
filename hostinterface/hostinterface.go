package hostinterface

type HostInterface struct {
	InterfaceId  string                 `json:"interfaceid,omitempty"`
	Hostid       string                 `json:"hostid,omitempty"`
	Main         string                 `json:"main,omitempty"`
	Type         string                 `json:"type,omitempty"`
	UseIp        string                 `json:"useip,omitempty"`
	Ip           string                 `json:"ip,omitempty"`
	Dns          string                 `json:"dns,omitempty"`
	Port         string                 `json:"port,omitempty"`
	Available    string                 `json:"available,omitempty"`
	Error        string                 `json:"error,omitempty"`
	ErrorsFrom   string                 `json:"errors_from,omitempty"`
	DisableUntil string                 `json:"disable_until,omitempty"`
	Details      map[string]interface{} `json:"details,omitempty"`
}
