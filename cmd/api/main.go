package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

type Driver struct {
	gorm.Model

	Name string

	License string

	Cars []Car
}

type Car struct {
	gorm.Model

	Year int

	Make string

	ModelName string

	DriverID int
}

var db *gorm.DB

var err error

var (
	drivers = []Driver{

		{Name: "Jimmy Johnson", License: "ABC123"},

		{Name: "Howard Hills", License: "XYZ789"},

		{Name: "Craig Colbin", License: "DEF333"},
	}

	cars = []Car{

		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},

		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},

		{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},

		{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
	}
)

func main() {

	log.Println("PostgreSQL...")

	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")

	if err != nil {

		panic("failed to connect database")

	}

	defer db.Close()

	db.AutoMigrate(&Driver{})

	db.AutoMigrate(&Car{})

	for index := range cars {

		db.Create(&cars[index])

	}

	for index := range drivers {

		db.Create(&drivers[index])

	}

	router.HandleFunc("/cars", GetCars).Methods("GET")

	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")

	router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")

	router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
