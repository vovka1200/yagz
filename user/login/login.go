package login

const Method = "user.login"

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
	UserData bool   `json:"userData,omitempty"`
}

type User struct {
	Userid        string `json:"userid"`
	Alias         string `json:"alias"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Url           string `json:"url"`
	Autologin     string `json:"autologin"`
	Autologout    string `json:"autologout"`
	Lang          string `json:"lang"`
	Refresh       string `json:"refresh"`
	Type          int    `json:"type"`
	Theme         string `json:"theme"`
	AttemptFailed string `json:"attempt_failed"`
	AttemptIp     string `json:"attempt_ip"`
	AttemptClock  string `json:"attempt_clock"`
	RowsPerPage   string `json:"rows_per_page"`
	DebugMode     int    `json:"debug_mode"`
	UserIp        string `json:"userip"`
	Node          struct {
		Name   string      `json:"name"`
		NodeId interface{} `json:"nodeid"`
	} `json:"node"`
	SessionId string `json:"sessionid"`
	GuiAccess int    `json:"gui_access"`
}
