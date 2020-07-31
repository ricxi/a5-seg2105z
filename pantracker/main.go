// Author: Richard Xiong

package main

import (
	"net/http"
	"pantracker/utilities"
	views "pantracker/views/pages"

	"github.com/gorilla/mux"
)

// import "pantracker/models"

// "encoding/json"

func main() {
	// patientDatabase := models.NewPatientDB("db.json")
	// patient1 := models.Patient{FirstName: "Dave", LastName: "Chappelle", HealthNumber: 1202853986}
	// patient2 := models.Patient{FirstName: "Martin", LastName: "King", HealthNumber: 802853986}
	// patientDatabase.AddPatient(patient1)
	// patientDatabase.AddPatient(patient2)

	// patientDatabase.GetPatientByID(1202853986)
	// patientDatabase.GetPatients()

	r := mux.NewRouter()

	v := views.NewView("views/pages/patient.html")

	r.Handle("/", v)

	address := utilities.GetConfig("config.json")
	http.ListenAndServe(address, r)
}
