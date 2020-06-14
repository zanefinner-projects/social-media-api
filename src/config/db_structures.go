package config

import "github.com/jinzhu/gorm"

//User is the structure used to generate the sql tables
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	//JoinedAt string `json:"joinedat"`
}

//Upload is the structure used to generate the sql tables
type Upload struct {
	gorm.Model
	Slug       string `json:"slug"`
	Username   string `json:"username"`
	FileType   string `json:"filetype"`
	Visibility string `json:"visibility"`
	Content    string `json:"content"`
}
