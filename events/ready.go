package events

import (
	"nyx/logger"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	err := session.UpdateListeningStatus("/help")
	if err != nil {
		logger.Logger.WarningF("[ERROR]: %s", err.Error())
	}

	logger.Logger.InfoF("[%s:%s]", session.State.User.Username, "is now online.")
}
