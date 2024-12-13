package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

var server = "DESKTOP-3A2FIBS"
var port = 1433
var database = "WorkingGo"

func Connect() {
	fmt.Println("Connectig to database")
}
