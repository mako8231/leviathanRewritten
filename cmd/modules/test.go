package modules

import (
	"github.com/bwmarrin/discordgo"
	"leviathanRewritten/utils"
)

func Test(s *discordgo.Session, m *discordgo.Message, args... string){
	list := utils.ArgList(args)
	final := utils.ArgsTag(list)
	
	s.ChannelMessageSend(m.ChannelID, final)
}