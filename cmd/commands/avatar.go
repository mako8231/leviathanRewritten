package commands

import (
	"leviathanRewritten/utils"

	"github.com/bwmarrin/discordgo"
)

//Avatar Get the user avatar
func Avatar(s *discordgo.Session, m *discordgo.Message, args ...string) {
	var username string
	var avatarurl string
	text := "'s avatar"

	//Verify args
	args = utils.ArgList(args)
	//Get users
	usr, err := s.User(args[0])
	if err != nil {
		username = m.Author.Username
		avatarurl = m.Author.AvatarURL("1024")
	} else {
		username = usr.Username
		avatarurl = usr.AvatarURL("1024")
	}
	//Creating embed message
	e := utils.NewEmbed()
	e.SetColor(utils.Blue)
	e.SetTitle(username + " " + text)
	e.SetImage(avatarurl)
	e.SetAuthor(m.Author.AvatarURL("1024"), m.Author.Username)
	msg := e.MessageEmbed
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, msg)
	if err != nil {
		return

	}
}
