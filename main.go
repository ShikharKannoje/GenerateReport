package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/sqltocsv"
	_ "github.com/joho/sqltocsv"
	_ "github.com/lib/pq"
)

const (
	hostname     = "localhost"
	hostport     = 5432
	username     = "postgres"
	password     = "root"
	databasename = "ConstructionPlan"
)

func startupServer() {

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Println("The server is up")
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
	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, hostport, username, password, databasename)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Database connected")
	}

	rows, err := db.Query(`SELECT * FROM public.masterplan ORDER BY string_to_array(slno, '.')::int[]`)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error Fetching from DB"))
		return
	}
	log.Println(rows)
	err = sqltocsv.WriteFile("./report.csv", rows)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	w.Header().Set("Content-type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")
	log.Println("done write till here")
	sqltocsv.Write(w, rows)
	return
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
