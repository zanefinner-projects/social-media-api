package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/zanefinner-projects/social-media-api/src/config"
	"golang.org/x/crypto/bcrypt"
)

//Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Create User Endpoint Hit")

	db := config.ConnectDatabase(config.GetDBCreds())

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	//To show json, use string(body). To use data, call from ud
	var ud UserDataForUserCreate
	err = json.Unmarshal(body, &ud)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		go fmt.Println(err)
		fmt.Fprintln(w, `{"err":`+`"`+string(err.Error())+`"`+`}`)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.Password), 8)
	if err != nil {
		fmt.Println(err)
	}
	validAndUnique := isUni(db, ud) && isValid(db, ud)
	if validAndUnique {
		db.Create(&config.User{Username: ud.Username, Password: string(hashedPassword)})
		fmt.Fprintln(w, `{"created_user":`+`"`+"true"+`"`+`}`)
	} else {
		fmt.Fprintln(w, `{"err":`+`"`+"Credentials aren't sufficient"+`"`+`}`)
		fmt.Fprintln(w, `{"created_user":`+`"`+"false"+`"`+`}`)
	}
	fmt.Println("Recieved->")
	fmt.Println(string(body))
}

func isValid(db *gorm.DB, creds UserDataForUserCreate) bool {
	usernameOK := (len([]byte(creds.Username)) >= 4)
	passwordOK := (len([]byte(creds.Password)) >= 6)
	if usernameOK && passwordOK {
		return true
	}
	return false
}
func isUni(db *gorm.DB, creds UserDataForUserCreate) bool {
	var evidence config.User
	db.Where(&config.User{Username: creds.Username}).Find(&config.User{}).Scan(&evidence)
	if evidence.ID == 0 {
		return true
	}
	return false
}
