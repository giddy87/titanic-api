package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/giddy87/titanic-api/pkg/models"
	"github.com/giddy87/titanic-api/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllPassengers(w http.ResponseWriter, r *http.Request) {
	AllPassengers := models.GetAllPassengers()
	result, _ := json.Marshal(AllPassengers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetPassengerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	passId := vars["PassId"]
	ID, err := strconv.ParseInt(passId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	passData, _ := models.GetPassengerById(ID)
	result, _ := json.Marshal(passData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func CreatePassenger(w http.ResponseWriter, r *http.Request) {
	createPass := &models.Passenger{}
	utils.ParseBody(r, createPass)
	p := createPass.CreatePassenger() //CreatePassenger() is the Passenger models Method
	result, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func DeletePassenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//mux.Vars(r) retrieve the map of route variables for the request r
	passId := vars["passId"] //passId is the value of key passId in the map mux.Vars(r)
	ID, err := strconv.ParseInt(passId, 0, 0)
	if err != nil {
		fmt.Println("Parse error")
	}
	result := models.DeletePassenger(ID)
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdatePassenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UpdateData := &models.Passenger{}
	utils.ParseBody(r, UpdateData)
	passId := vars["PassId"]
	ID, err := strconv.ParseInt(passId, 0, 0)
	if err != nil {
		fmt.Println("Parse error")
	}
	original_record, db := models.GetPassengerById(ID)
	if UpdateData.Survived != 0 {
		original_record.Survived = UpdateData.Survived
	}
	if UpdateData.Pclass != "" {
		original_record.Pclass = UpdateData.Pclass
	}
	if UpdateData.Name != "" {
		original_record.Name = UpdateData.Name
	}
	if UpdateData.Age != 0 {
		original_record.Age = UpdateData.Age
	}
	if UpdateData.Sex != "" {
		original_record.Sex = UpdateData.Sex
	}
	if UpdateData.Siblings != "" {
		original_record.Siblings = UpdateData.Siblings
	}
	if UpdateData.Children != "" {
		original_record.Children = UpdateData.Children
	}
	if UpdateData.Fare != 0.0 {
		original_record.Fare = UpdateData.Fare
	}
	db.Save(&original_record)
	res, _ := json.Marshal(original_record)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
