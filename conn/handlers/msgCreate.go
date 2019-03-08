package handlers

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	"leviathanRewritten/cmd"
	"leviathanRewritten/config"
)


func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	//Instance the channel object for verification 
	channel, err := s.Channel(m.ChannelID)
	 if err != nil{
	 	return 
	}
	
	//Ignores messages from bots
	if m.Author.Bot{
		return 
	}

	//only the prefix will crash the bot, so we'll handle this issue
	if m.Content == config.Data.Prefix{
		return 
	}
	

	//If this message is from a guild 
	if channel.Type == 0{
		if strings.HasPrefix(m.Content, config.Data.Prefix){
			cmd.MessageController(m.Message, s)
		}
	}
	
}
