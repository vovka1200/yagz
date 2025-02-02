package get

import (
	"github.com/vovka1200/yagz/generic"
)

const Method = "trigger.get"

type Params struct {
	generic.Params
	GroupIds      []string `json:"groupids,omitempty"`
	SelectHosts   []string `json:"selectHosts,omitempty"`
	SkipDependent bool     `json:"skipDependent,omitempty"`
	Active        bool     `json:"active,omitempty"`
	MinSeverity   int      `json:"min_Severity"`
}
