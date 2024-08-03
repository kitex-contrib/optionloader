package main

import (
	"fmt"
	"github.com/kitex-contrib/optionloader/configloader/yaml"
	"github.com/kitex-contrib/optionloader/optionloader/client/callopt/streamcall"
)

func main() {
	filePath := "./examples/yaml/client/callopt/streamcall/config.yaml"
	cfg, err := yaml.LoadStreamCallConfig(filePath)
	if err != nil {
		panic(err)
	}
	loader := streamcall.NewOptionLoader()
	options, err := loader.Load(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(options))
}
