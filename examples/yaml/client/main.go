package main

import (
	"fmt"
	"github.com/kitex-contrib/optionloader/configloader/yaml"
	"github.com/kitex-contrib/optionloader/optionloader/client"
)

func main() {
	filePath := "./examples/yaml/client/config.yaml"
	cfg, err := yaml.LoadClientConfig(filePath)
	if err != nil {
		panic(err)
	}
	loader := client.NewOptionLoader()
	options, err := loader.Load(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(options))
}
