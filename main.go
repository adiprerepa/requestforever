package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		fmt.Println("no endpoint found in $ENDPOINT, continuing with http://localhost:80...")
		endpoint = "http://127.0.0.1:80"
	}
	method := os.Getenv("METHOD")
	if method == "" {
		fmt.Println("no method found in $METHOD, defaulting to GET")
		method = "GET"
	}
	intervalDurationStr := os.Getenv("INTERVAL")
	if intervalDurationStr == "" {
		fmt.Println("no duration found in $INTERVAL (ms), defaulting to a request interval duration of 2s")
		intervalDurationStr = "2000"
	}
	intervalDuration, err := strconv.Atoi(intervalDurationStr)
	if err != nil {
		fmt.Printf("couldn't parse duration in $INTERVAL: %v. Defaulting to 2s.\n", err)
		intervalDuration = 2000
	}
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		fmt.Printf("unable to create request: %v\n", err)
		return
	}

	client := &http.Client{}
	for true {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("unable to make request: %v\n", err)
		} else {
			fmt.Printf("request succeeded with code %v\n", resp.StatusCode)
		}
		time.Sleep(time.Duration(intervalDuration) * time.Millisecond)
	}
}
