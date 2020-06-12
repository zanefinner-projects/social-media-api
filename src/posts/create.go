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
		result := &config.ResponseOk{
			Err: err.Error(),
			Ok:  "no",
		}
		resultJSON, err := json.Marshal(result)
		if err != nil {
			return
		}
		fmt.Fprintln(w, string(resultJSON))
		return
	}

	matched := match(db, pd)
	if matched {
		db.Create(&config.Upload{Slug: "", FileType: "nofile", Visibility: "public", Content: pd.Content, Source: pd.Username})
		result := &config.ResponsePost{
			Action: "Post created",
			//ID       : find id
			Source: pd.Username,
			Ok:     "yes",
			//Time get the time
		}
		resultJSON, err := json.Marshal(result)
		if err != nil {
			return
		}
		fmt.Fprintln(w, string(resultJSON))
	} else {
		result := &config.ResponseOk{
			Ok:  "no",
			Err: "Invalid credentials",
		}
		resultJSON, err := json.Marshal(result)
		if err != nil {
			return
		}
		fmt.Fprintln(w, string(resultJSON))
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
