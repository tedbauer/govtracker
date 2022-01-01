package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
    client := http.Client{}

    req, _ := http.NewRequest("GET", "https://api.propublica.org/congress/v1/house/votes/recent.json", nil)
    req.Header.Add("X-API-Key", "UfbGdkn5Fl3wmDTB2gMZTkroCEfiqgDbNGH0kYb0")

    resp, _ := client.Do(req)
    body, _ := ioutil.ReadAll(resp.Body)

    fmt.Printf("%s", body)
}
