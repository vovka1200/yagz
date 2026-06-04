package get

import (
	"git.grav.su/andromeda/yagz/generic"
)

const Method = "host.get"

type Params struct {
	generic.Params
	GroupIds         []string `json:"groupids,omitempty"`
	SelectInterfaces []string `json:"selectInterfaces,omitempty"`
}
