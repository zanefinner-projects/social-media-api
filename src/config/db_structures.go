package config

//User is the structure used to generate the sql tables
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	JoinedAt string `json:"joinedat"`
}
