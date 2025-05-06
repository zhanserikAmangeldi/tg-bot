package config

import (
	"github.com/joho/godotenv"
	"github.com/zhanserikAmangeldi/tg-bot/internal/domain"
	"os"
)

var (
	alerts     = make(map[int64][]domain.PriceAlert)
	BinanceAPI = "https://api.binance.com/api/v3/ticker/price?symbol="
)

type Config struct {
	TELEGRAM_BOT_TOKEN string
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		TELEGRAM_BOT_TOKEN: getEnv("TELEGRAM_BOT_TOKEN", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
