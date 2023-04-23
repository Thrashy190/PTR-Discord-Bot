package main

import (
	"github.com/Thrashy190/PTR-Discord-bot/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Thrashy190/PTR-Discord-bot/pkg/handlers"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config.Init()
	token := os.Getenv("TOKEN")

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	discord.AddHandler(handlers.Ping)
	discord.AddHandler(handlers.Health)
	discord.AddHandler(handlers.ShowRoles)

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	log.Println("Bot is running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
