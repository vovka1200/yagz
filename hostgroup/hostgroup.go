package hostgroup

type HostGroup struct {
	Id       string `json:"groupid"`
	Name     string `json:"name"`
	Internal string `json:"internal"`
}

type HostGroups []HostGroup
