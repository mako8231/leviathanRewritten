package handlers

import (
	"fmt"
	"leviathanRewritten/commands"

	"github.com/bwmarrin/discordgo"
)

func MessageReactionRemove(s *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	commands.EventGoogleMessageReaction(s, reaction.MessageReaction)
	fmt.Println("React removido", reaction.Emoji.Name)
}
