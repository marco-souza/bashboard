package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func MakeRequest(url string, params url.Values) *http.Request {
	url += "?" + params.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	return req
}
