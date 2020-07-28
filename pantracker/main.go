package main

import "pantracker/models"

// "encoding/json"

func main() {
	patientDatabase := models.NewPatientDB("db.json")
	patient1 := models.Patient{FirstName: "Frank", LastName: "Sinatra", HealthNumber: 1202853986}
	patient2 := models.Patient{FirstName: "Stevie", LastName: "Wonder", HealthNumber: 802853986}
	patientDatabase.CreatePatient(patient1)
	patientDatabase.CreatePatient(patient2)

	patientDatabase.GetPatientByID(1202853986)

}
