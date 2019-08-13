package utils

import (
	"github.com/bwmarrin/discordgo"
)

func Pool(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.MessageReactionAdd(m.ChannelID, m.ID, "👍")
	s.MessageReactionAdd(m.ChannelID, m.ID, "👎")
}
