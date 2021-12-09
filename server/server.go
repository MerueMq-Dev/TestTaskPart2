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
	row := db.QueryRow(`SELECT applicationName, firstParam, secondParam, version
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
	createTable(db, "applications")

	row := db.QueryRow(`SELECT applicationName, firstParam, secondParam, Version
	FROM applications WHERE applicationName in ("` + appFromRequest.AplicationName + `");`)
	var appFromDb App
	row.Scan(&appFromDb.AplicationName, &appFromDb.Param1, &appFromDb.Param2, &appFromDb.Version)

	if appFromDb.AplicationName != appFromRequest.AplicationName {
		stmt, _ := db.Prepare(`INSERT INTO applications (applicationName,firstParam, secondParam, version) values (?, ?, ?, ?)`)
		stmt.Exec(appFromRequest.AplicationName, appFromRequest.Param1, appFromRequest.Param2, 1)
		stmt.Close()
		return
	} else if appFromDb.Param1 != appFromRequest.Param1 || appFromDb.Param2 != appFromRequest.Param2 {
		stamt, _ := db.Prepare(`UPDATE applications SET firstParam = ?, secondParam = ?, Version = Version + 1
	      WHERE applicationName = ?`)
		stamt.Exec(appFromRequest.Param1, appFromRequest.Param2, appFromRequest.AplicationName)
		stamt.Close()
	}
}

func createTable(db *sql.DB, tableName string) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS ` + tableName + `
	(
	applicationName TEXT UNIQUE NOT NULL,
	firstParam INTEGER NOT NULL,
	secondParam TEXT NOT NULL,
	version INTEGER NOT NULL
	)`)
	return err
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/getstate/{appName}", handleGetRequest)
	router.HandleFunc("/savestate", handlePostRequest)
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
