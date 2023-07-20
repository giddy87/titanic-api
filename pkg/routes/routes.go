package routes

import (
	"github.com/giddy87/titanic-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAppRoutes = func(router *mux.Router) {
	router.HandleFunc("/passenger", controllers.CreatePassenger).Methods("POST")
	router.HandleFunc("/passenger", controllers.GetAllPassengers).Methods("GET")
	router.HandleFunc("/passenger/{PassId}", controllers.GetPassengerById).Methods("GET")
	router.HandleFunc("/passenger/{PassId}", controllers.DeletePassenger).Methods("DELETE")
	router.HandleFunc("/passenger/{PassId}", controllers.UpdatePassenger).Methods("PUT")

}
