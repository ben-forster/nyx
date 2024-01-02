package events

import (
	"nyx/commands"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if cmd, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
		cmd(s, i)
	}
}
