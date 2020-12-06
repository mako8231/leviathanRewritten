package commands

import (
	"leviathanRewritten/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandAjudaExec representa o comando de ajuda
func CommandAjudaExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	e := utils.NewEmbed()

	e.SetTitle(("Lista de comandos"))
	e.SetColor(0x36393F) // cor background do Discord, dá uma aparência legal

	for _, cmd := range Commands {
		// discriminação contra o comando ping: ele não aparece na lista
		if cmd.name == "ping" {
			continue
		}

		cmdName := cmd.name

		// incluir aliases no nome
		if len(cmd.aliases) > 0 {
			cmdName = cmdName + " (" + strings.Join(cmd.aliases, ", ") + ")"
		}

		e.AddField(false, cmdName, strings.Join([]string{
			cmd.description,
			"`" + cmd.usage + "`",
		}, "\n"))
	}

	// se enviar a lista de canais para o servidor ficar chato demais, é sempre
	// possível enviá-la nos DMs do usuário também...
	s.ChannelMessageSendEmbed(m.ChannelID, e.MessageEmbed)
}
