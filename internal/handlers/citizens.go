package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adityasshinde/Golang-Backend/internal/models"
	"github.com/adityasshinde/Golang-Backend/pkg/db"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func GetCitizens(w http.ResponseWriter, r *http.Request) {
	citizens, err := db.GetAllCitizens()
	if err != nil {
		http.Error(w, "Error fetching citizens", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(citizens)
}

func GetCitizen(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	citizenID := params["id"]
	citizen, err := db.GetCitizenByID(citizenID)
	if err != nil {
		http.Error(w, "Error fetching citizen", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(citizen)
}

func CreateCitizen(w http.ResponseWriter, r *http.Request) {
	var citizen models.NewCitizen
	json.NewDecoder(r.Body).Decode(&citizen)
	//generate Id for the citizen
	uuId := uuid.NewV4().String()
	newCitizen := models.Citizen{
		Id:          uuId,
		FirstName:   citizen.FirstName,
		LastName:    citizen.LastName,
		DateOfBirth: citizen.DateOfBirth,
		Gender:      citizen.Gender,
		Address:     citizen.Address,
		City:        citizen.City,
		State:       citizen.State,
		Pincode:     citizen.Pincode,
	}
	fmt.Fprint(w, citizen)
	err := db.CreateCitizen(newCitizen)
	if err != nil {
		http.Error(w, "Error creating citizen", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UpdateCitizen(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	citizenID := params["id"]
	var updatedCitizen models.Citizen
	json.NewDecoder(r.Body).Decode(&updatedCitizen)
	err := db.UpdateCitizen(citizenID, updatedCitizen)
	if err != nil {
		http.Error(w, "Error updating citizen", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteCitizen(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	citizenID := params["id"]
	fmt.Println(" Deleting citizen with id: ", citizenID)
	err := db.DeleteCitizen(citizenID)
	if err != nil {
		http.Error(w, "Error deleting citizen", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
