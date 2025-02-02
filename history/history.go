package history

type Record struct {
	Clock       string `json:"clock,omitempty"`
	ItemId      string `json:"itemid,omitempty"`
	NanoSeconds string `json:"ns,omitempty"`
	Value       string `json:"value,omitempty"`
}
