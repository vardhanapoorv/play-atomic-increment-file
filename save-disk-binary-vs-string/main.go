package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	saveCounter(-42)
	saveTime := time.Since(start)
	fmt.Println("Save time:", saveTime)
	readStart := time.Now()
	loadCounter()
	readTime := time.Since(readStart)
	fmt.Println("Read time:", readTime)

	startBinary := time.Now()
	writeInBinaryToFile(421)
	writeBinaryTime := time.Since(startBinary)
	fmt.Println("Write binary time:", writeBinaryTime)

	startReadBinary := time.Now()
	readIntValue := readIntegerFromFile("outputt.bin")
	readBinaryTime := time.Since(startReadBinary)
	fmt.Println("Read binary time:", readBinaryTime)
	fmt.Println("Read integer value:", readIntValue)

}

// Save counter value to disk
func saveCounter(value int) {
	// Integer value to save
	intValue := value

	// Convert integer to string
	start := time.Now()
	intString := strconv.Itoa(intValue)
	fmt.Println("Time spent converting integer to string", time.Since(start))

	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Close the file when done

	// Write string to file
	var b int
	start = time.Now()
	b, err = file.WriteString(intString)
	fmt.Println("Time spent writing string to file", time.Since(start))
	fmt.Println("Number of bytes written:", b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Integer value saved to file successfully.") // Write string to file
}

// Load counter value from disk
func loadCounter() {

	content, err := ioutil.ReadFile("output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parse content as integer
	intString := string(content)
	intValue, err := strconv.Atoi(intString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Integer value read from file:", intValue)
}

func writeInBinaryToFile(value int) {
	intValue := value

	// Open file for writing
	file, err := os.OpenFile("outputt.bin", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Close the file when done

	// Convert integer to binary encoded form
	intBytes := make([]byte, 8) // 8 bytes to hold an int64
	start := time.Now()
	binary.LittleEndian.PutUint64(intBytes, uint64(intValue))
	fmt.Println("Time spent converting integer to binary", time.Since(start))

	// Write binary encoded integer to file
	var b int
	start = time.Now()
	b, err = file.Write(intBytes)
	fmt.Println("Time spent writing binary to file", time.Since(start))
	fmt.Println("Number of bytes written:", b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Integer value saved to file successfully.")
}

func readIntegerFromFile(filename string) int {
	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer file.Close() // Close the file when done

	// Read binary encoded integer from file
	intBytes := make([]byte, 8) // 8 bytes to hold an int64
	_, err = file.Read(intBytes)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	// Convert binary encoded form to integer
	intValue := int(binary.LittleEndian.Uint64(intBytes))

	return intValue
}
