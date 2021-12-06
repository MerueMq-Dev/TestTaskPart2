package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	AplicationName string `json:"AplicationName"`
	Param1         int    `json:"Param1"`
	Param2         string `json:"Param2"`
	Version        int    `json:"Version"`
}

func handleGetRequest(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	appName := params["appName"]
	db, err := sql.Open("sqlite3", "./apps.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row := db.QueryRow(`SELECT applicationName, firstParam, secondParam, Version
	FROM applications WHERE applicationName in ("` + appName + `");`)

	var appFromDb App
	errs := row.Scan(&appFromDb.AplicationName, &appFromDb.Param1, &appFromDb.Param2, &appFromDb.Version)
	if errs != nil {
		log.Fatal(errs)
	}
	json.NewEncoder(writer).Encode(appFromDb)
}

func handlePostRequest(writer http.ResponseWriter, request *http.Request) {

	var appFromRequest App
	json.NewDecoder(request.Body).Decode(&appFromRequest)

	db, _ := sql.Open("sqlite3", "./apps.db")
	defer db.Close()
	stmt, _ := db.Prepare(`INSERT INTO applications (applicationName,firstParam, secondParam, version) values (?, ?, ?, ?)`)
	stmt.Exec(appFromRequest.AplicationName, appFromRequest.Param1, appFromRequest.Param2, 1)

	row := db.QueryRow(`SELECT applicationName, firstParam, secondParam, Version
	FROM applications WHERE applicationName in ("` + appFromRequest.AplicationName + `");`)
	var appFromDb App
	row.Scan(&appFromDb.AplicationName, &appFromDb.Param1, &appFromDb.Param2, &appFromDb.Version)
	if appFromRequest.Param1 != appFromDb.Param1 || appFromRequest.Param2 != appFromDb.Param2 {
		stamt, _ := db.Prepare(`UPDATE applications SET firstParam = ?, secondParam = ?, Version = Version + 1
	    WHERE applicationName=?`)
		stamt.Exec(appFromRequest.Param1, appFromRequest.Param2, appFromRequest.AplicationName)
		stamt.Close()
	}

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/getstate/{appName}", handleGetRequest)
	router.HandleFunc("/savestate", handlePostRequest)
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

// func createTable(){

// 	db, _ := sql.Open("sqlite3", "./app.db")

// 	stmt, _ := db.Prepare(`CREATE TABLE applications(
// 		id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
// 		applicationName TEXT UNIQUE NOT NULL,
// 		firstParam INTEGER NOT NULL,
// 		secondParam TEXT NOT NULL,
// 		Version INTEGER NOT NULL
// 	  )`)

// 	stmt.Exec()
// }
