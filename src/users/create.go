package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zanefinner-projects/social-media-api/src/config"
)

//Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Create User Endpoint Hit")

	db := config.ConnectDatabase(config.GetDBCreds())

	fmt.Println(db) //can use db as gorm.DB as of now
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
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
	//See if creds are valid
	//yes? -> add record
	//no? -> send json of errors
	fmt.Fprintln(w, `{"err":`+`"`+"Credentials aren't sufficient"+`"`+`}`)
	fmt.Println("Recieved->")
	fmt.Println(string(body))
}
