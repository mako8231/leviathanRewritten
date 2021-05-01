package commands

import (
	"github.com/bwmarrin/discordgo"
)

// CommandPingExec implementa o comando ping... grande comando
func CommandPingExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	sent, _ := s.ChannelMessageSend(m.ChannelID, "\\🏓")
	lastCommandOutputMsgChannelID = sent.ChannelID
	lastCommandOutputMsgID = sent.ID
}
