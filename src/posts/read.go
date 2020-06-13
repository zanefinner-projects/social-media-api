package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zanefinner-projects/social-media-api/src/config"
)

//Read ...
func Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer fmt.Println("Read Posts endpoint hit!")

	db := config.ConnectDatabase(config.GetDBCreds())

	vars := mux.Vars(r)
	type recieved struct {
		ID string
	}
	current := recieved{
		ID: vars["id"],
	}
	var evidence config.Upload
	db.First(&config.Upload{}, current.ID).Scan(&evidence)
	result := &config.ResponsePost{
		Action:  "Post retrieved via ID",
		ID:      evidence.ID,
		Source:  evidence.Source,
		Ok:      "yes",
		Time:    evidence.CreatedAt,
		Content: evidence.Content,
	}
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return
	}
	fmt.Fprintln(w, string(resultJSON))

}
