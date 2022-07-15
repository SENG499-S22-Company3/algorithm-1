package server

import (
	"algorithm-1/genetic"
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

	// Test there are no problems with the courses given by trying to generate an initial solution
	_, err = scheduling.Assignments(parsedInput.HardScheduled.FallCourses, parsedInput.CoursesToSchedule.FallCourses, parsedInput.Professors, "Fall")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	_, err = scheduling.Assignments(parsedInput.HardScheduled.SpringCourses, parsedInput.CoursesToSchedule.SpringCourses, parsedInput.Professors, "Spring")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	_, err = scheduling.Assignments(parsedInput.HardScheduled.SummerCourses, parsedInput.CoursesToSchedule.SummerCourses, parsedInput.Professors, "Summer")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Run genetic algorithm on each semester, will update to run concurrently
	if len(parsedInput.CoursesToSchedule.FallCourses) != 0 {
		schedule.FallCourses, err = genetic.RunGeneticAlg(parsedInput.HardScheduled.FallCourses, parsedInput.CoursesToSchedule.FallCourses, parsedInput.Professors, "Fall")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
	if len(parsedInput.CoursesToSchedule.SpringCourses) != 0 {
		schedule.SpringCourses, err = genetic.RunGeneticAlg(parsedInput.HardScheduled.SpringCourses, parsedInput.CoursesToSchedule.SpringCourses, parsedInput.Professors, "Spring")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
	if len(parsedInput.CoursesToSchedule.SummerCourses) != 0 {
		schedule.SummerCourses, err = genetic.RunGeneticAlg(parsedInput.HardScheduled.SummerCourses, parsedInput.CoursesToSchedule.SummerCourses, parsedInput.Professors, "Summer")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}

	marshalledJSON, err := structs.StructToJSON(schedule)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, string(marshalledJSON))
}

// Given a schedule, hard requirement checks are done to confirm none are violated
// Send schedule object in the "hardScheduled" place. The coursesToSchedule will be ignored so it can just be empty, "professors" will be used to make sure constraints are satisfied
func CheckSchedule(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No body parameters provided")
	}

	parsedInput, err := structs.ParseInput(reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	requirementsViolated := false

	// Check for professor and basic violations
	err = scheduling.ScheduleConstraintsCheck("Fall", parsedInput.HardScheduled.FallCourses, parsedInput.Professors)
	if err != nil {
		fmt.Fprint(w, err.Error())
		requirementsViolated = true
	}
	err = scheduling.ScheduleConstraintsCheck("Spring", parsedInput.HardScheduled.SpringCourses, parsedInput.Professors)
	if err != nil {
		fmt.Fprint(w, err.Error())
		requirementsViolated = true
	}
	err = scheduling.ScheduleConstraintsCheck("Summer", parsedInput.HardScheduled.SummerCourses, parsedInput.Professors)
	if err != nil {
		fmt.Fprint(w, err.Error())
		requirementsViolated = true
	}

	// Check for timeslot violations
	_, err = scheduling.BaseTimeslotMaps(parsedInput.HardScheduled.FallCourses, "Fall")
	if err != nil {
		fmt.Fprint(w, err.Error())
		requirementsViolated = true
	}
	_, err = scheduling.BaseTimeslotMaps(parsedInput.HardScheduled.SpringCourses, "Spring")
	if err != nil {
		fmt.Fprint(w, err.Error())
		requirementsViolated = true
	}
	_, err = scheduling.BaseTimeslotMaps(parsedInput.HardScheduled.SummerCourses, "Summer")
	if err != nil {
		fmt.Fprint(w, err.Error())
		requirementsViolated = true
	}

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
