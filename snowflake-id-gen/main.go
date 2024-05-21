package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sony/sonyflake"
	"log"
	"strconv"
	"strings"
	"time"
)

func CreateDBConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "stnduser:stnduser@tcp(127.0.0.1:3309)/snowflakeidgen")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	db, err := CreateDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sf := useSonyflake()

	id, err := sf.NextID()
	fmt.Println("Sonyflake ID:", id)
	id, err = sf.NextID()
	fmt.Println("Sonyflake ID:", id)

	createSnowflakeID()
}

func createSnowflakeID() {
	now := time.Now()
	// Convert to milliseconds since epoch
	epochMillis := now.UnixNano() / int64(time.Millisecond)
	fmt.Println("Epoch Milliseconds:", epochMillis)

	// Check size limits
	if epochMillis >= (1 << 41) {
		fmt.Println("Epoch milliseconds too large to fit in 41 bits")
		return
	}

	machineID := 1000 // Ensure it is less than 2^10
	if machineID >= (1 << 10) {
		fmt.Println("MachineID too large to fit in 10 bits")
		return
	}

	counterID := 2000 // Would be incremented in reality with a mutex
	if counterID >= (1 << 12) {
		fmt.Println("CounterID too large to fit in 12 bits")
		return
	}

	// Construct the snowflakeID
	var snowflakeID int64 = (epochMillis << (22)) // Leave space for 10 + 12 bits
	snowflakeID |= (int64(machineID) << 12)       // Shift left by 12 bits to leave space for the last number
	snowflakeID |= int64(counterID)               // Place last number in the lowest 12 bits

	// Convert snowflakeID to a binary string
	binaryStr := strconv.FormatInt(int64(snowflakeID), 2) // Use FormatUint to avoid two's complement
	// Pad the binary string to ensure it has 64 bits
	binaryStr = fmt.Sprintf("%064s", binaryStr)         // Pad the string with spaces
	binaryStr = strings.ReplaceAll(binaryStr, " ", "0") // Replace spaces with zeros

	fmt.Printf("Stored snowflakeID: %d\n", snowflakeID)
	fmt.Printf("Binary Representation: %s\n", binaryStr)
}

func useSonyflake() *sonyflake.Sonyflake {
	var st sonyflake.Settings
	sf := sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("failed to create Sonyflake")
	}
	return sf
}
