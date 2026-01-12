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

// a list of allowed query types, where would pass query type like AlterColumn and have a template query for it
var create_shadow_table string = "CREATE TABLE shadow_%s AS TABLE %s;"
var alter_column_type string = "ALTER TABLE %s ALTER COLUMN %s TYPE %s;"
var drop_og_table string = "DROP TABLE %s;"
var rename_shadow_table string = "ALTER TABLE shadow_%s RENAME TO %s;"
var add_column string = "ALTER TABLE %s ADD COLUMN %s %s;"

var instructions = map[string]string{
	"AlterColumn": alter_column_type,
	"AddColumn":   add_column,
}

func SetupReplication(db *sql.DB, table_name string) error {
	return nil
}
