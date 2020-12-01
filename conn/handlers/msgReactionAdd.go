package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageReactionAdd(s *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	// pegar o objeto de mensagem
	m, err := s.ChannelMessage(reaction.ChannelID, reaction.MessageID)
	if err != nil {
		// erro ao obter a mensagem... e daí? não creio que chegaremos a esse ponto
		// no programa em algum momento
		return
	}

	// Ignores messages from bots
	if m.Author.Bot {
		return
	}

	fmt.Println("React adicionado", reaction.Emoji.Name)
}
