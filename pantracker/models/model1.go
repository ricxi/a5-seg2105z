package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//PatientDB represents the patient database
type PatientDB struct {
	filename string
	patients Patients // store a array/slice of patients at all times, so that json file doesn't have to be opened each time
}

// NewPatientDB is called to generate a temporary database
// in order to store patient information
func NewPatientDB(filename string) *PatientDB {

	// create a json file
	_, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	patientDB := &PatientDB{
		filename: filename,
		patients: Patients{},
	}

	patientDB.updateDatabase()

	return patientDB
}

// AddPatient adds a new patient to the database
func (p *PatientDB) AddPatient(patient Patient) {

	p.patients.Patient = append(p.patients.Patient, patient)
	p.updateDatabase()

}

// GetPatients returns a list of patients
// FOR NOW, it just prints the list of patients to the console
func (p *PatientDB) GetPatients() {
	for _, pa := range p.patients.Patient {
		fmt.Println(fmt.Sprintf("%s %s", pa.FirstName, pa.LastName))
	}
}

// GetPatientByID returns information for the patient by their ID (health card number)
// FOR NOW, it just prints their info to the console
func (p *PatientDB) GetPatientByID(healthNumber int) {

	var patientInfo Patient
	for _, patient := range p.patients.Patient {
		if patient.HealthNumber == healthNumber {
			patientInfo = patient
		}
	}

	fmt.Println(patientInfo)
}

func (p *PatientDB) updateDatabase() {
	patientsByteSlice, err := json.MarshalIndent(p.patients, "", "  ") // create a byte slice of Patients struct with indendation
	err = ioutil.WriteFile(p.filename, patientsByteSlice, 0644)        // write the byte slice to the json file as a json object
	if err != nil {
		fmt.Println(err)
	}
}

// Patients holds a slice of type Patient
type Patients struct {
	Patient []Patient `json:"patients"`
}

// Patient struct is used to create/store patient info
// in the form of JSON data
type Patient struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	HealthNumber int    `json:"healthNumber"`
	Address      `json:"address"`
}

// Address struct used to create and store
// a JSON object that records address information
type Address struct {
	StreetNumber int    `json:"streetNumber"`
	StreetName   string `json:"streetName"`
	City         string `json:"city"`
	Province     string `json:"province"`
	PostalCode   string `json:"postalCode"`
}
