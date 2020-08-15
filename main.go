// MasterPlan csv/excel report generation Service API
//
// .
//
//     Schemes: http, https
//     Host: localhost:8080
//     Version: 0.1.0
//     basePath: /
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"database/sql"
	"encoding/json"
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

// WriteJSONResponse represents a utility function which writes status code and JSON to response
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func startupServer() {

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Println("The server is up")
	http.HandleFunc("/", home)
	http.HandleFunc("/getActivities", getActivities)
	http.HandleFunc("/getAllDetails", getAllDetails)
	http.HandleFunc("/getSortedByWBS", getSortedWBS)
	http.HandleFunc("/getSortedByDatesAndWBS", getSortedByDatesAndWBS)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Hello, The server is running")
}

func getActivities(w http.ResponseWriter, r *http.Request) {
	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, hostport, username, password, databasename)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, 500, "database connection failed")
		return
	} else {
		log.Println("Database connected")
	}

	rows, err := db.Query(`SELECT activity FROM public.masterplan`)
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, 500, "database connection issue")
		//w.Write([]byte("Error Fetching from DB"))
		return
	}
	log.Println(rows)
	// err = sqltocsv.WriteFile("./report.csv", rows)
	// if err != nil {
	// 	panic(err)
	// }
	defer rows.Close()

	w.Header().Set("Content-type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")
	log.Println("done write till here")
	sqltocsv.Write(w, rows)

	return
}

func getAllDetails(w http.ResponseWriter, r *http.Request) {
	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, hostport, username, password, databasename)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, 500, "database connection issue")
		return
	} else {
		log.Println("Database connected")
	}

	rows, err := db.Query(`SELECT * FROM public.masterplan`)
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, 500, "database connection issue")
		return
	}
	log.Println(rows)
	// err = sqltocsv.WriteFile("./report.csv", rows)
	// if err != nil {
	// 	panic(err)
	// }
	defer rows.Close()

	w.Header().Set("Content-type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")
	log.Println("done write till here")
	sqltocsv.Write(w, rows)

	return
}

func getSortedWBS(w http.ResponseWriter, r *http.Request) {
	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, hostport, username, password, databasename)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, 500, "database connection issue")
		return
	} else {
		log.Println("Database connected")
	}

	rows, err := db.Query(`SELECT * FROM public.masterplan ORDER BY string_to_array(slno, '.')::int[]`)
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, 500, "database connection issue")
		return
	}
	log.Println(rows)
	// err = sqltocsv.WriteFile("./report.csv", rows)
	// if err != nil {
	// 	panic(err)
	// }
	defer rows.Close()

	w.Header().Set("Content-type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")
	log.Println("done write till here")
	sqltocsv.Write(w, rows)

	return
}

func getSortedByDatesAndWBS(w http.ResponseWriter, r *http.Request) {
	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, hostport, username, password, databasename)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		WriteJSONResponse(w, 500, "database connection issue")
		return
	} else {
		log.Println("Database connected")
	}

	rows, err := db.Query(`SELECT startdate, slno, activity FROM (SELECT startdate, slno, activity FROM masterplan ORDER BY string_to_array(slno, '.')::int[])subtable`)
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, 500, "database connection issue")
		return
	}
	log.Println(rows)
	defer rows.Close()

	w.Header().Set("Content-type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")
	log.Println("done write till here")
	sqltocsv.Write(w, rows)

	return
}

func main() {

	startupServer()

}
