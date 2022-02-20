package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/marco-souza/bashbot/pkg/config"
	"github.com/marco-souza/bashbot/pkg/entities"
)

func FetchAccountSnapshot(walletType string) *entities.AccountSnapshotResponse {
	accountSnapURL := getBinanceEndpoint("account-snapshot")

	params := url.Values{}
	params.Add("type", walletType)
	params.Add("endTime", fmt.Sprint(time.Now().Unix() * 1000))

	req := makeSignedRequest(accountSnapURL, params)
	responseBody := Fetch(req)

	var accountSnapshot entities.AccountSnapshotResponse
	if err := json.Unmarshal(responseBody, &accountSnapshot); err != nil {
		panic(err)
	}

	return &accountSnapshot
}

func FetchTicker(currencyPair string) *entities.Ticker {
	// API ref: https://binance-docs.github.io/apidocs/spot/en/#symbol-price-ticker
	tickerURL := getBinanceEndpoint("ticker")

	params := url.Values{}
	params.Add("symbol", currencyPair)

	req := MakeRequest(tickerURL, params)
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

func makeSignedRequest(url string, params url.Values) *http.Request {
	signedParams := signParams(params)
	req := MakeRequest(url, signedParams)

	req.Header.Set("X-MBX-APIKEY", config.BINANCE_API_KEY)

	return req
}

func signParams(params url.Values) url.Values {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	params.Add("timestamp", timestamp)

	signature := sign(params.Encode())
	params.Add("signature", signature)

	return params
}

func sign(text string) string {
	hash := hmac.New(sha256.New, []byte(config.BINANCE_API_SECRET))
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
