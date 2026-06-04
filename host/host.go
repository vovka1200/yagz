package host

import (
	"git.grav.su/andromeda/yagz/hostinterface"
)

type Host struct {
	Id         string                        `json:"hostid,omitempty"`
	Host       string                        `json:"host,omitempty"`
	Name       string                        `json:"name,omitempty"`
	Interfaces []hostinterface.HostInterface `json:"interfaces,omitempty"`
}
