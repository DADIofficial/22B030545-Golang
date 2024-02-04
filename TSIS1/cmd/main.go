package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/Total_Drama", Total_Drama).Methods("GET")
	router.HandleFunc("/Total_Drama/Gender/{Gender}", find_gender).Methods("GET")
	router.HandleFunc("/Total_Drama/Favteam/{Favteam}", find_team).Methods("GET")
	router.HandleFunc("/Total_Drama/Hair/{Hair}", find_hair).Methods("GET")
	router.HandleFunc("/Total_Drama/Position/{Position}", find_hair).Methods("GET")
	router.HandleFunc("/Total_Drama/Power/{Power}", find_power).Methods("GET")
	router.HandleFunc("/Total_Drama/Number_seasons/{Number_seasons}", find_season).Methods("GET")
	router.HandleFunc("/Total_Drama/{Name}", find_char).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
