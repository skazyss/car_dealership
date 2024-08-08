package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/skazyss/car_dealership/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterCarDealershipRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8888", r))
}
