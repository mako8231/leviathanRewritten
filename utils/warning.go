package utils

import (
	"github.com/bwmarrin/discordgo"

)

func SendWarning(text string, s *discordgo.Session, m *discordgo.Message){
	e := NewEmbed()
	e.SetAuthor(m.Author.AvatarURL("1024"), m.Author.Username)
	e.SetDescription(text)
	e.SetColor(Yellow)
	s.ChannelMessageSendEmbed(m.ChannelID, e.MessageEmbed)
}