package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func authReq(req *http.Request) (*http.Request) {
    req.Header.Add("X-API-Key", "UfbGdkn5Fl3wmDTB2gMZTkroCEfiqgDbNGH0kYb0")
    return req
}

func memberVotesReq(memberId string) (*http.Request, error) {
    query := fmt.Sprintf("https://api.propublica.org/congress/v1/members/%s/votes.json", memberId)
    req, err := http.NewRequest("GET", query, nil)
    return authReq(req), err
}

func memberListReq() (*http.Request, error) {
    // TODO un-hardcode congress/chamber
    query := "https://api.propublica.org/congress/v1/116/senate/members.json"
    req, err := http.NewRequest("GET", query, nil)
    return authReq(req), err
}

type person struct {
    Id string
    Title string
}

type result struct {
    Members []person
}

type memberListResp struct {
    Results []result
}

func main() {
    client := http.Client{}

    req, _ := memberListReq()
    resp, _ := client.Do(req)
    body, _ := ioutil.ReadAll(resp.Body)

    var memberIds []string

    members := memberListResp{}
    fmt.Printf("%s", body)
    json.Unmarshal(body, &members)
    for _, r := range members.Results {
        for _, m := range r.Members {
            memberIds = append(memberIds, m.Id)
        }
    }

    /*
    _, _ := memberVotesReq(memberIds[0])
    _, _ := client.Do(req)
    _, _ := ioutil.ReadAll(resp.Body)
    */
    
    fmt.Printf("%s", body)
}
