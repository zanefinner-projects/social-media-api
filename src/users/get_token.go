package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

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
		fmt.Fprintln(w, `{"err":`+`"`+"Invalid login"+`"`+`}`)
	} else {

		if evidence.Token != "" {
			fmt.Fprintln(w, `{"token":`+`"`+evidence.Token+`",`)
			fmt.Fprintln(w, `"ID":`+`"`+strconv.FormatUint(uint64(evidence.ID), 10)+`"`+`}`)

		} else {
			rstr := randomString(64)
			fmt.Println(rstr)
			evidence.Token = rstr
			db.Save(&evidence)
			fmt.Fprintln(w, `{"token":`+`"`+rstr+`"`+`}`)
			fmt.Fprintln(w, `{"ID":`+`"`+string(evidence.ID)+`"`+`}`)
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
