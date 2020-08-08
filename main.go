package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

func startupServer() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	http.HandleFunc("/home", home)
	http.HandleFunc("/getMasterPlan", getMasterPlan)
	http.HandleFunc("/getDates", getDates)
	http.HandleFunc("/getSortedByWBS", getSortedWBS)
	http.HandleFunc("/getSortedByDates", getSortedByDates)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Hello, The server is running")
}

func getMasterPlan(w http.ResponseWriter, r *http.Request) {

}

func getDates(w http.ResponseWriter, r *http.Request) {

}

func getSortedWBS(w http.ResponseWriter, r *http.Request) {

}

func getSortedByDates(w http.ResponseWriter, r *http.Request) {

}

func main() {

	startupServer()

}
