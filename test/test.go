package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	url := "http://127.0.0.1:8080/ping"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Thunder Client (https://www.thunderclient.com)")

	times := 300

	start := time.Now()

	for i := range times {
		res, err := http.DefaultClient.Do(req)
		if err == nil {
			fmt.Println(i)
		}
		res.Body.Close()
	}
	fmt.Println(time.Since(start))
}
