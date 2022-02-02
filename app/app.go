package app

import (
	"github.com/6adfeniks/banking/domain"
	"github.com/6adfeniks/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe("localhost: 8000", router))

}
