package conn

import (
	"leviathanRewritten/config"
	"github.com/bwmarrin/discordgo"
	"log"
	"leviathanRewritten/conn/handlers"
	"leviathanRewritten/httpServer"
)

var (
	Session *discordgo.Session
	BotID string
)
//StartClient starts a new discord session
func StartClient(){
	//Reading the configuration file...
	config.LoadConfig()
	//Starting session
	s, err := discordgo.New("Bot "+config.Data.Token)
	if err != nil{
		log.Fatal(err.Error())
	}

	//Get user data
	user, err := s.User("@me")
	if err != nil{
		log.Fatal(err.Error())
	}


	BotID = user.ID
	Session = s

	err = Session.Open()

	Session.AddHandler(handlers.MessageCreate)

	if err != nil{
		log.Fatal(err.Error())
	}
	log.Println(user.Username + " started")
	httpServer.StartServer(config.Data.DefaultPort)
	
}

