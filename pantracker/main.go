// Author: Richard Xiong

package main

import (
	"net/http"
	"pantracker/controllers"
	"pantracker/models"
	"pantracker/utilities"

	"github.com/gorilla/mux"
)

// import "pantracker/models"

// "encoding/json"

func main() {
	patientDB := models.NewPatientDB("db.json")
	c := controllers.NewRepository(patientDB)

	r := mux.NewRouter()

	frameworks := http.FileServer(http.Dir("./frameworks/"))
	r.PathPrefix("/frameworks/").Handler(http.StripPrefix("/frameworks/", frameworks))

	r.Handle("/", c.Registration).Methods("GET")
	r.HandleFunc("/register", c.NewPatient).Methods("Post")

	r.Handle("/complete", c.Completed).Methods("GET")

	address := utilities.GetConfig("config.json")
	http.ListenAndServe(address, r)
}
