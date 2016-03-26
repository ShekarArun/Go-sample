package main

import (
"log"
"net/http"
"github.com/gorilla/mux"
"github.com/Pruthvik-n/Go-sample/api/v1/controllers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sign_up",controllers.Register.Create).Methods("POST")
	r.HandleFunc("/fetch/{id:[0-9]+}",controllers.Register.Fetch).Methods("GET")
	r.HandleFunc("/delete/{id:[0-9]+}",controllers.Register.Delete).Methods("DELETE")
	r.HandleFunc("/fetch_all",controllers.Register.FetchAll).Methods("GET")
	http.Handle("/", r)
	log.Println("main: Bombs Away on Port 3000 !")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}