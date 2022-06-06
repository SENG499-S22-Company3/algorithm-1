package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Algorithm 1 REST server is alive!")
}

func GenerateSchedule(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Call parsing method and scheduling algorithm
	// For now, only outputs the body provided (JSON data)
	// and if it's empty just returns a basic string.
	if len(reqBody) == 0 {
		fmt.Fprintf(w, "This will generate a schedule!")
	} else {
		fmt.Fprint(w, string(reqBody))
	}
}

func CheckSchedule(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No body parameters provided")
	}

	// TODO: Call parsing method and scheduling checking
	// method. For now, only outputs the body provided
	// (JSON data), and if it's empty just returns a basic string.
	if len(reqBody) == 0 {
		fmt.Fprintf(w, "This will check a schedule is valid!")
	} else {
		fmt.Fprint(w, string(reqBody))
	}
}

func StartHTTPServer() {
	// Define all routes for REST API
	http.HandleFunc("/", Root)
	http.HandleFunc("/generate_schedule", GenerateSchedule)
	http.HandleFunc("/check_schedule", CheckSchedule)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
