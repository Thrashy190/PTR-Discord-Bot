package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ShowRoles(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	//genera una carta con un header y un footer de prueba y la envia al canal Roles ademas de poner 3 botones abajo
	if m.Content == "roles" {

		//crea el embed
		embed := &discordgo.MessageEmbed{
			Title:       "Roles",
			Description: "Selecciona un rol para obtener acceso a los canales de texto y voz correspondientes",
			Color:       0x4459DF,
			Author: &discordgo.MessageEmbedAuthor{
				IconURL: "https://cdn.discordapp.com/attachments/819496062769440020/820000000000000000/LogoPTR.png",
				Name:    "PTR5.0",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Roles generales",
					Value:  "@Tlatoani  --Admin principal\n @Semidioses  --Moderadores\n @Humanos  --Rol general",
					
				},
				{
					Name:   "Roles que puedes elegir",
					Value:  "Descripcion del rol 2",
				},
			},
		}

		button := []discordgo.Button{
			{
				Label:    "Minecraft Player",
				Style:    discordgo.PrimaryButton,
				CustomID: "rol1",
			},
			{
				Label: "BanditaEsport",
				Style: discordgo.PrimaryButton,
				CustomID: "rol2",
			},
			{
				Label: "DnD Player",
				Style: discordgo.PrimaryButton,
				CustomID: "rol3",
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
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
