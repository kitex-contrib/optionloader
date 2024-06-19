package main

import (
	"fmt"
	"github.com/Printemps417/optionloader/utils"
	ymlclient "github.com/Printemps417/optionloader/yml/client"
)

func main() {
	utils.Printpath()
	loader, err := ymlclient.NewClientLoader()
	myreader := ymlclient.NewReader("./client.yml")
	cfg, err := myreader.GetConfig()
	if err != nil {
		println(err)
		return
	}
	fmt.Println("Config: \n", cfg)
	fmt.Println(cfg.Connection.LongConnection.MaxIdlePerAddress)

	loader.SetSource(myreader)
	loader.Load()
	opts, err := loader.GetOptions()
	if err != nil {
		fmt.Println("Error creating clientLoader:", err)
		return
	}
	println("Options: ", opts)
}
