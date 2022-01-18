package cmd

import (
	"net/http"
	"time"
)

const BASE_BINANCE_URL = "https://api.binance.com"

var client = &http.Client{Timeout: time.Second * 10}
