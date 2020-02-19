package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(fmt.Sprintf("Weather will appear soon")))
}
