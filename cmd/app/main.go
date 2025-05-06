package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhanserikAmangeldi/tg-bot/internal/config"
	"github.com/zhanserikAmangeldi/tg-bot/internal/domain"
	"log"
)

var (
	alerts     = make(map[int64][]domain.PriceAlert)
	binanceAPI = "https://api.binance.com/api/v3/ticker/price?symbol="
)

func main() {
	var cfg *config.Config

	cfg = config.NewConfig()
	if cfg.TELEGRAM_BOT_TOKEN == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable not founded. Please set in env var.")
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TELEGRAM_BOT_TOKEN)
	if err != nil {
		log.Fatalf("Failed to connect to telegram bot: %s", err)
		return
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}
