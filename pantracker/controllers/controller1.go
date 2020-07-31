// Author: Richard Xiong

package controllers

import (
	"log"
	"net/http"
	"pantracker/models"
	"pantracker/views"

	"github.com/gorilla/schema"
)

// Repository is used to update views
// and update information for the the models
type Repository struct {
	Registration *views.View
	Completed    *views.View
	patientDB    *models.PatientDB
}

// NewRepository creates a Repository controller
func NewRepository(pdb *models.PatientDB) *Repository {

	return &Repository{
		patientDB:    pdb,
		Registration: views.NewView("pages/patient.html"),
		Completed:    views.NewView("pages/completed.html"),
	}
}

// NewPatient handles POST request when the patient clicks
// on Register on the registration form, and a new patient is added
// to the system
func (rep *Repository) NewPatient(w http.ResponseWriter, r *http.Request) {

	var patientInfo PatientInformation // create a Patient struct to store data from the registration form

	// getFormData obtains the fields from the form and
	// populates the PatientInformation struct
	if err := getFormData(r, &patientInfo); err != nil {
		log.Println(err)
	}

	patient := models.Patient{
		FirstName:    patientInfo.FirstName,
		LastName:     patientInfo.LastName,
		Email:        patientInfo.Email,
		Password:     patientInfo.Password,
		Phone:        patientInfo.Phone,
		HealthNumber: patientInfo.HealthNumber,
		Address: models.Address{
			StreetNumber: patientInfo.StreetNumber,
			StreetName:   patientInfo.StreetName,
			City:         patientInfo.City,
			Province:     patientInfo.Province,
			PostalCode:   patientInfo.PostalCode,
		},
	}

	rep.patientDB.AddPatient(patient)
	http.Redirect(w, r, "/complete", 301) // go to completion to tell patient they have finished registering
}

// getFormData takes a request pointer and uses it to obtain
// data from the form and parses that data into a struct
// in an order dictated by the schema labels of that struct
func getFormData(r *http.Request, form interface{}) error {

	if err := r.ParseForm(); err != nil {
		return err
	}

	d := schema.NewDecoder()
	if err := d.Decode(form, r.PostForm); err != nil {
		return err
	}

	return nil
}

// PatientInformation is used to store
// form data for registering a patient
type PatientInformation struct {
	FirstName    string `schema:"firstName"`
	LastName     string `schema:"lastName"`
	Email        string `schema:"email"`
	Password     string `schema:"password"`
	Phone        int    `schema:"phone"`
	HealthNumber int    `schema:"healthNumber"`
	StreetNumber int    `schema:"streetNumber"`
	StreetName   string `schema:"streetName"`
	City         string `schema:"city"`
	Province     string `schema:"province"`
	PostalCode   string `schema:"postalCode"`
}

// OpenRepository is called when an Employee
// accesses the repository page. It will signal
// Views layer to render an HTML page with a list of
// patients who have registered for a test
func (rep *Repository) OpenRepository(w http.ResponseWriter, r *http.Request) {
	patientList := rep.patientDB.GetPatients()
	v := views.NewView("pages/repository.html")
	v.RenderTemplate(w, struct {
		PatientList []models.Patient
	}{
		PatientList: patientList,
	})
}
