package main

import (
	"fmt"
	"github.com/Printemps417/optionloader/utils"
	ymlclient "github.com/Printemps417/optionloader/yml/client"
)

func main() {
	utils.Printpath()
	println("Client")
	myreader := ymlclient.NewReader("./client.yml")
	cfg, err := myreader.GetConfig()
	if err != nil {
		println(err)
		return
	}
	fmt.Println("Config: \n", cfg)
	fmt.Println(cfg.Connection.LongConnection.MaxIdlePerAddress)
}
