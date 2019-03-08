package httpServer

import (
	"os"
	"log"
	"net/http"
	"time"
)

func StartServer(port string){
	RefreshRequest(port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func RefreshRequest(port string){
	tick := time.Tick(2 * time.Minute)
	var serverName string

	for {
		select{
		case <-tick:
			serverName = os.Getenv("DOMAIN_NAME")
			if serverName == ""{
				serverName = "http://127.0.0.1:8080"
			}
			_, err := http.Get(serverName)
			if err != nil{
				log.Println(err.Error())
			} 
		}
	}

}
