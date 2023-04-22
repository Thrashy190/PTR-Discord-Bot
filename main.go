package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)


func roles(s*discordgo.Session, m*discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	//genera una carta con un header y un footer de prueba
	if m.Content == "carta" {
		embed := &discordgo.MessageEmbed{
			Title: "Carta de prueba",
			Description: "Esto es una carta de prueba",
			Color: 0x00ff00,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Footer de prueba",
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}

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

	token := os.Getenv("TOKEN")

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	discord.AddHandler(ping)
	discord.AddHandler(hello)
	discord.AddHandler(roles)

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