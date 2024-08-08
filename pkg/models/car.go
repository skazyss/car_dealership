package models

import (
	"github.com/jinzhu/gorm"
	"github.com/skazyss/car_dealership/pkg/config"
)

var db *gorm.DB

type Car struct {
	gorm.Model
	CarMake  string `gorm:""json:"carmake"`
	CarModel string `json:"carmodel"`
	Year     string `json:"year"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Car{})
}

func (c *Car) AddCar() *Car {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllCars() []Car {
	var Cars []Car
	db.Find(&Cars)
	return Cars
}

func GetCarById(Id int64) (*Car, *gorm.DB) {
	var getCar Car
	db := db.Where("ID=?", Id).Find(&getCar)
	return &getCar, db
}

func DeleteCar(Id int64) Car {
	var car Car
	db.Where("ID=?", Id).Delete(car)
	return car
}
