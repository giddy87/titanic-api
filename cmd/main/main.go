package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/giddy87/titanic-api/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterAppRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server starting on port 9010...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
