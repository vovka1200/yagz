package get

import (
	"github.com/vovka1200/yagz/generic"
)

const Method = "history.get"

const Float = 0
const Character = 1
const Log = 2
const Integer = 3
const Text = 4

type Params struct {
	generic.Params
	History  int      `json:"history"`
	ItemIds  []string `json:"itemids,omitempty"`
	TimeFrom int      `json:"time_from,omitempty"`
	TimeTill int      `json:"time_till,omitempty"`
}
