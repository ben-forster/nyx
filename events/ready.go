package events

import (
	"nyx/logger"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	err := s.UpdateListeningStatus("/help")
	if err != nil {
		logger.Logger.WarningF("[ERROR]: %s", err.Error())
	}

	logger.Logger.InfoF("[%s:%s]", s.State.User.Username, "is now online.")
}
