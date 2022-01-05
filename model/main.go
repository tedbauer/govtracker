package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func authReq(req *http.Request) *http.Request {
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

type vote struct {
	Position    string
	Description string
}

type memberVotesRes struct {
	Votes []vote
}

type memberVotesResp struct {
	Results []memberVotesRes
}

type person struct {
	Id         string
	Title      string
	First_name string
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
	//fmt.Printf("%s", body)
	json.Unmarshal(body, &members)
	for _, r := range members.Results {
		for _, m := range r.Members {
			memberIds = append(memberIds, m.Id)
		}
	}

	req2, _ := memberVotesReq(memberIds[0])
	resp2, _ := client.Do(req2)
	body2, _ := ioutil.ReadAll(resp2.Body)

	votesResp := memberVotesResp{}
	json.Unmarshal(body2, &votesResp)

	fmt.Printf("Votes for %s, %s:\n", members.Results[0].Members[0].First_name, members.Results[0].Members[0].Title)
	for _, r := range votesResp.Results {
		for _, v := range r.Votes {
			fmt.Printf("Bill: %s\n", v.Description)
			fmt.Printf("Vote: %s\n", v.Position)
		}
	}

	//fmt.Printf("%s", body)
}
