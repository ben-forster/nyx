package commands

import (
	"time"

	"nyx/config"
	"nyx/logger"

	"github.com/bwmarrin/discordgo"
)

var AvatarData = discordgo.ApplicationCommand{
	Name:        "avatar",
	Description: "Sends a user's avatar.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "user",
			Description: "The user you want the avatar of.",
			Required:    true,
		},
	},
}

func AvatarHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var user string

	for _, option := range i.ApplicationCommandData().Options {
		if option.Name == "user" {
			user = option.Value.(string)
		}
	}

	member, err := s.GuildMember(i.GuildID, user)
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

	size := "512"

	avatar := member.User.AvatarURL("") + "?size=" + size
	
		avatarEmbed := discordgo.MessageEmbed{
			Title:      "\U0001f5bc " + member.User.Username + "'s Avatar",
			Image:      &discordgo.MessageEmbedImage{URL: avatar},
			Timestamp:  time.Now().Format(time.RFC3339),
			Footer: 	&discordgo.MessageEmbedFooter{
				Text:    i.Member.User.Username,
				IconURL: i.Member.User.AvatarURL(""),
			},
			Color:      config.EmbedColor,
		}	
		if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{&avatarEmbed},
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
			errorEmbed.Color = config.EmbedColor
		
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{errorEmbed},
			})
			logger.Logger.WarningF("[ERROR]: %s", err.Error())
			return
		}
	}
