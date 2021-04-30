package commands

import (
	"fmt"
	"leviathanRewritten/utils"
	"leviathanRewritten/utils/search"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var last_results []search.Result // guardar os resultados obtidos anteriormente
var results_index int            // variável de controle de paginação

var last_query_channel_id string
var last_query_provider string // lembrar se a última pesquisa utilizou Google ou Bing

var last_msg_author_id string         // lembrar quem que pesquisou anteriormente
var last_google_msg_id string         // lembrar qual foi a mensagem que contém os resultados
var last_google_command_msg_id string // lembrar qual a mensagem que contém o comando ";g" e seus argumentos

// EventGoogleMessageEdit é executado quando uma mensagem com ;g for editada por seu autor
func EventGoogleMessageEdit(s *discordgo.Session, usrMsg *discordgo.Message) {
	if usrMsg.ID != last_google_command_msg_id {
		// não executar se a mensagem editada não for a que tem o ;g
		return
	}

	s.ChannelMessageDelete(usrMsg.ChannelID, last_google_msg_id) // apagar a mensagem anterior após re-executar o comando
	HandleCommand(s, usrMsg)
}

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

	// remover react da pessoa para que ela possa reagir novamente
	err := s.MessageReactionRemove(botMessage.ChannelID, botMessage.ID, r.Emoji.Name, r.UserID)

	if err != nil {
		fmt.Println(err)
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
	final := strconv.Itoa(results_index+1) + "/" + strconv.Itoa(len(last_results)) + " " + link

	if last_query_provider == "bing" {
		final = final + " (Bing)"
	}

	s.ChannelMessageEdit(botMessage.ChannelID, botMessage.ID, final)
}

func CommandGoogleExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	if len(args) > 0 {
		query := utils.ArgsTag(args)
		channel, err := s.Channel(m.ChannelID)
		if err != nil {
			return
		}

		query = strings.Replace(query, "_", "%20", -1)

		res, provider, err := search.Search(query, channel.NSFW)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//remove bad args
		if len(res) > 0 {
			if len(last_results) > 0 {
				// colocar um react na pesquisa anterior dizendo que está expirada
				// s.MessageReactionAdd(last_query_channel_id, last_google_msg_id, "🕥")
			}

			last_results = res
			last_msg_author_id = m.Author.ID
			last_query_channel_id = m.ChannelID
			results_index = 0
			last_query_provider = provider

			substr := strings.SplitAfter(res[0].Link, "&sa")
			final := strings.Replace(substr[0], "&sa", "", -1)
			final = strconv.Itoa(results_index+1) + "/" + strconv.Itoa(len(res)) + " " + final

			if provider == "bing" {
				final = final + " (Bing)"
			}

			sent_msg, err := s.ChannelMessageSend(m.ChannelID, final)

			if err != nil {
				// erro ao enviar mensagem
				fmt.Println(err.Error())
				return
			}

			last_google_msg_id = sent_msg.ID
			last_google_command_msg_id = m.ID

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
