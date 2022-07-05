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

	schedule.FallCourses = scheduling.Assignments(parsedInput.HardScheduled.FallCourses, parsedInput.CoursesToSchedule.FallCourses, parsedInput.Professors)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	schedule.SpringCourses = scheduling.Assignments(parsedInput.HardScheduled.SpringCourses, parsedInput.CoursesToSchedule.SpringCourses, parsedInput.Professors)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	schedule.SummerCourses = scheduling.Assignments(parsedInput.HardScheduled.SummerCourses, parsedInput.CoursesToSchedule.SummerCourses, parsedInput.Professors)
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

// Given a schedule, hard requirement checks are done to confirm none are violated
func CheckSchedule(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No body parameters provided")
	}

	parsedSchedule, err := structs.ParseHistorical(reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	requirementsViolated := false

	// Check for timeslot violations
	_, err = scheduling.BaseTimeslotMaps(parsedSchedule.FallCourses)
	if err != nil {
		fmt.Fprint(w, err.Error())
		requirementsViolated = true
	}
	_, err = scheduling.BaseTimeslotMaps(parsedSchedule.SpringCourses)
	if err != nil {
		fmt.Fprint(w, err.Error())
		requirementsViolated = true
	}
	_, err = scheduling.BaseTimeslotMaps(parsedSchedule.SummerCourses)
	if err != nil {
		fmt.Fprint(w, err.Error())
		requirementsViolated = true
	}

	// Check for professor violations
	// TO-DO

	if !requirementsViolated {
		fmt.Fprintf(w, "Schedule given is valid")
	} else {
		fmt.Fprint(w, "Schedule given has some violations that should be resolved")
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
