package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	version := "v1"

	w.Write([]byte(fmt.Sprintf("Time now is: %s\n", t.Format(time.RFC850))))
	w.Write([]byte(fmt.Sprintf("Hostname: %s\n", hostname)))
	w.Write([]byte(fmt.Sprintf("Version: %s\n", version)))
}
