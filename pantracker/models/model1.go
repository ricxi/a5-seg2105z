// Author: Richard Xiong

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	patientDB := &PatientDB{
		filename: filename,
		patients: Patients{},
	}
	patientDB.load()

	return patientDB
}

// ResetPatientDB can be called to completely erase
// the patient database and create a new one
func (p *PatientDB) ResetPatientDB(filename string) {
	// create a json file
	_, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	p.sync()
}

// AddPatient adds a new patient to a slice/array that
// is currently loaded in memory, and then it syncs
// with the database to add the new patient
func (p *PatientDB) AddPatient(patient Patient) {

	p.patients.Patient = append(p.patients.Patient, patient)
	p.sync()

}

// GetPatients returns a array/slice of JSON structs that
// contain all the patients and their information
func (p *PatientDB) GetPatients() []Patient {
	return p.patients.Patient
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

// sync updates the make-shift JSON database every time it is called
func (p *PatientDB) sync() {
	patientsByteSlice, err := json.MarshalIndent(p.patients, "", "  ") // create a byte slice of Patients struct with indendation
	err = ioutil.WriteFile(p.filename, patientsByteSlice, 0644)        // write the byte slice to the json file as a json object
	if err != nil {
		fmt.Println(err)
	}
}

// load is called to load the database into
// memory, so that it can accessed faster
func (p *PatientDB) load() {
	dbInfo, err := ioutil.ReadFile(p.filename)
	if err != nil {
		log.Println(err)
	}

	if json.Unmarshal(dbInfo, &p.patients); err != nil {
		log.Println(err)
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
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone        int    `json:"phone"`
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
