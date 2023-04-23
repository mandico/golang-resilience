package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "127.0.0.1"
	database = "demo"
	user     = "mandico"
	password = "P@ssw0rd1234"
)

// Initialize connection string.
var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func hostname() string {
	str, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	return str
}

func createTable() {
	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")

	// Drop previous table of same name if one exists.
	drop := "DROP TABLE IF EXISTS rpolog;"
	_, err = db.Exec(drop)
	checkError(err)
	fmt.Println("Finished dropping table (if existed).")

	// Create table.
	create := "CREATE TABLE rpolog (id serial PRIMARY KEY, hostname VARCHAR(50), datetime_local VARCHAR(50), DateUpdated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP);"

	_, err = db.Exec(create)
	checkError(err)
	fmt.Println("Finished creating table.")
}

func main() {

	createTable()
	host := hostname()

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	for {
		loc, _ := time.LoadLocation("America/Sao_Paulo")
		currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05")
		result, err := db.Exec(`INSERT INTO rpolog (hostname, datetime_local) VALUES (?, ?)`, host, currentTime)
		checkError(err)

		if result != nil {
			fmt.Println(result.LastInsertId())
		}
		time.Sleep(1 * time.Second)
	}
}
