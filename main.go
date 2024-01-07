package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"nyx/commands"
	"nyx/config"
	"nyx/database"
	"nyx/events"
	"nyx/logger"

	"github.com/bwmarrin/discordgo"
)

var (
	session 	*discordgo.Session
)

func main() {
	var flagMigrateCommands bool

	flag.BoolVar(&flagMigrateCommands, "commands", true, "Update commands.")

	flag.Parse()

	database.Connect()
	config.ReadEnv()
	config.ReadConfig()

	s, err := discordgo.New(fmt.Sprintf("Bot %v", config.Token))
   	if err != nil {
      		logger.Logger.FatalF("[ERROR]: %v", err.Error())
        	return
    	}
	defer s.Close()

	session = s

	s.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildPresences | discordgo.IntentsGuildMembers

	s.AddHandler(events.Ready)
	s.AddHandler(events.InteractionCreate)

	err = s.Open()
	if err != nil {
		logger.Logger.FatalF("[ERROR]: %v", err.Error())
        return
    }

	if flagMigrateCommands {
		commands.Create(s)
	}

	shutdown()
}

func shutdown() {	
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logger.Logger.InfoF(session.State.User.Username + " is shutting down.")

	commands.Remove(session)
	session.Close()
}
