package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server Start")
	defer fmt.Println("Server End")
	router := mux.NewRouter().
		StrictSlash(true)

	//Index Endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Route Guide
		fmt.Fprintln(w, `Routes:`)
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, `USERS`)
		fmt.Fprintln(w, ` | ANY:      [ /           ]   ➤   Route List`)
		fmt.Fprintln(w, ` | POST:     [ /users      ]   ➤   Create User`)
		fmt.Fprintln(w, ` | GET:      [ /users      ]   ➤   List all Users`)
		fmt.Fprintln(w, ` | GET:      [ /users/{id} ]   ➤   Singular User Info`)
		fmt.Fprintln(w, ` | DELETE:   [ /users/{id} ]   ➤   Delete a User`)
		fmt.Fprintln(w, ` | PUT:      [ /users/{id} ]   ➤   Modify a User`)
		fmt.Fprintln(w, "")

		fmt.Fprintln(w, `POSTS`)
		fmt.Fprintln(w, ` | POST:     [ /posts      ]   ➤   Create Post`)
		fmt.Fprintln(w, ` | GET:      [ /posts      ]   ➤   List all Posts`)
		fmt.Fprintln(w, ` | GET:      [ /posts/{id} ]   ➤   Singular Post Info`)
		fmt.Fprintln(w, ` | DELETE:   [ /posts/{id} ]   ➤   Delete a Post`)
		fmt.Fprintln(w, ` | PUT:      [ /posts/{id} ]   ➤   Modify a Post`)

		fmt.Fprintln(w, `MULTIMEDIA`)
		fmt.Fprintln(w, ` | POST:     [ /media      ]   ➤   Create Media`)
		fmt.Fprintln(w, ` | GET:      [ /media      ]   ➤   List all Media`)
		fmt.Fprintln(w, ` | GET:      [ /media/{id} ]   ➤   Serve Media File`)
		fmt.Fprintln(w, ` | DELETE:   [ /media/{id} ]   ➤   Delete Media`)
		fmt.Fprintln(w, ` | PUT:      [ /media/{id} ]   ➤   Modify Media`)
	})

	//Static Files
	router.
		PathPrefix("/media/"). //Will be activated with perms in the future
		Handler(http.
			StripPrefix("/media/", http.
				FileServer(http.Dir("./media/"))))

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
