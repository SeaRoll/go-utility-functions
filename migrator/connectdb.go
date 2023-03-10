package migrator

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // required for mysql
	_ "github.com/jackc/pgx/v4/stdlib" // required for postgres
	_ "github.com/mattn/go-sqlite3"    // required for sqlite
)

// connects to a postgres database
// connectionString is a string in the format:
// "postgres://user:password@host:port/database"
func ConnectToPostgres(host string, port string, username string, password string, dbname string) *sql.DB {
	conn, err := sql.Open("pgx", "postgres://"+username+":"+password+"@"+host+":"+port+"/"+dbname)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return conn
}

// connects to a sqlite database
func ConnectToSqlite(path string) *sql.DB {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return conn
}

// connects to a mysql database
func ConnectToMysql(host string, port string, username string, password string, dbname string) *sql.DB {
	conn, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return conn
}
