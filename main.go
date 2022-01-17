package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type GithubResponse struct {
	Name          string `json:"name"`
	AvatarUrl     string `json:"avatar_url"`
	GithubProfile string `json:"url"`
	Bio           string `json:"bio"`
}

var client = &http.Client{
	Timeout: time.Second * 10,
}

func fetch(url, verb, bodyText string) {
	requestBody := bytes.NewBufferString(bodyText)
	req, err := http.NewRequest(verb, url, requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	query := req.URL.Query()
	query.Set("signature", "test")
	req.URL.RawQuery = query.Encode()

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var githubResp GithubResponse
	if err = json.Unmarshal(body, &githubResp); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(githubResp.Name)
	fmt.Println(githubResp.Bio)
	// fmt.Println(githubResp)
}

func main() {
	fetch("https://api.github.com/users/marco-souza?signature=shity-shit", "GET", "")
}
