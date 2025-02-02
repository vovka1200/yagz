package get

import (
	"github.com/vovka1200/yagz/generic"
)

const Method = "item.get"

type Params struct {
	generic.Params
	GroupIds    []string `json:"groupids,omitempty"`
	SelectHosts []string `json:"selectHosts,omitempty"`
}
