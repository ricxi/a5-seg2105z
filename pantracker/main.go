package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// import "pantracker/models"

// "encoding/json"

// Config struct is used to load information from a config file
type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT"`
		Host string `envconfig:"SERVER_HOST"`
	}
}

func main() {
	// patientDatabase := models.NewPatientDB("db.json")
	// patient1 := models.Patient{FirstName: "Frank", LastName: "Sinatra", HealthNumber: 1202853986}
	// patient2 := models.Patient{FirstName: "Stevie", LastName: "Wonder", HealthNumber: 802853986}
	// patientDatabase.CreatePatient(patient1)
	// patientDatabase.CreatePatient(patient2)

	// patientDatabase.GetPatientByID(1202853986)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello homies")
	})

	http.ListenAndServe(":4000", r)
}
