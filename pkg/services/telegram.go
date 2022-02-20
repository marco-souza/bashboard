package services

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/marco-souza/bashbot/pkg/config"
)

// TODO: make it dynamic
const CHAT_ID = 161456907

func SendChatMessage(chatID int, message string) {
	fmt.Println("Sending message: ", message)

	telegramURL := fmt.Sprintf("%s/sendMessage", config.TELEGRAM_BASE_URL)

	params := url.Values{}
	params.Add("chat_id", strconv.Itoa(chatID))
	params.Add("parse_mode", "Markdown")
	params.Add("text", message)

	req := MakeRequest(telegramURL, params)
	body := Fetch(req)

	fmt.Println(string(body))
}
