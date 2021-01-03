package commands

import (
	"leviathanRewritten/config"
	"leviathanRewritten/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandAjudaExec representa o comando de ajuda
func CommandAjudaExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	e := utils.NewEmbed()

	e.SetColor(0x36393F) // cor background do Discord, dá uma aparência legal

	commandsString := "```md\n"

	for _, cmd := range Commands {
		// discriminação contra o comando ping: ele não aparece na lista
		if cmd.name == "ping" || cmd.name == "ajuda" {
			continue
		}

		cmdName := cmd.name

		// incluir aliases no nome
		if len(cmd.aliases) > 0 {
			cmdName = cmdName + " (" + strings.Join(cmd.aliases, ", ") + ")"
		}

		commandsString = commandsString + cmd.usage + "\n * " + cmd.description + "\n"
	}

	commandsString = commandsString + "```"

	e.AddField(false, "Comandos", commandsString)
	e.AddField(false, "Sistema de enquetes", "Use `[1..x]` na sua enquete para fazer uma votação de múltipla escolha, com reações de 1 até x. Exemplo: https://a.pomf.cat/gqkffy.png")
	e.SetFooter("Leviathan " + config.BotVersion)

	// se enviar a lista de canais para o servidor ficar chato demais, é sempre
	// possível enviá-la nos DMs do usuário também...
	s.ChannelMessageSendEmbed(m.ChannelID, e.MessageEmbed)
}
