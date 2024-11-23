package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		currentTime := time.Now()
		formattedTime := fmt.Sprintf("%s, %d %s %d",
			currentTime.Weekday(),
			currentTime.Day(),
			currentTime.Month(),
			currentTime.Year())
		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(formattedTime))
		
	}
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
