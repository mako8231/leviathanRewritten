package handlers

import (
	"fmt"
	"leviathanRewritten/commands"

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

	currentUser, err := s.User("@me")
	if err != nil {
		// erro ao obter o usuário atual
		return
	}

	// Somente executar se a mensagem reagida é nossa
	if m.Author.ID == currentUser.ID {
		commands.EventGoogleMessageReaction(s, m, reaction.MessageReaction)
		return
	}

	fmt.Println("React adicionado", reaction.Emoji.Name)
}
