package posts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/zanefinner-projects/social-media-api/src/config"
)

//Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer fmt.Println("Create Post Endpoint Hit")

	db := config.ConnectDatabase(config.GetDBCreds())

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var pd UserAuth
	err = json.Unmarshal(body, &pd)
	if err != nil {
		go fmt.Println(err)
		fmt.Fprintln(w, `{"err":`+`"`+string(err.Error())+`"`+`}`)
		return
	}

	matched := match(db, pd)
	if matched {
		db.Create(&config.Upload{Slug: "", FileType: "nofile", Visibility: "public", Content: pd.Content, Source: pd.Username})
		fmt.Fprintln(w, `{"created_post":`+`"`+"true"+`"`+`}`)
	} else {
		fmt.Fprintln(w, `{"err":`+`"`+"Account not authorized"+`"`+`}`)
	}
}

func match(db *gorm.DB, creds UserAuth) bool {
	var evidence config.User
	db.Where(&config.User{Username: creds.Username}).Find(&config.User{}).Scan(&evidence)
	if creds.Username == evidence.Username && creds.Token == evidence.Token {
		return true
	}
	return false
}
