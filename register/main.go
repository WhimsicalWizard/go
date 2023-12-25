package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "db_sms"
)

func dns() string {
	dataDriver := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	return dataDriver
}
func main() {

	db, err := sql.Open("mysql", dns())

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	table := "SELECT * FROM users"
	rows, _ := db.Query(table)
	for rows.Next() {
		col, _ := rows.Columns()
		fmt.Print(col, "\n")
	}
	db.Close()
	fmt.Print("\nconnection closed")
}
