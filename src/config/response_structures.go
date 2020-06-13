package config

import "time"

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
	Action  string `json:"action"`
	ID      uint
	Source  string    `json:"source"`
	Err     string    `json:"err"`
	Ok      string    `json:"ok"`
	Time    time.Time `json:"time"`
	Content string    `json:"content"`
}

//ResponseOk returns a response that says yes/no based on errors
type ResponseOk struct {
	Ok  string `json:"ok"`
	Err string `json:"err"`
}
