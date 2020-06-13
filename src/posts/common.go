package posts

import (
	"github.com/jinzhu/gorm"
	"github.com/zanefinner-projects/social-media-api/src/config"
)

//UserAuth is the struct to organize authorization
type UserAuth struct {
	gorm.Model
	Username string
	Token    string
	Content  string
}

func match(db *gorm.DB, creds UserAuth) bool {
	var evidence config.User
	db.Where(&config.User{Username: creds.Username}).Find(&config.User{}).Scan(&evidence)
	if creds.Username == evidence.Username && creds.Token == evidence.Token {
		return true
	}
	return false
}
