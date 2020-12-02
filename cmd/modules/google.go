package modules

import (
	"fmt"
	"leviathanRewritten/utils"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var last_results []utils.Result // guardar os resultados obtidos anteriormente
var results_index int           // variável de controle de paginação
var last_msg_author_id string   // lembrar quem que pesquisou anteriormente
var last_google_msg_id string   // lembrar qual foi a mensagem mais recente do ;g

// EventGoogleMessageReaction é executado quando há uma reação na mensagem do comando de Google.
func EventGoogleMessageReaction(s *discordgo.Session, botMessage *discordgo.Message, r *discordgo.MessageReaction) {
	if r.UserID != last_msg_author_id {
		// não executar se quem reagiu é não for quem executou ;g
		return
	}

	if botMessage.ID != last_google_msg_id {
		// não executar se essa não for a pesquisa mais recente
		return
	}

	if r.Emoji.Name == "◀️" && results_index > 0 {
		results_index = results_index - 1
	} else if r.Emoji.Name == "▶️" && results_index < len(last_results)-1 {
		results_index = results_index + 1
	} else {
		return
	}

	substr := strings.SplitAfter(last_results[results_index].Link, "&sa")
	link := strings.Replace(substr[0], "&sa", "", -1)

	// remover react da pessoa para que ela possa reagir novamente
	err := s.MessageReactionRemove(botMessage.ChannelID, botMessage.ID, r.Emoji.Name, r.UserID)

	if err != nil {
		fmt.Println(err)
	}

	s.ChannelMessageEdit(botMessage.ChannelID, botMessage.ID, "Resultado #"+strconv.Itoa(results_index+1)+" "+link)
}

func Google(s *discordgo.Session, m *discordgo.Message, args ...string) {
	if len(args) > 0 {
		query := utils.ArgsTag(args)
		channel, err := s.Channel(m.ChannelID)
		if err != nil {
			return
		}

		query = strings.Replace(query, "_", "%20", -1)

		res, err := utils.GoogleParse(query, channel.NSFW)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//remove bad args
		if len(res) > 0 {
			last_results = res
			last_msg_author_id = m.Author.ID
			results_index = 0

			substr := strings.SplitAfter(res[0].Link, "&sa")
			final := strings.Replace(substr[0], "&sa", "", -1)
			sent_msg, err := s.ChannelMessageSend(m.ChannelID, "Resultado #"+strconv.Itoa(results_index+1)+" "+final)

			if err != nil {
				// erro ao enviar mensagem
				fmt.Println(err.Error())
				return
			}

			last_google_msg_id = sent_msg.ID

			// adicionar reacts na mensagem
			s.MessageReactionAdd(m.ChannelID, sent_msg.ID, "◀️")
			s.MessageReactionAdd(m.ChannelID, sent_msg.ID, "▶️")
		} else {
			msg := utils.NewEmbed()
			msg.SetColor(utils.Yellow)
			msg.SetTitle("Erro")
			msg.SetDescription("Nenhum resultado encontrado")
			s.ChannelMessageSendEmbed(m.ChannelID, msg.MessageEmbed)
		}

	}

	return

	//fmt.Println(teste[1].Link)
}
