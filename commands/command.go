package commands

import (
	"nyx/logger"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		&WeatherData,
		&PingData,
	}
	
	Handlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate) {
		"weather": WeatherHandler,
		"ping": PingHandler,
	}
)

func Create(s *discordgo.Session) {
	logger.Logger.Info("Loading commands.")
	for _, command := range commands {
		logger.Logger.InfoF("%vLoading command: ", command.Name)
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
		if err != nil {
			logger.Logger.WarningF("[WARNING]: %s", err.Error())
		}
	}
}

func Remove(s *discordgo.Session) {
	logger.Logger.Info("Loading commands.")
	for _, command := range commands {
		logger.Logger.InfoF("%vRemoving command: ", command.Name)
		err := s.ApplicationCommandDelete(s.State.User.ID, "", command.ID)
		if err != nil {
			logger.Logger.WarningF("[WARNING]: %s", err.Error())
		}
	}
}
