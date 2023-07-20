package models

import (
	"github.com/giddy87/titanic-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Passenger struct {
	gorm.Model
	Survived int8    `json:"survived"`
	Pclass   string  `json:"pclass"`
	Name     string  `gorm:""json:"name"`
	Sex      string  `json:"sex"`
	Age      int8    `json:"age"`
	Siblings string  `json:"siblings"`
	Children string  `json:"children"`
	Fare     float64 `json:"fare"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Passenger{})
}

func (p *Passenger) CreatePassenger() *Passenger {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllPassengers() []Passenger {
	var Passengers []Passenger
	db.Find(&Passengers)
	return Passengers
}

func GetPassengerById(ID int64) (*Passenger, *gorm.DB) {
	var getPassenger Passenger
	db := db.Where("ID=?", ID).Find(&getPassenger)
	return &getPassenger, db
}

func DeletePassenger(Id int64) Passenger {
	var pass Passenger
	db.Where("ID = ?", Id).Delete(pass)
	return pass
}
