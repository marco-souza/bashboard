package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

// TODO: make it dynamic
const CHAT_ID = 161456907

var TELEGRAM_BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")
var TELEGRAM_BASE_URL = "https://api.telegram.org/bot" + TELEGRAM_BOT_TOKEN

func SendChatMessage(chatID int, message string) {
	fmt.Println("Sending message: ", message)

	telegramURL := fmt.Sprintf("%s/sendMessage", TELEGRAM_BASE_URL)

	params := url.Values{}
	params.Add("chat_id", strconv.Itoa(chatID))
	params.Add("parse_mode", "Markdown")
	params.Add("text", message)

	req := makeRequest(telegramURL, params.Encode())
	body := fetch(req)

	fmt.Println(string(body))
}

func SendWalletReport() {
	resp := FetchAccountSnapshot()
	snap := resp.SnapshotVos[len(resp.SnapshotVos)-1]

	totalBtcAmount, err := strconv.ParseFloat(snap.Data.TotalBtcAsset, 32)
	if err != nil {
		panic(err)
	}

	respTiker := FetchTicker("BTCUSDT")
	tikerPrice, err := strconv.ParseFloat(respTiker.Price, 32)
	if err != nil {
		panic(err)
	}

	// Get total wallet amount in USD
	totalUSDAmount := totalBtcAmount * tikerPrice

	// TODO: Fetch exchange ticker USD-BRL
	tickerBRLPrice := 5.0
	totalBRLAmount := totalUSDAmount * tickerBRLPrice

	msg := fmt.Sprintf("*Binance Wallet Report*\n\n - USD: $ %.2f\n - BRL: R$ %.2f", totalUSDAmount, totalBRLAmount)

	SendChatMessage(CHAT_ID, msg)
}
