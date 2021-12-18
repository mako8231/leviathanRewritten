package commands

import (
	"encoding/json"
	"fmt"
	"leviathanRewritten/utils"
	"strings"

	"net/url"

	"github.com/bwmarrin/discordgo"
)

type translateOutput [][][]string

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

	var output translateOutput
	var translated string = ""
	json.Unmarshal(responseBody, &output)

	// Essa API adora mostrar os resultados em vetor 4D. Muito bonito!
	for i := range output[0] {
		translated = translated + output[0][i][0]
	}

	if len(translated) < 1 {
		sent, _ := s.ChannelMessageSend(m.ChannelID, "❌ ***Tradução nula — pode ter sido parâmetros incorretos, ou a API mudou***")

		lastCommandOutputMsgChannelID = sent.ChannelID
		lastCommandOutputMsgID = sent.ID
	} else {
		e := utils.NewEmbed()
		e.SetColor(0x2F3136)
		e.SetDescription(utils.StringLimitChar(translated, 999))
		sent, _ := s.ChannelMessageSendEmbed(m.ChannelID, e.MessageEmbed)

		lastCommandOutputMsgChannelID = sent.ChannelID
		lastCommandOutputMsgID = sent.ID
	}
}
