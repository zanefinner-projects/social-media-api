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
	router := mux.NewRouter()
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
	})

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
