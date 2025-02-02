package generic

const Extend = "extend"

type Auth struct {
	SessionId string `json:"auth,omitempty"`
}

type Filter map[string]interface{}

type Params struct {
	Auth
	Output any    `json:"output,omitempty"`
	Filter Filter `json:"filter,omitempty"`
}
