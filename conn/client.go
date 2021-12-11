package conn

import (
	"leviathanRewritten/commands"
	"leviathanRewritten/config"
	"leviathanRewritten/conn/handlers"
	"leviathanRewritten/httpServer"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Session *discordgo.Session
	BotID   string
)

//StartClient starts a new discord session
func StartClient() {
	//Reading the configuration file...
	config.LoadConfig()
	//Starting session
	s, err := discordgo.New("Bot " + config.Data.Token)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Get user data
	user, err := s.User("@me")
	if err != nil {
		log.Fatal(err.Error())
	}

	BotID = user.ID
	Session = s

	err = Session.Open()

	Session.AddHandler(handlers.MessageCreate)
	Session.AddHandler(handlers.MessageReactionAdd)
	Session.AddHandler(handlers.MessageReactionRemove)
	Session.AddHandler(handlers.MessageEdit)
	Session.AddHandler(handlers.MessageDelete)

	commands.RegisterCommands()

	Session.UpdateStreamingStatus(0, config.Data.Prefix+"ajuda | v"+config.BotVersion, "https://www.twitch.tv/masao24")

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(user.Username + " started")
	httpServer.StartServer(config.Data.DefaultPort)
}
