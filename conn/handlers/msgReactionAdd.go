package handlers

import (
	"fmt"
	"leviathanRewritten/commands"

	"github.com/bwmarrin/discordgo"
)

func MessageReactionAdd(s *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	commands.EventGoogleMessageReaction(s, reaction.MessageReaction)
	fmt.Println("React adicionado", reaction.Emoji.Name)
}
