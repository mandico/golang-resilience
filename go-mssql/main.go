package main

// 1
import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
	"os"
	"time"

	"github.com/go-co-op/gocron"

	log "github.com/sirupsen/logrus"
)

var db *sql.DB

var server = "localhost"
var port = 1433
var user = "sa"
var password = "P@ssw0rd1234"
var database = "master"

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
}

func main() {
	runCronJobs()
}

func message() {
	// Current Time
	currentTime := time.Now()

	// Get Hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	message := fmt.Sprintf("| %v | %v ", currentTime.Format("2006.01.02 15:04:05.000000"), hostname)
	log.Info("* " + message)
}

func healthCheckDB() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Warn("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Warn(err.Error())
	} else {
		log.Info("Connected!")
		CreateLog()
	}

}

func runCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(500).Milliseconds().Do(func() {
		healthCheckDB()
	})
	s.StartBlocking()
}

func CreateLog() (int64, error) {
	// Get Hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/**

		query := `CREATE TABLE IF NOT EXISTS Logs(id int primary key auto_increment, hostname text,
	        created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
		resultado, e := db.Exec(query)
		if e != nil {
			fmt.Println("Error create table " + err.Error())
			return -1, err
		}
		log.Info(resultado)
		***/

	tsql := fmt.Sprintf("INSERT INTO dbo.Logs (Hostname) VALUES ('%s');",
		hostname)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}
