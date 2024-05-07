package main

import (
	"encoding/binary"
	"fmt"
	"os"
	//"sync"
	"time"
)

func main() {
	startBinary := time.Now()

	//writeInBinaryToFile(0)
	// Total time: 3.322291ms
	//var wg sync.WaitGroup
	//var mu sync.Mutex
	var value int32 = readIntegerFromFile("outputt.bin") // 1 here is for safety, let's say id generator returned 100 but didn't write to file, so we will start from 101 to avoid ID collision
	fmt.Println("Value:", value)
	fmt.Println("New safest starting value:", value+100+1)
	/*for i := 0; i < 1000; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			value++
			// Lets write to file not every time instead after every few values
			if value%100 == 0 {
				if value == 200 {
					panic("Panic")
				}
				writeInBinaryToFile(int(value))
				//fmt.Println("Value:", value)
			}
			fmt.Println("Id: ", value)
		}()
	}*/
	//wg.Wait()
	totalTime := time.Since(startBinary)
	fmt.Println("Total time:", totalTime)
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
	binary.LittleEndian.PutUint64(intBytes, uint64(intValue))

	// Write binary encoded integer to file
	_, err = file.Write(intBytes)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Integer value saved to file successfully.")
}

func readIntegerFromFile(filename string) int32 {
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
	intValue := int32(binary.LittleEndian.Uint64(intBytes))

	return intValue
}
