package main

import (
  "time"
  "fmt"
  "net/http"
  "os"
  "io"
  "io/ioutil"
)

func main() {
  start := time.Now()
  fmt.Printf("Time now is %v\n", start)

  ch := make(chan string)

  for _, url := range os.Args[1:] {
    go fetch(url, ch)
  }

  for range os.Args[1:] {
    fmt.Println(<-ch)
  }

  fmt.Printf("%d millis elapsed\n", time.Since(start).Milliseconds())
}

func fetch(url string, ch chan <- string) {
  start := time.Now()
  resp, err := http.Get(url)

  if err != nil {
    ch <- fmt.Sprint(err)
    return
  }

  nbytes, err := io.Copy(ioutil.Discard, resp.Body)

  resp.Body.Close()

  if err != nil {
    ch <- fmt.Sprintf("while reading %s, %v", url, err)
    return
  }

  secs := time.Since(start).Milliseconds()

  ch <- fmt.Sprintf("%d %7d %s", secs, nbytes, url)
}
