package commands

import (
	"github.com/bwmarrin/discordgo"
)

// CommandPingExec implementa o comando ping... grande comando
func CommandPingExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	s.ChannelMessageSend(m.ChannelID, "\\🏓")
	return
}
