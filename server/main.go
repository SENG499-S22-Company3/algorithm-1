package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const PORT = "8080"

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

// Gets default value passed if no value exist for given environment variable.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func StartHTTPServer() {
	port := getEnv("PORT", PORT)

	// Define all routes for REST API
	http.HandleFunc("/", Root)
	http.HandleFunc("/generate_schedule", GenerateSchedule)
	http.HandleFunc("/check_schedule", CheckSchedule)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
