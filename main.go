package main

import (
	"leviathanRewritten/conn"
)

func main(){
	conn.StartClient()
	<-make(chan(struct{}))
}