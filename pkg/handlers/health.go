package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)


func Health(s*discordgo.Session, m*discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	//give me de bot status with the command Health with a 200 status code

	if m.Content == "health" {
		s.ChannelMessageSend(m.ChannelID, s.State.User.Username+" is running!")
		log.Println("Health command executed")
	}
}

func Ping(s*discordgo.Session, m*discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
		log.Println("Ping command executed")
	}
}