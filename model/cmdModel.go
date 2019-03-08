package model

import (
	"fmt"
	"leviathanRewritten/config"
	"leviathanRewritten/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CmdModel struct {
	Author  *discordgo.User
	Message string
	CmdName string
	Args    []string
}

//NewCommand instances a new "object" of type CmdModel
func NewCommand(m *discordgo.Message) *CmdModel {
	author := m.Author
	message := m.Content
	fmt.Println(author, message)
	//get fields
	fields := strings.Fields(message)
	fields[0] = strings.Replace(fields[0], config.Data.Prefix, "", -1)
	cmdName := strings.Replace(fields[0], config.Data.Prefix, "", -1)

	//Parsing the mentions
	message = strings.Replace(message, config.Data.Prefix, "", -1)
	message = strings.Replace(message, "<@", "", -1)
	message = strings.Replace(message, "!", "", -1)
	message = strings.Replace(message, ">", "", -1)
	args := utils.GetArgs(message, fields[0])

	return &CmdModel{Author: author, Message: message, CmdName: cmdName, Args: args}

}
