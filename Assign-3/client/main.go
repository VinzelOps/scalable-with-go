package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type StatusResponse struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}

func main() {
	fetchHTTP()
	
	go func() {
		// timer 15 detik
		timer := time.NewTicker(15 * time.Second)
		defer timer.Stop()

		for {
			<-timer.C
			//15 detik
			go fetchHTTP()
		}
	}()

	select {}
}

func fetchHTTP() {

	url := "http://localhost:3000/status"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error parsing body:", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", response.StatusCode)
		return
	}

	var statusRes StatusResponse


	err = json.Unmarshal(body, &statusRes)
	if err != nil {
		log.Fatalln("Error unmarshalling JSON:", err)
		return
	}

	var waterStatus string
	switch {
	case statusRes.Status.Water < 5:
		waterStatus = "Aman"
	case statusRes.Status.Water >= 6 && statusRes.Status.Water <= 8:
		waterStatus = "Siaga"
	default:
		waterStatus = "Bahaya"
	}
	var windStatus string
	switch {
	case statusRes.Status.Wind < 6:
		windStatus = "Aman"
	case statusRes.Status.Wind >= 7 && statusRes.Status.Wind <= 15:
		windStatus = "Siaga"
	default:
		windStatus = "Bahaya"
	}

	// tampilkan waktu
	log.Printf("%s - Response from %s: water=%d m (%s), wind=%d m/s (%s)\n",
		time.Now().Format("2006-01-02 15:04:05"), url, statusRes.Status.Water, waterStatus, statusRes.Status.Wind, windStatus)
}
