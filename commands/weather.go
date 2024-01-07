package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/text/cases"

	"nyx/config"
	"nyx/logger"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"golang.org/x/text/language"
)

const OWMURL string = "https://api.openweathermap.org/data/2.5/weather?"

type WeatherHTTPData struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
}

var WeatherData = discordgo.ApplicationCommand{
	Name:        "weather",
	Description: "Sends the current weather.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "city",
			Description: "The city you want the weather for.",
			Required:    true,
		},
	},
}

func WeatherHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
        }); err != nil {
		logger.Logger.WarningF("[ERROR]: %s", err.Error())
        return
    }
	
	var city string

	for _, option := range i.ApplicationCommandData().Options {
		if option.Name == "city" {
			city = option.Value.(string)
		}
	}

	godotenv.Load()

	OW_TOKEN := os.Getenv("OW_TOKEN")
	units := "metric"

	weatherURL := fmt.Sprintf("%sappid=%s&q=%s&units=%s", OWMURL, OW_TOKEN, city, units)

	client := http.Client{Timeout: 10 * time.Second}

	response, err := client.Get(weatherURL)

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

	var data WeatherHTTPData
	err = json.NewDecoder(response.Body).Decode(&data)
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

	cityName := data.Name
	firstconditions := data.Weather[0].Main
	secconditions := cases.Title(language.Und).String(data.Weather[0].Description)
	temperaturec := strconv.FormatFloat(data.Main.Temp, 'f', 2, 64)
	temperaturef := strconv.FormatFloat((data.Main.Temp * 9/5) + 32, 'f', 2, 64)
	humidity := strconv.Itoa(data.Main.Humidity)
	wind := strconv.FormatFloat(data.Wind.Speed, 'f', 2, 64)
	
	weatherEmbed := &discordgo.MessageEmbed{
		Title: "\U0001f324 Weather for " + cityName,
		Description: "The temperature for " + cityName + " is currently `" + firstconditions + "` with `" + secconditions + "`.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Conditions",
				Value:  firstconditions,
				Inline: true,
			},
			{
				Name:   "Temperature (°C)",
				Value:  temperaturec + "°C",
				Inline: true,
			},
			{
				Name:   "Temperature (°F)",
				Value:  temperaturef + "°F",
				Inline: true,
			},
			{
				Name:   "Humidity",
				Value:  humidity + "%",
				Inline: true,
			},
			{
				Name:   "Wind",
				Value:  wind + " mph",
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: i.Member.User.Username + " | Powered by openweathermap.com",
			IconURL: i.Member.User.AvatarURL(""),
		},
		Color: config.EmbedColor,
	}	
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{weatherEmbed},
	})
	}
