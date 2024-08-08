package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/skazyss/car_dealership/pkg/models"
	"github.com/skazyss/car_dealership/pkg/utils"
)

var NewCar models.Car

func GetCar(w http.ResponseWriter, r *http.Request) {
	newCars := models.GetAllCars()
	res, _ := json.Marshal(newCars)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["carId"]
	Id, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}
	carDetails, _ := models.GetCarById(Id)
	res, _ := json.Marshal(carDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	AddCar := &models.Car{}
	utils.ParseBody(r, AddCar)
	c := AddCar.AddCar()
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["carId"]
	Id, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}
	car := models.DeleteCar(Id)
	res, _ := json.Marshal(car)
	w.Header().Set("Content-Type", "pkgcation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	var updCar = &models.Car{}
	utils.ParseBody(r, updCar)
	vars := mux.Vars(r)
	carId := vars["carId"]
	Id, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	carDetails, db := models.GetCarById(Id)
	if updCar.CarMake != "" {
		carDetails.CarMake = updCar.CarMake
	}
	if updCar.CarModel != "" {
		carDetails.CarModel = updCar.CarModel
	}
	if updCar.Year != "" {
		carDetails.Year = updCar.Year
	}
	db.Save(&carDetails)
	res, _ := json.Marshal(carDetails)
	w.Header().Set("Content-Type", "pkgcation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
