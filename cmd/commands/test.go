package commands

import (
	"leviathanRewritten/utils"

	"github.com/bwmarrin/discordgo"
)

func Test(s *discordgo.Session, m *discordgo.Message, args ...string) {
	list := utils.ArgList(args)
	final := utils.ArgsTag(list)

	s.ChannelMessageSend(m.ChannelID, final)
}
