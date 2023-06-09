package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Health(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	

	if m.Content == "health" {
		s.ChannelMessageSend(m.ChannelID, s.State.User.Username+" is running! ")
		log.Println("Health command executed")
	}
}

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "png")
		log.Println("Ping command executed")
	}
}

func RolesOfUser(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	discordgo. 

	if m.Content == "rolesList" {
		s.ChannelMessageSend(m.ChannelID, m.Author.Username+" has the following roles: "+m.GuildID)
		log.Println("RolesList command executed")
	}
}
