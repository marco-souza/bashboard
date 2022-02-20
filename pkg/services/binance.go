package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/marco-souza/bashbot/pkg/config"
	"github.com/marco-souza/bashbot/pkg/entities"
)

func FetchSystemStatus() *entities.SystemStatusResponse {
	url := getBinanceEndpoint("system-status")
	req := MakeRequest(url, "")

	// Send the request
	responseBody := Fetch(req)

	var systemStatus entities.SystemStatusResponse
	if err := json.Unmarshal(responseBody, &systemStatus); err != nil {
		panic(err)
	}

	return &systemStatus
}

func FetchAccountSnapshot() *entities.AccountSnapshotResponse {
	url := getBinanceEndpoint("account-snapshot")
	params := fmt.Sprintf("type=SPOT&endTime=%d", time.Now().Unix()*1000) // params: https://binance-docs.github.io/apidocs/spot/en/#daily-account-snapshot-user_data
	req := makeSignedRequest(url, params)

	responseBody := Fetch(req)

	var accountSnapshot entities.AccountSnapshotResponse
	if err := json.Unmarshal(responseBody, &accountSnapshot); err != nil {
		panic(err)
	}

	return &accountSnapshot
}

func FetchTicker(currencyPair string) *entities.Ticker {
	// API ref: https://binance-docs.github.io/apidocs/spot/en/#symbol-price-ticker
	url := getBinanceEndpoint("ticker")
	params := "symbol=" + currencyPair
	req := MakeRequest(url, params)

	responseBody := Fetch(req)

	var ticker entities.Ticker
	if err := json.Unmarshal(responseBody, &ticker); err != nil {
		panic(err)
	}

	log.Println("Tiker", ticker)

	return &ticker
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
	return config.BASE_BINANCE_URL + url
}

func makeSignedRequest(url, params string) *http.Request {
	// TODO handle query params properly
	req := MakeRequest(signUrl(url, params), "")
	req.Header.Set("X-MBX-APIKEY", config.BINANCE_API_KEY)
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
	hash := hmac.New(sha256.New, []byte(config.BINANCE_API_SECRET))
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
