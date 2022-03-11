package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Convert JSON to Go Struct
// https://mholt.github.io/json-to-go/
type Ticker struct {
	Last      float64 `json:"last"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
	Timestamp int64   `json:"timestamp"`
}

func main() {
	fmt.Println("Start!")

	url := "https://coincheck.com/api/ticker"

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		switch err2 := err.(type) {
		default:
			fmt.Println("Error Request:", err2)
		}
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error Response:", resp.Status)
		return
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	var ticker Ticker
	json.Unmarshal(body, &ticker)

	// [Option+¥]で[\]が入力できる
	fmt.Printf("Ask=%f\n", ticker.Ask)
	fmt.Printf("Bid=%f\n", ticker.Bid)
	fmt.Printf("High=%f\n", ticker.High)
	fmt.Printf("Last=%f\n", ticker.Last)
	fmt.Printf("Low=%f\n", ticker.Low)
	fmt.Printf("Volume=%f\n", ticker.Volume)

	dtFromUnix := time.Unix(ticker.Timestamp, 0)
	fmt.Printf("Timestamp=%s\n", dtFromUnix.Format("2006/01/02 15:04:05"))

	fmt.Println("End!")
}
