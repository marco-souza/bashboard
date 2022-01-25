package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const BASE_BINANCE_URL = "https://api.binance.com"

var client = &http.Client{Timeout: time.Second * 10}
var staticFilePath = "./static/"
var IP = ""
var PORT = "8001"

func StartPageServer(port string) {
	fs := http.FileServer(http.Dir(staticFilePath))
	http.Handle("/", fs)

	address := fmt.Sprintf("%s:%s", IP, PORT)

	log.Println("Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
