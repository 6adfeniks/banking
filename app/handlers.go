package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/6adfeniks/banking/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Nikolay", "Minsk", "222228"},
	//	{"Ivan", "Zhodino", "222167"},
	//}

	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
