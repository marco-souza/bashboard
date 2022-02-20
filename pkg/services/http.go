package services

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(req *http.Request) []byte {
	log.Println("RequestURL: ", req.URL)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	// Parse responseBody
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println("ResposeBody", string(responseBody))

	return responseBody
}

func MakeRequest(url, params string) *http.Request {
	if params != "" {
		url += "?" + params
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	return req
}
