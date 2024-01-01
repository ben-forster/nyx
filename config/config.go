package config

import (
	"os"

	"nyx/logger"

	"github.com/joho/godotenv"
)

var (
	Token			string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("[ERROR]: " + err.Error())
	}

	Token = os.Getenv("BOT_TOKEN")
}
