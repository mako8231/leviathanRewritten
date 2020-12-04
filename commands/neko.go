package commands

import (
	"fmt"
	"leviathanRewritten/utils"

	"github.com/bwmarrin/discordgo"
)

var url = "https://nekos.life/api/v2/img/"

func CommandNekoExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	//First, instance a channel struct
	var finalURL string
	var color int
	channel, _ := s.Channel(m.ChannelID)
	//Now, lets work with args
	args = utils.ArgList(args)

	//Now, lets handle the args
	if utils.Compare(args[0], "lewd") {
		if !channel.NSFW {
			utils.SendWarning(utils.NSFWarning, s, m)
			return
		}
		color = utils.Green
		finalURL = url + "lewd"
	} else {
		color = utils.Blue
		finalURL = url + "neko"
	}

	b, err := utils.GetDoc(finalURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	obj, err := utils.MapJSON(b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Get the image url from map
	img := fmt.Sprint(string(obj["url"]))
	img = utils.Remove(img, `"`)
	//================================
	//Make a new embed message
	embed := utils.NewEmbed()
	embed.SetColor(color)
	embed.SetTitle("🐱 Neko 🐱")
	embed.SetAuthor(m.Author.AvatarURL("1024"), m.Author.Username)
	embed.SetImage(img)
	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)

}
