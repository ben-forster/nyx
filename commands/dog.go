package commands

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/text/cases"

	"nyx/config"
	"nyx/logger"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/language"
)

type DogHTTPData struct {
	Image string `json:"image"`
	Fact string `json:"fact"`
}

var DogData = discordgo.ApplicationCommand{
	Name:        "dog",
	Description: "Sends an image of a dog.",
}

func DogHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
        }); err != nil {
		logger.Logger.WarningF("[ERROR]: %s", err.Error())
        return
    }

	client := http.Client{Timeout: 10 * time.Second}

	response, err := client.Get("https://some-random-api.com/animal/dog")

	if err != nil {
		errorEmbed := &discordgo.MessageEmbed{}
		errorEmbed.Title = config.Info + " Invalid Format"
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
	defer response.Body.Close()

	var data DogHTTPData
	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		errorEmbed := &discordgo.MessageEmbed{}
		errorEmbed.Title = config.Info + " Invalid Format"
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

	dogEmbed := &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: data.Image,
		},
		Description: "Did you know `" + cases.Lower(language.Und).String(data.Fact) + "`",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       config.EmbedColor,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Powered by somerandomapi.com",
		},
	}

	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{dogEmbed},
	})
}
