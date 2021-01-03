package commands

import (
	"leviathanRewritten/config"
	"leviathanRewritten/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Command representa um comando individual
type Command struct {
	name        string
	description string
	usage       string
	aliases     []string
	exec        func(s *discordgo.Session, m *discordgo.Message, args ...string)
}

// Commands representa o registro de comandos
var Commands = make(map[string]Command)

// Aliases representa o registro de aliases, apontando uma alias para o nome real do comando
var Aliases = make(map[string]string)

// RegisterCommands faz o registro de comandos. Deve ser executado pelo menos uma vez na inicialização do bot
func RegisterCommands() {
	// função utilitária de registro
	register := func(cmd Command) {
		Commands[cmd.name] = cmd

		// fazer todas as aliases apontar para o comando também
		for i := 0; i < len(cmd.aliases); i++ {
			Aliases[cmd.aliases[i]] = cmd.name
		}
	}

	// não tem outro jeito...
	// registrar cada comando individualmente
	register(Command{
		name:        "avatar",
		description: "Mostra o avatar de alguém",
		usage:       "avatar @usuário",
		aliases:     []string{},
		exec:        CommandAvatarExec,
	})

	register(Command{
		name:        "dice",
		description: "Gera um número aleatório",
		usage:       "dice 12",
		aliases:     []string{"d"},
		exec:        CommandDiceExec,
	})

	register(Command{
		name:        "google",
		description: "Realiza uma pesquisa no Google",
		usage:       "google termos",
		aliases:     []string{"g"},
		exec:        CommandGoogleExec,
	})

	register(Command{
		name:        "neko",
		description: "Mostra uma neko",
		usage:       "neko lewd?",
		aliases:     []string{},
		exec:        CommandNekoExec,
	})

	register(Command{
		name:        "ajuda",
		description: "Mostra os comandos disponíveis",
		usage:       "ajuda",
		aliases:     []string{"help", "?"},
		exec:        CommandAjudaExec,
	})

	register(Command{
		name:        "calc",
		description: "Resolve uma expressão aritmética",
		usage:       "calc 2 + 2",
		aliases:     []string{"c", "=", "+"},
		exec:        CommandCalcExec,
	})

	register(Command{
		name:        "ping",
		description: "🏓",
		usage:       "ping",
		aliases:     []string{},
		exec:        CommandPingExec,
	})
}

// HandleCommand faz o parsing das mensagens e a execução dos comandos
func HandleCommand(s *discordgo.Session, m *discordgo.Message) {
	// se a mensagem não começa com prefixo, retornar
	if !strings.HasPrefix(m.Content, config.Data.Prefix) {
		return
	}

	args := strings.Split(m.Content, " ")
	cmdName := strings.ToLower(strings.TrimPrefix(args[0], config.Data.Prefix))

	// remover nome de comando da lista de parâmetros
	args = args[1:]

	// remover parâmetros vazios ("")
	args = utils.FilterSliceString(args, func(element string) bool {
		return element != ""
	})

	// procurar por aliases
	if realName, aliasExists := Aliases[cmdName]; aliasExists {
		cmdName = realName
	}

	// verificar se o comando existe
	if cmd, exists := Commands[cmdName]; exists {
		// começar a digitar no canal...
		s.ChannelTyping(m.ChannelID)

		// executar :)
		cmd.exec(s, m, args...)
	}
}
