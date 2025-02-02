package get

import (
	"github.com/vovka1200/yagz/generic"
)

const Method = "host.get"

type Params struct {
	generic.Params
	GroupIds         []string `json:"groupids,omitempty"`
	SelectInterfaces []string `json:"selectInterfaces,omitempty"`
}
