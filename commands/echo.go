package commands

import (
	"time"

	"nyx/config"
	"nyx/logger"

	"github.com/bwmarrin/discordgo"
)

var EchoData = discordgo.ApplicationCommand{
	Name:        "echo",
	Description: "Repeats what you say.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "message",
			Description: "The message you want the bot to repeat.",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionChannel,
			Name:        "channel",
			Description: "The channel you want the message to send TO.",
			ChannelTypes: []discordgo.ChannelType{0},
			Required:    false,
		},
	},
}

func EchoHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Flags: discordgo.MessageFlagsEphemeral,
        },
    }); err != nil {
		logger.Logger.WarningF("[ERROR]: %s", err.Error())
        return
    }

	var messageText string
	var channelID string

	for _, option := range i.ApplicationCommandData().Options {
		if option.Name == "message" {
			messageText = option.Value.(string)
		}
		if option.Name == "channel" {
			channelID = option.Value.(string)
		}
	}

	if channelID == "" {
		_, err := s.ChannelMessageSend(i.ChannelID, messageText)
		if err != nil {
			errorEmbed := &discordgo.MessageEmbed{}
			errorEmbed.Title = config.Cross + " Error"
			errorEmbed.Description = err.Error()
			errorEmbed.Footer = &discordgo.MessageEmbedFooter{
				Text:    i.Member.User.Username,
				IconURL: i.Member.User.AvatarURL(""),
			}
			errorEmbed.Timestamp = time.Now().Format(time.RFC3339)
			errorEmbed.Color = config.EmbedColor
		
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{errorEmbed},
			})
			logger.Logger.WarningF("[ERROR]: %s", err.Error())
			return
		}

		channel := "<#" + i.ChannelID + ">"
		echoEmbed := &discordgo.MessageEmbed{}
		echoEmbed.Title = config.Checkmark + " Message sent"
		echoEmbed.Description = "Message `" + messageText + "` was sent to " + channel + "."
		echoEmbed.Footer = &discordgo.MessageEmbedFooter{
			Text:    i.Member.User.Username,
			IconURL: i.Member.User.AvatarURL(""),
		}
		echoEmbed.Timestamp = time.Now().Format(time.RFC3339)
		echoEmbed.Color = config.EmbedColor

		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{echoEmbed},
		})
	} else {
		_, err := s.ChannelMessageSend(channelID, messageText)
		if err != nil {

			errorEmbed := &discordgo.MessageEmbed{}
			errorEmbed.Title = config.Cross + " Error"
			errorEmbed.Description = err.Error()
			errorEmbed.Footer = &discordgo.MessageEmbedFooter{
				Text:    i.Member.User.Username,
				IconURL: i.Member.User.AvatarURL(""),
			}
			errorEmbed.Timestamp = time.Now().Format(time.RFC3339)
			errorEmbed.Color = config.EmbedColor
		
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{errorEmbed},
			})
			logger.Logger.WarningF("[ERROR]: %s", err.Error())
			return
		}
	
		channel := "<#" + channelID + ">"
		echoEmbed := &discordgo.MessageEmbed{}
		echoEmbed.Title = config.Checkmark + " Message sent"
		echoEmbed.Description = "Message `" + messageText + "` was sent to " + channel + "."
		echoEmbed.Footer = &discordgo.MessageEmbedFooter{
			Text:    i.Member.User.Username,
			IconURL: i.Member.User.AvatarURL(""),
		}
		echoEmbed.Timestamp = time.Now().Format(time.RFC3339)
		echoEmbed.Color = config.EmbedColor

		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{echoEmbed},
		})
		return
	}
}
