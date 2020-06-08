package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Create User Endpoint Hit")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	//To show json, use string(body). To use data, call from ud
	var ud UserDataForUserCreate
	err = json.Unmarshal(body, &ud)
	if err != nil {
		fmt.Println(err)
	}
	//See if creds are valid
	//yes? -> add record
	//no? -> send json of errors
	fmt.Println("Recieved->")
	fmt.Println(string(body))
}
