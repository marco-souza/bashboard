package config

import "os"


var TELEGRAM_BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")
var TELEGRAM_BASE_URL = "https://api.telegram.org/bot" + TELEGRAM_BOT_TOKEN
