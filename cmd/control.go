package cmd

import (
	"github.com/bwmarrin/discordgo"
	"leviathanRewritten/model"
	"leviathanRewritten/cmd/modules"
	"leviathanRewritten/utils"
)

var (

)

func MessageController(m *discordgo.Message, s *discordgo.Session){
	var c *model.CmdModel
	c = model.NewCommand(m)
	commandCaller(c, s, m) 
}

func commandCaller(cmd *model.CmdModel, s *discordgo.Session, m *discordgo.Message){
	s.ChannelTyping(m.ChannelID)
	if utils.Compare(cmd.CmdName, "ping"){
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if utils.Compare(cmd.CmdName, "avatar"){
		modules.Avatar(s, m, cmd.Args...)
	}

	if utils.Compare(cmd.CmdName, "test"){
		modules.Test(s, m, cmd.Args...)
	}

	if utils.Compare(cmd.CmdName, "neko"){
		modules.Neko(s, m, cmd.Args...)
	}

	if utils.Compare(cmd.CmdName, "google") || utils.Compare(cmd.CmdName, "g"){
		modules.Google(s, m, cmd.Args...)
	}
}
