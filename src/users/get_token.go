package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/zanefinner-projects/social-media-api/src/config"
)

//GetToken ...
func GetToken(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Recieved->")
	fmt.Println(string(body))
	//Match with db records

	var token string
	db.
		Model(&config.User{}).
		Where("username = ?", uc.Username).
		Select("token").
		Row().
		Scan(&token)

	if token != "" /*metch*/ {
		fmt.Fprintln(w, `{"token":`+`"`+token+`"`+`}`)
	} else {
		rstr := randomString(64)
		fmt.Println(rstr)
		fmt.Fprintln(w, `{"token":`+`"`+rstr+`"`+`}`)
		//add rstr to assoc record
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
