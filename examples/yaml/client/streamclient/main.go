package main

import (
	"fmt"
	"github.com/kitex-contrib/optionloader/configloader/yaml"
	"github.com/kitex-contrib/optionloader/optionloader/client/streamclient"
)

func main() {
	filePath := "./examples/yaml/client/streamclient/config.yaml"
	cfg, err := yaml.LoadStreamClientConfig(filePath)
	if err != nil {
		panic(err)
	}
	loader := streamclient.NewOptionLoader()
	options, err := loader.Load(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(options))
}
