package cmd

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var apiKey = os.Getenv("BINANCE_API_KEY")
var apiSecret = os.Getenv("BINANCE_API_SECRET")

type SystemStatusResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func FetchSystemStatus() *SystemStatusResponse {
	url := getBinanceEndpoint("system-status")
	req := makeRequest(url, "")

	// Send the request
	responseBody := fetch(req)

	var systemStatus SystemStatusResponse
	if err := json.Unmarshal(responseBody, &systemStatus); err != nil {
		panic(err)
	}

	return &systemStatus
}

type AccountAssets struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type AccountData struct {
	TotalBtcAsset string          `json:"totalAssetOfBtc"`
	Balances      []AccountAssets `json:"balances"`
}

type AccountSnapshot struct {
	Type       string      `json:"type"`
	UpdateTime int         `json:"updateTime"`
	Data       AccountData `json:"data"`
}

type AccountSnapshotResponse struct {
	Code        int               `json:"code"`
	Msg         string            `json:"msg"`
	SnapshotVos []AccountSnapshot `json:"snapshotVos"`
}

func FetchAccountSnapshot() *AccountSnapshotResponse {
	url := getBinanceEndpoint("account-snapshot")
	params := fmt.Sprintf("type=SPOT&endTime=%d", time.Now().Unix()*1000) // params: https://binance-docs.github.io/apidocs/spot/en/#daily-account-snapshot-user_data
	req := makeSignedRequest(url, params)

	responseBody := fetch(req)

	var accountSnapshot AccountSnapshotResponse
	if err := json.Unmarshal(responseBody, &accountSnapshot); err != nil {
		panic(err)
	}

	return &accountSnapshot
}

type Ticker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func FetchTicker(currencyPair string) *Ticker {
	// API ref: https://binance-docs.github.io/apidocs/spot/en/#symbol-price-ticker
	url := getBinanceEndpoint("ticker")
	params := "symbol=" + currencyPair
	req := makeRequest(url, params)

	responseBody := fetch(req)

	var ticker Ticker
	if err := json.Unmarshal(responseBody, &ticker); err != nil {
		panic(err)
	}

	log.Println("Tiker", ticker)

	return &ticker
}

func fetch(req *http.Request) []byte {
	log.Println("RequestURL: ", req.URL)

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

	log.Println("ResposeBody", string(responseBody))

	return responseBody
}

var endpoint = map[string]string{
	"account-snapshot": "/sapi/v1/accountSnapshot",
	"system-status":    "/sapi/v1/system/status",
	"ticker":           "/api/v3/ticker/price",
}

func getBinanceEndpoint(name string) string {
	url := endpoint[name]
	if url == "" {
		panic(fmt.Sprintf("No '%s' endpoint found", name))
	}
	return BASE_BINANCE_URL + url
}

func makeSignedRequest(url, params string) *http.Request {
	// TODO handle query params properly
	req := makeRequest(signUrl(url, params), "")
	req.Header.Set("X-MBX-APIKEY", apiKey)
	return req
}

func makeRequest(url, params string) *http.Request {
	if params != "" {
		url += "?" + params
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	return req
}

func signUrl(url string, params string) string {
	if params != "" {
		params += "&"
	}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	params += "timestamp=" + timestamp

	signature := sign(params)
	signedUrl := fmt.Sprintf("%s?%s&signature=%s", url, params, signature)

	return signedUrl
}

func sign(text string) string {
	hash := hmac.New(sha256.New, []byte(apiSecret))
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
