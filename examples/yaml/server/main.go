package main

import (
	"fmt"
	"github.com/kitex-contrib/optionloader/configloader/yaml"
	"github.com/kitex-contrib/optionloader/optionloader/server"
)

func main() {
	filePath := "./examples/yaml/server/config.yaml"
	cfg, err := yaml.LoadServerConfig(filePath)
	if err != nil {
		panic(err)
	}
	loader := server.NewOptionLoader()
	options, err := loader.Load(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(options))
}
