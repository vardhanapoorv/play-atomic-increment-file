package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	StartID int `json:"startid"`
	EndID   int `json:"endid"`
}

func main() {
	requestData := map[string]interface{}{"batchsize": 500, "servicename": "payments"} // Replace 123 with the actual order ID

	// Convert the request data to JSON
	payload, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	response, err := http.Post("http://localhost:8081/getbatchIds", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error booking food packet:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error booking food packet:", response.Status)
		return
	}
	var res Response

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		fmt.Println("Error decoding response:", err)
	}
	fmt.Println("Batch IDs", res.StartID, res.EndID)

}
