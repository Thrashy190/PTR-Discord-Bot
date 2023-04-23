package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Thrashy190/PTR-Discord-bot/pkg/config"
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

	discord.AddHandler(handlers.RolesOfUser)
	discord.AddHandler(handlers.Ping)
	discord.AddHandler(handlers.Health)
	discord.AddHandler(handlers.ShowAndSelectRoles)

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	fmt.Println("Bot is running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
