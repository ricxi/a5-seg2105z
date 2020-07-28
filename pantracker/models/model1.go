package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Patients is a slice of patients
type Patients struct {
	Patient []Patient `json:"patients"`
}

// Patient records patient into
type Patient struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	HealthNumber int    `json:"healthNumber"`
	Address      `json:"address"`
}

// Address information
type Address struct {
	StreetNumber int    `json:"streetNumber"`
	StreetName   string `json:"streetName"`
	City         string `json:"city"`
	Province     string `json:"province"`
	PostalCode   string `json:"postalCode"`
}

//PatientDB represents the patient database
type PatientDB struct {
	filename string
}

// NewPatientDB is called to generate the database
func NewPatientDB(filename string) *PatientDB {

	patientByteSlice, err := json.MarshalIndent(Patients{}, "", "  ")

	_, err = os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(filename, patientByteSlice, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return &PatientDB{
		filename: filename,
	}
}

// CreatePatient adds a new patient to the DB
func (p *PatientDB) CreatePatient(patient Patient) {

	// get the json file
	file, err := ioutil.ReadFile(p.filename)
	if err != nil {
		panic(err)
	}

	// create a Patients struct
	patients := Patients{}
	err = json.Unmarshal(file, &patients) // store json file in patients struct
	if err != nil {
		panic(err)
	}

	patients.Patient = append(patients.Patient, patient)
	// fmt.Println(patients)
	byteArray, err := json.MarshalIndent(patients, "", "  ")
	err = ioutil.WriteFile(p.filename, byteArray, 0644)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}

}

// GetPatients returns a list of patients
// FOR NOW, it just prints the list of patients to the console
func (p *PatientDB) GetPatients() {
	// get the json file
	file, err := ioutil.ReadFile(p.filename)
	if err != nil {
		panic(err)
	}

	// create a Patients struct
	patients := Patients{}
	err = json.Unmarshal(file, &patients) // store json file in patients struct
	if err != nil {
		panic(err)
	}

	for _, pa := range patients.Patient {
		fmt.Println(pa.HealthNumber)
	}
}

// GetPatientByID returns information for the patient by their ID (health card number)
// FOR NOW, it just prints their info to the console
func (p *PatientDB) GetPatientByID(healthNumber int) {
	// get the json file
	file, err := ioutil.ReadFile(p.filename)
	if err != nil {
		panic(err)
	}

	// create a Patients struct
	patients := Patients{}
	err = json.Unmarshal(file, &patients) // store json file in patients struct
	if err != nil {
		panic(err)
	}

	var patientInfo Patient
	for _, pa := range patients.Patient {
		if pa.HealthNumber == healthNumber {
			patientInfo = pa
		}
	}

	fmt.Println(patientInfo)
}

// // checkFile checks to see if the file exists,
// // if it does not, it creates it
func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
