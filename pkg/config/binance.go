package config

import "os"

var BASE_BINANCE_URL = "https://api.binance.com"
var BINANCE_API_KEY = os.Getenv("BINANCE_API_KEY")
var BINANCE_API_SECRET = os.Getenv("BINANCE_API_SECRET")
