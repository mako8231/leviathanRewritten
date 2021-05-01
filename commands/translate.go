package commands

import (
	"encoding/json"
	"fmt"
	"leviathanRewritten/utils"
	"strings"

	"net/url"

	"github.com/bwmarrin/discordgo"
)

type translateResult [][]string

// CommandTranslateExec representa o comando de tradução
func CommandTranslateExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	if len(args) < 2 {
		customHelpMsg = "❌ ***Parâmetros incorretos! Mostrando __;ajuda__ ***"
		CommandAjudaExec(s, m)
		return
	}

	targetLang := args[:1][0]
	text := strings.Join(args[1:], " ")

	responseBody, err := utils.GetDoc("https://translate.googleapis.com/translate_a/single" +
		"?client=gtx" +
		"&sl=auto" +
		"&tl=" + url.QueryEscape(targetLang) +
		"&dt=t" +
		"&q=" + url.QueryEscape(text) +
		"&ie=UTF-8" +
		"&oe=UTF-8")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result []translateResult
	json.Unmarshal(responseBody, &result)
	translated := result[0][0][0]

	if len(translated) < 1 {
		sent, _ := s.ChannelMessageSend(m.ChannelID, "❌ ***Tradução somehow nula — pode ter sido parâmetros incorretos***")

		lastCommandOutputMsgChannelID = sent.ChannelID
		lastCommandOutputMsgID = sent.ID
	} else {
		e := utils.NewEmbed()
		e.SetColor(0x2F3136)
		e.SetDescription(utils.StringLimitChar(translated, 999))
		e.SetFooter(utils.StringLimitChar(text, 20))
		sent, _ := s.ChannelMessageSendEmbed(m.ChannelID, e.MessageEmbed)

		lastCommandOutputMsgChannelID = sent.ChannelID
		lastCommandOutputMsgID = sent.ID
	}
}
