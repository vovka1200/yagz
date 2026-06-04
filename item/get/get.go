package get

import (
	"git.grav.su/andromeda/yagz/generic"
)

const Method = "item.get"

type Params struct {
	generic.Params
	GroupIds    []string `json:"groupids,omitempty"`
	SelectHosts []string `json:"selectHosts,omitempty"`
}
