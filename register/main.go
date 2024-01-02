package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

const (
	name     = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "db_sms"
)

var (
	id           int
	username     string
	userpassword string
	usertype     string
)

func dns() string {
	dataDriver := fmt.Sprintf("%s:%s@tcp(%s)/%s", name, password, hostname, dbname)
	return dataDriver
}
func data() {

	db, err := sql.Open("mysql", dns())

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	//row := "SELECT * FROM users"
	rows, _ := db.Query("SELECT * FROM users	")
	fmt.Printf("id \t username \t userpassword \tusertype\n")
	for rows.Next() {

		if err := rows.Scan(&id, &username, &userpassword, &usertype); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d \t %s \t %s \t %s \n", id, username, userpassword, usertype)
	}
	db.Close()
	fmt.Print("\nconnection closed")
}

var tmp *template.Template

func main() {

	tmp, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/register", registerUser)
	http.HandleFunc("/registerForm", register)
	http.ListenAndServe(":8080", nil)
	// Fixed typo in the path
	data()
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "register.html", 0)
}

type dataForm struct {
	Name string
	Ps   string
}

func register(w http.ResponseWriter, r *http.Request) {
	data()
	// var nm dataForm
	// nm.Name = r.FormValue("username")
	// nm.Ps = r.FormValue("password")
	// dataBase, err := sql.Open("mysql", dns())
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// userType := "student"
	// query := "INSERT INTO users (username, userpassword, usertype) VALUES (?, ?,?) "
	// _, err = dataBase.Exec(query, nm.Name, nm.Ps, userType)
	// if err != nil {
	fmt.Fprint(w, "helllo")
	// }
}
