package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"time"
)

func main() {
	var host string
	var limit int
	var interval int
	var stationID int
	var token string

	flag.StringVar(&host, "host", "http://localhost:8000", "API server host")
	flag.IntVar(&limit, "limit", 1000, "Number of data entries to simulate")
	flag.IntVar(&interval, "interval", 10, "Interval in seconds between data entries")
	flag.IntVar(&stationID, "station-id", 1, "Station ID")
	flag.StringVar(&token, "token", "", "Authentication token")
	flag.Parse()
	path := host + "/api/v1/sensors"

	var entry struct {
		StationID int `json:"station_id"`
		TempData  struct {
			Temperature float64 `json:"data"`
		} `json:"temp_data"`
	}
	entry.StationID = stationID
	entry.TempData.Temperature = 33

	for i := 0; i < limit; i++ {
		r := rand.Float64()
		if rand.IntN(2) == 0 {
			r *= -1
		}
		entry.TempData.Temperature += r
		payload, err := json.Marshal(entry)
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest("POST", path, bytes.NewBuffer(payload))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var js map[string]any
		json.Unmarshal(body, &js)
		fmt.Println(js)

		time.Sleep(time.Second * time.Duration(interval))
	}
}
