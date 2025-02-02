package trigger

import (
	"github.com/vovka1200/yagz/host"
)

type Trigger struct {
	Id          string      `json:"triggerid"`
	Description string      `json:"description"`
	Priority    string      `json:"priority"`
	LastChange  string      `json:"lastchange"`
	Hosts       []host.Host `json:"hosts,omitempty"`
}
