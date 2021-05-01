package handlers

import (
	"fmt"
	"leviathanRewritten/commands"

	"github.com/bwmarrin/discordgo"
)

func MessageDelete(s *discordgo.Session, mDel *discordgo.MessageDelete) {
	commands.HandleCommandDelete(s, mDel.ID)

	fmt.Println("Mensagem apagada")
}
