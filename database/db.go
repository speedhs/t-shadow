package database

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

func GetConnectionString() string {
	return "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
}

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", GetConnectionString())
	if err != nil {
		panic(err)
	}
	return db
}

func ExecuteQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

// define a list of allowed query types, where would pass query type like AlterColumn and have a template query for it
var instructionSet = map[string]string{
	"AlterColumnType": "ALTER TABLE %s ALTER COLUMN %s TYPE %s;",
}
