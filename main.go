package main

import (
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	if len(url) == 0 {
		os.Exit(1)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
		return
	}
	statusCode := resp.StatusCode
	if statusCode < 200 || statusCode >= 400 {
		os.Exit(1)
		return
	}
	os.Exit(0)
}
