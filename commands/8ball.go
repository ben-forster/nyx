package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"nyx/config"
	"nyx/logger"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type EightBall struct {
	Lines []string `json:"lines"`
}

var EightBallData = discordgo.ApplicationCommand{
	Name:        "8ball",
	Description: "Ask the magic eightball a question.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "question (Must end in ?)",
			Description: "The question you want to ask.",
			Required:    true,
		},
	},
}

func EightBallHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
        }); err != nil {
		logger.Logger.WarningF("[ERROR]: %s", err.Error())
        return
    }
	
	var question string

	for _, option := range i.ApplicationCommandData().Options {
		if option.Name == "question" {
			question = option.Value.(string)
		}
	}

	if strings.HasSuffix(question, "?") {
		lines, err := os.ReadFile("./config/8ball.json")

		var obj EightBall

		if err != nil {
			fmt.Println("err", err)
			return
		}

		err = json.Unmarshal(lines, &obj)

		if err != nil {
			fmt.Println("err", err)
			return
		}

		rand.New(rand.NewSource(time.Now().UnixNano()))

		randomIndex := rand.Intn(len(obj.Lines))

		randomLine := obj.Lines[randomIndex]

		eightballEmbed := &discordgo.MessageEmbed{}
		eightballEmbed.Title = "\U0001f3b1" + randomLine
		eightballEmbed.Description = "Question: `" + question + "`"
		eightballEmbed.Timestamp = time.Now().Format(time.RFC3339)
		eightballEmbed.Footer = &discordgo.MessageEmbedFooter{
			Text:    i.Member.User.Username,
			IconURL: i.Member.User.AvatarURL(""),
		}
		eightballEmbed.Color = config.Embed

		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{eightballEmbed},
		})
		return
	} else {
		errorEmbed := &discordgo.MessageEmbed{}
		errorEmbed.Title = config.Info + " Invalid Format"
		errorEmbed.Description = "That is not a valid question."
		errorEmbed.Timestamp = time.Now().Format(time.RFC3339)
		errorEmbed.Footer = &discordgo.MessageEmbedFooter{
			Text:    i.Member.User.Username,
			IconURL: i.Member.User.AvatarURL(""),
		}
		errorEmbed.Color = config.Embed

		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{errorEmbed},
		})
		return
	}
}
