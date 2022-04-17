package app

//package main

import (
	"fmt"
	"github.com/fprogress17/banking/domain"
	"github.com/fprogress17/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define router
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	/*
		router.HandleFunc("/greet", greet).Methods(http.MethodGet)
		router.HandleFunc("/customers", GetAllCustomer).Methods(http.MethodGet)
		router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

		router.HandleFunc("/customers",createCustomer).Methods(http.MethodPost)
	*/

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
