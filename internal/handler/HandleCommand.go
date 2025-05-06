package handler

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhanserikAmangeldi/tg-bot/internal/config"
	"github.com/zhanserikAmangeldi/tg-bot/internal/domain"
	"log"
	"net/http"
	"strings"
)

func HandleCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	command := msg.Command()
	args := msg.CommandArguments()

	switch command {
	case "start":
		reply := "Hello! I am bot for tracking prices of cryptocurrencies!\n\n" +
			"Commands:\n" +
			"/price BTC - for know current price of BTC\n" +
			"/alert BTC 120000 above - make alert when btc will break the 120000\n" +
			"/alert BTC 700000 below - make alert when btc will break the 700000\n" +
			"/myalerts - show your active alerts\n" +
			"/deletealert ID - delete your active alerts by ID\n"
		sendMessage(bot, chatID, reply)
	case "price":
		if args == "" {
			sendMessage(bot, chatID, "Please, specify the cryptocurrency. For example: /price BTC")
			return
		}

		symbol := strings.ToUpper(args)
		price, err := getCryptoPrice(symbol)
		if err != nil {
			sendMessage(bot, chatID, fmt.Sprintf("Get cryptocurrency error: %v", err))
			return
		}

		reply := fmt.Sprintf("Price for %s: %.03f", symbol, price)
		sendMessage(bot, chatID, reply)
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error while sending message: %v", err)
	}
}

func getCryptoPrice(symbol string) (float64, error) {
	url := config.BinanceAPI + symbol + "USDT"

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return 0, fmt.Errorf("status code is not OK: %v", resp.StatusCode)
	}

	var data domain.CryptoPrice
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return data.Price, nil
}
