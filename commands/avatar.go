package commands

import (
	"leviathanRewritten/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandAvatar Get the user avatar
func CommandAvatarExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	var userID string
	var user *discordgo.User

	userID = m.Author.ID

	if len(args) > 0 {
		userID = args[0]                           // pegar o ID puro OU a menção
		userID = strings.TrimPrefix(userID, "<@!") // se for menção (<@0123456789>), extrair ID
		userID = strings.TrimPrefix(userID, "<@")
		userID = strings.TrimSuffix(userID, ">")
	}

	// tentar pegar o usuário
	if fetchedUser, err := s.User(userID); err == nil {
		user = fetchedUser
	} else {
		// ocorreu um erro ao pegar o usuário... responder com o avatar do autor mesmo
		user = m.Author
	}

	//Creating embed message
	e := utils.NewEmbed()
	e.SetColor(0x36393F)
	e.SetTitle("🌟 " + user.Username)
	e.SetImage(user.AvatarURL("1024"))
	e.SetFooter("ID " + user.ID)

	sent, _ := s.ChannelMessageSendEmbed(m.ChannelID, e.MessageEmbed)
	lastCommandOutputMsgChannelID = sent.ChannelID
	lastCommandOutputMsgID = sent.ID
}
