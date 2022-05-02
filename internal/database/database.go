package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func start() {
	var dsn string = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_BASE")
	var err error
	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("error connecting to mysql: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("database connection established\n")

	err = conn.Ping()
	if err != nil {
		fmt.Printf("couldn't ping database: %s\n", err.Error())
		os.Exit(1)
	}

	_, err = conn.Exec("SET lc_time_names = 'es_ES';")
	if err != nil {
		fmt.Printf("error configuring database locale: %s\n", err.Error())
		os.Exit(1)
	}
	_, err = conn.Exec("SET @@session.time_zone='+00:00';")
	if err != nil {
		fmt.Printf("error configuring database timezone: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("database configured\n")
	db = conn
}

func GetDB() *sqlx.DB {
	if db == nil {
		start()
	}
	return db
}
