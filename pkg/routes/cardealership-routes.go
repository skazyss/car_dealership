package routes

import (
	"github.com/gorilla/mux"
	"github.com/skazyss/car_dealership/pkg/controllers"
)

var RegisterCarDealershipRoutes = func(r *mux.Router) {
	r.HandleFunc("/car/", controllers.AddCar).Methods("POST")
	r.HandleFunc("/car/", controllers.GetCar).Methods("GET")
	r.HandleFunc("/car/{carId}", controllers.GetCarById).Methods("GET")
	r.HandleFunc("/car/{carId}", controllers.DeleteCar).Methods("DELETE")
	r.HandleFunc("/car/{carId}", controllers.UpdateCar).Methods("PUT")
}
