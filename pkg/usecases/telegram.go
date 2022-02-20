package usecases

import (
	"fmt"
	"strconv"

	"github.com/marco-souza/bashbot/pkg/services"
)

func SendWalletReport() {
	resp := services.FetchAccountSnapshot("SPOT")
	snap := resp.SnapshotVos[len(resp.SnapshotVos)-1]

	totalBtcAmount, err := strconv.ParseFloat(snap.Data.TotalBtcAsset, 32)
	if err != nil {
		panic(err)
	}

	respTiker := services.FetchTicker("BTCUSDT")
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

	services.SendChatMessage(services.CHAT_ID, msg)
}
