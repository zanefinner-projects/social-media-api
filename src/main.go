package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/zanefinner-projects/social-media-api/src/config"
	"github.com/zanefinner-projects/social-media-api/src/posts"
	"github.com/zanefinner-projects/social-media-api/src/users"
)

func main() {
	fmt.Println("Server Start")
	config.Migrate(config.GetDBCreds())
	router := mux.NewRouter().
		StrictSlash(true)

	//Index Endpoint
	router.HandleFunc("/", documentation)
	router.
		HandleFunc("/echo/{msg}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			vars := mux.Vars(r)

			type msg struct {
				Message string
			}

			recieved := msg{
				Message: vars["msg"],
			}

			jsonResponse, err := json.Marshal(recieved)
			if err != nil {
				http.Error(w, err.Error(),
					http.
						StatusInternalServerError)
				return
			}
			fmt.Println(string(jsonResponse))
			w.Write(jsonResponse)
		})

	//Static Files
	router.
		PathPrefix("/media/"). //Will be activated with perms in the future
		Handler(http.
			StripPrefix("/media/", http.
				FileServer(http.Dir("./media/"))))

	//User Routes
	router.HandleFunc("/users", users.Create).
		Methods("POST")
	router.HandleFunc("/users", users.GetToken).
		Methods("GET")

	//Post Routers
	router.HandleFunc("/posts", posts.Create).
		Methods("POST")
	router.HandleFunc("/posts/{id}", posts.Read).
		Methods("GET")
	//Will have filtering routes later on. Ex: GET: /posts/bydate/recent "{'limit':'5'}"
	//Server Setup
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//Run Server
	log.Fatal(srv.ListenAndServe())
}

func documentation(w http.ResponseWriter, r *http.Request) {
	//Redirect to GITHUB page!
	http.Redirect(w, r, "https://github.com/zanefinner-projects/social-media-api", 301)
}
