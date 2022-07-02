package server

import (
	"algorithm-1/scheduling"
	"algorithm-1/structs"
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

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func Generate(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(reqBody) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request body cannot be empty."))
		return
	}

	parsedInput, err := structs.ParseInput(reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var schedule structs.Schedule

	schedule.FallCourses = scheduling.Assignments(parsedInput.HardScheduled.FallCourses, parsedInput.CoursesToSchedule.FallCourses, parsedInput.Professors, "Fall")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	
	schedule.SpringCourses = scheduling.Assignments(parsedInput.HardScheduled.SpringCourses, parsedInput.CoursesToSchedule.SpringCourses, parsedInput.Professors, "Spring")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	schedule.SummerCourses = scheduling.Assignments(parsedInput.HardScheduled.SummerCourses, parsedInput.CoursesToSchedule.SummerCourses, parsedInput.Professors, "Summer")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	marshalledJSON, err := structs.StructToJSON(schedule)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprint(w, string(marshalledJSON))
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
	http.HandleFunc("/healthcheck", HealthCheck)
	http.HandleFunc("/schedule", Generate)
	http.HandleFunc("/check_schedule", CheckSchedule)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
