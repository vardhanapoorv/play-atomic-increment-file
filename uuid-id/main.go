package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"time"
)

func CreateDBConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "stnduser:stnduser@tcp(127.0.0.1:3309)/amazonid")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Connect to MySQL database
	db, err := CreateDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert 1 million rows into the table_with_uuid table
	fmt.Println("Inserting 1 million rows into table_with_uuid_char...")
	startTimeUUIDChar := time.Now()
	for i := 1; i <= 1000000; i++ {
		uuid := uuid.New().String()
		_, err := db.Exec("INSERT INTO table_with_uuid_char (ID, age) VALUES (?, ?)", uuid, i)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Time spent for UUID Char", time.Since(startTimeUUIDChar))
	fmt.Println("Successfully inserted 1 million rows into table_with_uuid_char.")

	// Insert 1 million rows into the table_with_auto_increment table
	fmt.Println("Inserting 1 million rows into table_with_auto_increment...")
	startTimeAutoIncrement := time.Now()
	for i := 1; i <= 1000000; i++ {
		_, err := db.Exec("INSERT INTO table_with_auto_increment (age) VALUES (?)", i)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Time spent for int AutoIncrement", time.Since(startTimeAutoIncrement))
	fmt.Println("Successfully inserted 1 million rows into table_with_auto_increment.")

	fmt.Println("Inserting 1 million rows into table_with_uuid...")
	startTimeUUIDBinary := time.Now()
	for i := 1; i <= 1000000; i++ {
		uuid := uuid.New()
		// Convert UUID to binary format
		uuidBinary := uuid[:]
		_, err := db.Exec("INSERT INTO table_with_uuid (ID, age) VALUES (?, ?)", uuidBinary, i)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Time spent for UUID Binary", time.Since(startTimeUUIDBinary))
	fmt.Println("Successfully inserted 1 million rows into table_with_uuid.")
}
