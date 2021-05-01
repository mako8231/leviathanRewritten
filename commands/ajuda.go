package commands

import (
	"leviathanRewritten/config"
	"leviathanRewritten/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var customHelpMsg string

// CommandAjudaExec representa o comando de ajuda
func CommandAjudaExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	e := utils.NewEmbed()

	e.SetColor(0x2F3136) // cor background do Discord, dá uma aparência legal

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
	e.AddField(true, "Sistema de enquetes", "Use `[1..x]` na sua enquete para fazer uma votação de múltipla escolha, com reações de 1 até x. [Exemplo aqui (PNG)](https://a.pomf.cat/gqkffy.png)")
	e.AddField(true, "Linguagens para tradução", "[Aqui](https://omegat.sourceforge.io/manual-latest/pt_BR/appendix.languages.html) a lista completa de idiomas - os dois tipos de IDs funcionam (por exemplo, ambos `en` e `eng` se referem ao idioma inglês)")
	e.SetFooter("Leviathan " + config.BotVersion)

	// se enviar a lista de canais para o servidor ficar chato demais, é sempre
	// possível enviá-la nos DMs do usuário também...
	sent, _ := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Embed:   e.MessageEmbed,
		Content: customHelpMsg,
	})

	customHelpMsg = ""

	lastCommandOutputMsgChannelID = sent.ChannelID
	lastCommandOutputMsgID = sent.ID
}
