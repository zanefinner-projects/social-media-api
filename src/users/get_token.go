package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/zanefinner-projects/social-media-api/src/config"
	"golang.org/x/crypto/bcrypt"
)

//GetToken ...
func GetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer fmt.Println("Get Token Endpoint Hit")

	db := config.ConnectDatabase(config.GetDBCreds())

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var uc UserCredentials

	err = json.Unmarshal(body, &uc)
	if err != nil {
		fmt.Println(err)
	}

	var evidence config.User

	db.
		Where(&config.
			User{Username: uc.Username}).
		Find(&config.
			User{}).
		Scan(&evidence)

	err = bcrypt.
		CompareHashAndPassword([]byte(evidence.Password),
			[]byte(uc.Password))

	if err != nil {
		result := &config.ResponseOk{
			Err: "Credentials do not match any account",
			Ok:  "no",
		}
		resultJSON, err := json.Marshal(result)
		if err != nil {
			return
		}
		fmt.Fprintln(w, string(resultJSON))
	} else {

		if evidence.Token != "" {
			result := &config.ResponseUser{
				Action: "Grab Token",
				//ID:       string(evidence.ID), need to convert from uint
				Username: evidence.Username,
				Token:    evidence.Token,
				Ok:       "yes",
			}
			resultJSON, err := json.Marshal(result)
			if err != nil {
				return
			}
			fmt.Fprintln(w, string(resultJSON))

		} else {
			rstr := randomString(64)
			fmt.Println(rstr)
			evidence.Token = rstr
			db.Save(&evidence)
			result := &config.ResponseUser{
				Action: "Create Token",
				//ID:       string(evidence.ID), need to convert from uint
				Username: evidence.Username,
				Token:    rstr,
				Ok:       "yes",
			}
			resultJSON, err := json.Marshal(result)
			if err != nil {
				return
			}
			fmt.Fprintln(w, string(resultJSON))
		}
	}
}

func randomString(n int) string {
	var letters = []rune("QAZWSXEDCRFVTGBYHNUJMIKOLPqazwsxedcrfvtgbyhnujmikolp!1234567890")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
