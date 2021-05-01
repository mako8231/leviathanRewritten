package handlers

import (
	"leviathanRewritten/commands"
	"leviathanRewritten/config"
	"leviathanRewritten/subsystems/pool"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Instance the channel object for verification
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return
	}

	//Ignores messages from bots
	if m.Author.Bot {
		return
	}

	//only the prefix will crash the bot, so we'll handle this issue
	if m.Content == config.Data.Prefix {
		return
	}

	//If this message is from a guild
	if channel.Type == 0 {
		//código merda ae vou melhorar quando eu achar que a leviathan será algo maior pra ficar na SA (não vai acontecer)
		if m.ChannelID == config.Data.PoolChan && m.GuildID == config.Data.Server {
			pool.CreatePool(s, m.Message)
		}

		if strings.HasPrefix(m.Content, config.Data.Prefix) {
			commands.HandleCommand(s, m.Message)
		}
	}

}
