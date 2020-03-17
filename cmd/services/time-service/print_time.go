package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting server at 0.0.0.0:9090")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	v := "v1"

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("Time now is: %s\n", t.Format(time.RFC850))))
	w.Write([]byte(fmt.Sprintf("Hostname: %s\n", hostname)))
	w.Write([]byte(fmt.Sprintf("Version: %s\n", v)))
}
