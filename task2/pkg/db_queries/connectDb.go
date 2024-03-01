package dbqueries

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	Hostname = "localhost"
	Port     = 5432
	Username = "user"
	Password = "pass"
	Database = "dbname"
)

/*
коннект к постгрес
Hostname = "localhost"
Port     = 5432
Username = "user"
Password = "pass"
Database = "dbname"
*/
func ConnectPostgres() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Hostname, Port, Username, Password, Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
