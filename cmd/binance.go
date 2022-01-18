package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WalletInfoResponse struct {
	Name string `json:"name"`
}

func FetchWalletInfo() *WalletInfoResponse {
	url := getBinanceEndpoint("system-status")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Fetching ", url)

	// Setup query params
	query := req.URL.Query()
	// TODO: add wallet info query param
	// TODO: sign api req
	query.Set("signature", "test")
	req.URL.RawQuery = query.Encode()

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Parse responseBody
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var walletInfo WalletInfoResponse
	if err = json.Unmarshal(responseBody, &walletInfo); err != nil {
		fmt.Println(string(responseBody))
		panic(err)
	}

	return &walletInfo
}

type SystemStatusResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func FetchSystemStatus() *SystemStatusResponse {
	url := getBinanceEndpoint("system-status")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Parse responseBody
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var systemStatus SystemStatusResponse
	if err = json.Unmarshal(responseBody, &systemStatus); err != nil {
		fmt.Println(string(responseBody))
		panic(err)
	}

	return &systemStatus
}

var endpoint = map[string]string{
	"wallet":        "/wallet",
	"system-status": "/sapi/v1/system/status",
}

func getBinanceEndpoint(name string) string {
	url := endpoint[name]
	if url == "" {
		panic("No endpoint found")
	}
	return BASE_BINANCE_URL + url
}
