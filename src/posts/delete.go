package posts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zanefinner-projects/social-media-api/src/config"
)

//Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer fmt.Println("Delete Post Endpoint Hit")

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
		vars := mux.Vars(r)
		type upload struct { //Using this instead of config.upload b/c of a type conversion issue
			ID string
		}
		example := upload{
			ID: vars["id"],
		}

		db.Unscoped().Delete(&example)
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
