package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"nyx/commands"
	"nyx/config"
	"nyx/logger"
	"nyx/db"
	"nyx/events"

	"github.com/bwmarrin/discordgo"
)

var (
	session 	*discordgo.Session
)

func main() {
	var flagMigrateCommands bool

	flag.BoolVar(&flagMigrateCommands, "commands", false, "Update commands.")

	flag.Parse()

	db.Connect()
	config.ReadEnv()
	config.ReadConfig()

	bot, err := discordgo.New(fmt.Sprintf("Bot %v", config.Token))
    if err != nil {
      	logger.Logger.FatalF("[ERROR]: %v", err.Error())
        return
    }
	defer session.Close()

	bot.AddHandler(events.Ready)
	bot.AddHandler(events.InteractionCreate)

	bot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildPresences | discordgo.IntentsGuildMembers)

	err = bot.Open()
	if err != nil {
		logger.Logger.FatalF("[ERROR]: %v", err.Error())
        return
    }

	if flagMigrateCommands {
		commands.Migrate(bot)
	}

	shutdown()
}

func shutdown() {	
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logger.Logger.InfoF(session.State.User.Username + "is shutting down.")
	session.Close()
}
