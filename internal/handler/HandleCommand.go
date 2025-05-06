package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func HandleCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	command := msg.Command()

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
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error while sending message: %v", err)
	}
}
