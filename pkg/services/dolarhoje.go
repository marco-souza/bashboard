package services

import (
	"strconv"
	"strings"

	"github.com/marco-souza/bashbot/pkg/config"
)

func FetchDolarRealExchangeValue() float64 {
	req := MakeRequest(config.USDBRL_EXCHANGE_URL, nil)
	body := Fetch(req)

	exchangeValueString := strings.Replace(string(body), ",", ".", 1)
	exchangeRate, err := strconv.ParseFloat(exchangeValueString, 64)
	if err != nil {
		panic(err)
	}

	return exchangeRate
}
