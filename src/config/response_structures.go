package config

//ResponseUser returns a response to be unarshalled into valid json
type ResponseUser struct {
	Action   string `json:"action"`
	ID       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Err      string `json:"err"`
	Ok       string `json:"ok"`
	Time     string `json:"time"`
}

//ResponsePost returns a response to be unarshalled into valid json
type ResponsePost struct {
	Action   string `json:"action"`
	ID       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Err      string `json:"err"`
	Ok       string `json:"ok"`
	Time     string `json:"time"`
}
