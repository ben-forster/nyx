package config

import (
	"encoding/json"
	"os"

	"nyx/logger"

	"github.com/joho/godotenv"
)

var (
	Token					string
	EmbedColor 				int
	DeleteAfter 				int
	Checkmark      				string
	Cross 					string
	Info 					string

	config *configStruct
)

type configStruct struct {
	EmbedColor				int
	DeleteAfter 				int
	Checkmark      				string
	Cross 					string
	Info 					string
}

func ReadEnv() error {
	logger.Logger.InfoF("Loading environmental variables.")
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("[ERROR]: " + err.Error())
	}

	Token = os.Getenv("BOT_TOKEN")
	logger.Logger.Infof("Environmental variables loaded.")

	return nil
}

func ReadConfig() error {	
	logger.Logger.InfoF("Loading configuration file.")
	file, err := os.ReadFile("config/config.json")
	if err != nil {
		logger.Logger.Fatal("[ERROR]: " + err.Error())
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		logger.Logger.Fatal("[ERROR]: " + err.Error())
	}
	logger.Logger.InfoF("Configuration file loaded.")

	EmbedColor = config.EmbedColor
	DeleteAfter = config.DeleteAfter
	Checkmark = config.Checkmark
	Cross = config.Cross
	Info = config.Info

	return nil
}
