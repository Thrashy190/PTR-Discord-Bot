package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Thrashy190/PTR-Discord-bot/pkg/handlers"
	"github.com/bwmarrin/discordgo"
)


func hello(s*discordgo.Session, m*discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "Hello"{
		s.ChannelMessageSend(m.ChannelID, "Hola, "+m.Author.Mention()+"!, Como estas?")
	}
}

func ping(s*discordgo.Session, m*discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
}


func main() {

	//token := os.Getenv("TOKEN")

	discord, err := discordgo.New("Bot " + "MTA5OTQxOTc4MTI4ODk3MjM5MQ.GOMqH3.VOybSu3dK8E2TJhmZr6IolhK8zGMp_zUQLTS2U")
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

	fmt.Println("Bot is running!")


	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}