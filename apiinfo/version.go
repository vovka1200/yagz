package apiinfo

const Method = "apiinfo.version"

type Version struct {
	Result string `json:"result"`
}

type Params []string
