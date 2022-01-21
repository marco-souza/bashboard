package cmd

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	UpdateTime time.Time   `JSON:"updateTime"`
	Data       AccountData `json:"data"`
}

type AccountSnapshotResponse struct {
	Code        int               `json:"code"`
	Msg         string            `json:"msg"`
	SnapshotVos []AccountSnapshot `json:"snapshotVos"`
}

func FetchAccountSnapshot() *AccountSnapshotResponse {
	url := getBinanceEndpoint("account-snapshot")
	params := "type=SPOT" // params: https://binance-docs.github.io/apidocs/spot/en/#daily-account-snapshot-user_data
	req := makeSignedRequest(url, params)

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

	fmt.Println("ResposeBody", string(responseBody))

	var accountSnapshot AccountSnapshotResponse
	if err = json.Unmarshal(responseBody, &accountSnapshot); err != nil {
		fmt.Println(string(responseBody))
		panic(err)
	}

	return &accountSnapshot
}

var endpoint = map[string]string{
	"account-snapshot": "/sapi/v1/accountSnapshot",
	"system-status":    "/sapi/v1/system/status",
}

func getBinanceEndpoint(name string) string {
	url := endpoint[name]
	if url == "" {
		panic("No endpoint found")
	}
	return BASE_BINANCE_URL + url
}

func makeSignedRequest(url, params string) *http.Request {
	req, err := http.NewRequest("GET", signUrl(url, params), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-MBX-APIKEY", apiKey)
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
