package pool

import (
	"regexp"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var numberReactions = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣"}

var multichoicePollRegex = regexp.MustCompile("^\\[(?P<StartIndex>[1-9])\\.\\.(?P<EndIndex>[1-9])\\]")

// CreatePool cria uma votação na mensagem especificada.
func CreatePool(s *discordgo.Session, m *discordgo.Message) {
	matches := multichoicePollRegex.FindStringSubmatch(m.Content)

	if len(matches) > 0 {
		startIndex, _ := strconv.Atoi(matches[1])
		endIndex, _ := strconv.Atoi(matches[2])

		if endIndex < startIndex {
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			return
		}

		for i := startIndex - 1; i < endIndex; i++ {
			s.MessageReactionAdd(m.ChannelID, m.ID, numberReactions[i])
		}
	} else {
		s.MessageReactionAdd(m.ChannelID, m.ID, "👍")
		s.MessageReactionAdd(m.ChannelID, m.ID, "👎")
	}
}
