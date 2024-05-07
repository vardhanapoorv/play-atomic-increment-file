package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	DB *sql.DB
}

type Response struct {
	StartID int `json:"startid"`
	EndID   int `json:"endid"`
}

func main() {

	db, err := createDBConn()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		panic(err)
	}
	defer db.Close()
	// Pass db connection to reserveAgentHandler
	app := &App{DB: db}

	r := mux.NewRouter()
	r.HandleFunc("/getbatchIds", app.GetBatchIDHandler)
	fmt.Println("Server listening on port 8081...")
	http.ListenAndServe(":8081", r)
}

func createDBConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "stnduser:stnduser@tcp(127.0.0.1:3309)/amazonid")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (a *App) GetBatchIDHandler(w http.ResponseWriter, r *http.Request) {

	var requestData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get value of the "orderId" key
	batchSiz, ok := requestData["batchsize"].(float64)
	if !ok {
		fmt.Println("Error: key not found ")
		http.Error(w, "key not found or not a string", http.StatusBadRequest)
		return
	}
	serviceName, ok := requestData["servicename"].(string)

	if !ok {
		http.Error(w, "key not found or not a string", http.StatusBadRequest)
		return
	}
	batchSize := int(batchSiz)

	startID, endID, err := BatchIdGen(a.DB, batchSize, serviceName)
	if err != nil {
		fmt.Errorf("Error generating batch ids: %v", err)
	}

	// json response to client with range of ids
	response := Response{
		StartID: startID,
		EndID:   endID,
	}

	// Marshal the Response struct to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response writer
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON response:", err)
		return
	}
}

func BatchIdGen(db *sql.DB, batchSize int, serviceName string) (int, int, error) {
	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, 0, err
	}

	// Query to get a non-reserved food packet - With a exclusive lock to ensure no other transaction can access the same record
	query := "SELECT counter FROM amazonidgen WHERE servicename = ? FOR UPDATE"
	row := tx.QueryRow(query, serviceName)

	var counter int
	if err := row.Scan(&counter); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No results found.")
		}
		return 0, 0, err
	}

	// Update the record to mark it as reserved
	updateQuery := "UPDATE amazonidgen SET counter = ?  WHERE servicename = ?"
	_, err = tx.Exec(updateQuery, counter+batchSize, serviceName)
	if err != nil {
		tx.Rollback()
	}
	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return 0, 0, err
	}

	return counter, counter + batchSize, nil
}
