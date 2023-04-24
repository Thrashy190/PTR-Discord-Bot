package handlers

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func ShowAndSelectRoles(s *discordgo.Session, m *discordgo.MessageCreate) {

	RolesChannelID := os.Getenv("ROLES_CHANNEL_ID")

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "roles" {
		embed := &discordgo.MessageEmbed{
			Title:       "Roles",
			Description: "Selecciona un rol para obtener acceso a los canales de texto y voz correspondientes",
			Color:       0x4459DF,
			Author: &discordgo.MessageEmbedAuthor{
				IconURL: "",
				Name:    "PTR5.0",
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Roles generales",
					Value: "< @Tlatoani >",
				},
				{
					Name:  "Roles que puedes elegir",
					Value: "Descripcion del rol 2",
				},
			},
		}

		button := []discordgo.Button{
			{
				Emoji: discordgo.ComponentEmoji{
					Name: "🌎",
				},
				Label:    "Minecraft Player",
				Style:    discordgo.PrimaryButton,
				CustomID: "rol1",
			},
			{
				Emoji: discordgo.ComponentEmoji{
					Name: "🎮",
				},
				Label:    "BanditaEsport",
				Style:    discordgo.PrimaryButton,
				CustomID: "rol2",
			},
			{
				Emoji: discordgo.ComponentEmoji{
					Name: "🗡",
				},
				Label:    "DnD Player",
				Style:    discordgo.PrimaryButton,
				CustomID: "rol3",
			},
		}

		s.ChannelMessageSendEmbed(RolesChannelID, embed)
		s.ChannelMessageSendComplex(RolesChannelID, &discordgo.MessageSend{
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						button[0],
						button[1],
						button[2],
					},
				},
			},
		})

		log.Println("ShowRoles command executed")
	}
}
