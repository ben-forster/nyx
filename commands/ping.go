package commands

import (
	"time"

	"nyx/config"
	"nyx/logger"

	"github.com/bwmarrin/discordgo"
)

var PingData = discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Replies with the bot latency.",
}

func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	pingingEmbed := &discordgo.MessageEmbed{}
	pingingEmbed.Description = "**Pinging...**"
	pingingEmbed.Timestamp = time.Now().Format(time.RFC3339)
	pingingEmbed.Color = config.Embed

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{pingingEmbed},
		},
	}); err != nil {
		errorEmbed := &discordgo.MessageEmbed{}
		errorEmbed.Title = config.Cross + " Error"
		errorEmbed.Description = err.Error()
		errorEmbed.Footer = &discordgo.MessageEmbedFooter{
			Text:    i.Member.User.Username,
			IconURL: i.Member.User.AvatarURL(""),
		}
		errorEmbed.Timestamp = time.Now().Format(time.RFC3339)
		errorEmbed.Color = config.Embed
	
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{errorEmbed},
		})
		logger.Logger.WarningF("[ERROR]: %s", err.Error())
		return
	}

	id := i.Interaction.ID
	timestamp, err := discordgo.SnowflakeTimestamp(id)
	if err != nil {
		errorEmbed := &discordgo.MessageEmbed{}
		errorEmbed.Title = config.Cross + " Error"
		errorEmbed.Description = err.Error()
		errorEmbed.Footer = &discordgo.MessageEmbedFooter{
			Text:    i.Member.User.Username,
			IconURL: i.Member.User.AvatarURL(""),
		}
		errorEmbed.Timestamp = time.Now().Format(time.RFC3339)
		errorEmbed.Color = config.Embed
	
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{errorEmbed},
		})
		logger.Logger.WarningF("[ERROR]: %s", err.Error())
		return
	}

	duration := time.Since(timestamp)
	formattedDuration := duration.String()

	pingEmbed := &discordgo.MessageEmbed{}
	pingEmbed.Title = "\U0001f3d3 Pong!"
	pingEmbed.Description = "Heartbeat: `" + formattedDuration + "` "
	pingEmbed.Footer = &discordgo.MessageEmbedFooter{
		Text:    i.Member.User.Username,
		IconURL: i.Member.User.AvatarURL(""),
	}
	pingEmbed.Timestamp = time.Now().Format(time.RFC3339)
	pingEmbed.Color = config.Embed

	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{pingEmbed},
	})
}
