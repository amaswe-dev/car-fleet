package main

import (
	"amaswe-be/service"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	runServer(":8080")
}

func runServer(port string) {

	r := mux.NewRouter()

	r.HandleFunc("/getfleet", getFleetHandler).Methods("GET")

	//TODO
	r.HandleFunc("/getvehicle/{id}", getVehicleHandler).Methods("GET")

	//TODO
	r.HandleFunc("/metrics", getMetrics).Methods("POST")

	http.ListenAndServe(port, r)
}

func getFleetHandler(w http.ResponseWriter, r *http.Request) {
	res, err := service.GetFleet()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(res)
}

func getVehicleHandler(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)

	//TODO
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	_, _ = io.ReadAll(r.Body)

	//TODO
}
