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
	ID         uint   `json:"id" gorm:"primary_key"`
	Slug       string `json:"slug"`
	FileType   string `json:"filetype"`
	UploadedAt string `json:"uploadedat"`
	Visibility string `json:"visibility"`
}
