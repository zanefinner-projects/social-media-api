package posts

import "github.com/jinzhu/gorm"

//UserAuth is the struct to organize authorization
type UserAuth struct {
	gorm.Model
	Username string
	Token    string
	Content  string
}
