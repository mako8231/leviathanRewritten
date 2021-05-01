package handlers

import (
	"fmt"
	"leviathanRewritten/commands"

	"github.com/bwmarrin/discordgo"
)

func MessageEdit(s *discordgo.Session, mEdit *discordgo.MessageUpdate) {
	// pegar o objeto de mensagem
	m, err := s.ChannelMessage(mEdit.ChannelID, mEdit.ID)
	if err != nil || m.Author.Bot {
		// erro ao obter a mensagem... e daí? não creio que chegaremos a esse ponto
		// no programa em algum momento
		return
	}

	commands.HandleCommandEdit(s, m)

	fmt.Println("Mensagem editada")
}
