package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidkristia/go_students_crud_mysql/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterStudentsRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting Server localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

